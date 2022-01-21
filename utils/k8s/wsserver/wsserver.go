package main

/**
思路来源： 360 wayne 以及 https://github.com/kubernetes/client-go/issues/204
*/

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/utils"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gorilla/websocket"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/homedir"
)

var (
	restConf *rest.Config
	sshReq   *rest.Request
	executor remotecommand.Executor
	wsConn   *utils.WsConnection

	podName       string
	podNs         string
	containerName string
	handler       *streamHandler
	err           error
)

func GetRestConf() (restConf *rest.Config, err error) {

	var kubeconfig []byte
	var clientErr error
	var restconfig *rest.Config

	if home := homedir.HomeDir(); home != "" {
		if kubeconfig, clientErr = ioutil.ReadFile(filepath.Join(home, ".kube", "config")); clientErr != nil {
			return
		}

		restconfig, clientErr = clientcmd.RESTConfigFromKubeConfig(kubeconfig)
		if clientErr != nil {
			panic(clientErr)
		}

	}
	return restconfig, err
}

// ssh流式处理器
type streamHandler struct {
	wsConn      *utils.WsConnection
	resizeEvent chan remotecommand.TerminalSize
}

// web终端发来的包
type xtermMessage struct {
	MsgType string `json:"type"`  // 类型:resize客户端调整终端, input客户端输入
	Input   string `json:"input"` // msgtype=input情况下使用
	Rows    uint16 `json:"rows"`  // msgtype=resize情况下使用
	Cols    uint16 `json:"cols"`  // msgtype=resize情况下使用
}

// executor回调获取web是否resize
func (handler *streamHandler) Next() (size *remotecommand.TerminalSize) {
	ret := <-handler.resizeEvent
	size = &ret
	return
}

// executor回调读取web端的输入
//read called in a loop from remotecommand as long  as the process is running.
func (handler *streamHandler) Read(p []byte) (size int, err error) {
	var (
		msg      *utils.WsMessage
		xtermMsg xtermMessage
	)

	// 读web发来的输入
	if msg, err = handler.wsConn.WsRead(); err != nil { //du
		handler.wsConn.WsClose()
		return
	}

	fmt.Println("read msg:", string(msg.Data), msg.MessageType)

	if err = json.Unmarshal(msg.Data, &xtermMsg); err != nil {
		fmt.Println(err)
		return
	}

	// 解析客户端请求

	//web ssh调整了终端大小
	if xtermMsg.MsgType == "resize" {
		// 放到channel里，等remotecommand executor调用我们的Next取走
		handler.resizeEvent <- remotecommand.TerminalSize{Width: xtermMsg.Cols, Height: xtermMsg.Rows}
	} else if xtermMsg.MsgType == "input" { // web ssh终端输入了字符
		// copy到p数组中
		size = len(xtermMsg.Input)
		copy(p, xtermMsg.Input)
	}
	return

}

// executor回调向web端输出
func (handler *streamHandler) Write(p []byte) (size int, err error) {
	var (
		copyData []byte
	)

	fmt.Println("web端输出:", string(p))
	// 产生副本
	copyData = make([]byte, len(p))
	copy(copyData, p)
	size = len(p)
	err = handler.wsConn.WsWrite(websocket.TextMessage, copyData) //写到channel
	return
}

func wsHandler(resp http.ResponseWriter, req *http.Request) {

	// 解析GET参数
	if err = req.ParseForm(); err != nil {
		return
	}

	vars := req.URL.Query()

	// fmt.Println("vars:", vars)

	podNs = vars.Get("podNs")
	podName = vars.Get("podName")

	containerName = vars.Get("containerName")

	// 得到websocket长连接
	if wsConn, err = utils.InitWebsocket(resp, req); err != nil {
		wsConn.WsClose()
		return
	}

	validbashs := []string{"/bin/bash", "/bin/sh"}
	if isValidBash(validbashs, "") {
		cmds := []string{""}
		err = startProcess(podName, podNs, containerName, cmds)
	} else {
		for _, testShell := range validbashs {
			cmd := []string{testShell}
			if err = startProcess(podName, podNs, containerName, cmd); err != nil {
				continue
			}
		}
	}

}

func startProcess(podName string, podNs string, containerName string, cmd []string) error {

	restConf, err := GetRestConf()
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		panic(err.Error())
	}

	sshReq = clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(podNs).
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			Container: containerName,
			Command:   cmd,
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)

	// 创建到容器的连接
	if executor, err = remotecommand.NewSPDYExecutor(restConf, "POST", sshReq.URL()); err != nil {
		wsConn.WsClose()
		return err

	}

	// 配置与容器之间的数据流处理回调
	handler = &streamHandler{wsConn: wsConn, resizeEvent: make(chan remotecommand.TerminalSize)}
	if err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             handler,
		Stdout:            handler,
		Stderr:            handler,
		TerminalSizeQueue: handler,
		Tty:               true,
	}); err != nil {
		fmt.Println("handler", err)
		return err

	}
	return err

}

func isValidBash(isValidbash []string, shell string) bool {
	for _, isValidbash := range isValidbash {
		if isValidbash == shell {
			return true
		}
	}
	return false
}

func main() {

	http.HandleFunc("/ssh", wsHandler)
	http.ListenAndServe("localhost:7777", nil)
}
