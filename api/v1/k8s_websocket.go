package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

var (
	restConf *rest.Config
	sshReq   *rest.Request
	executor remotecommand.Executor
	wsConn   *utils.WsConnection

	podName       string
	podNs         string
	containerName string
	userName      string
	pipelineRunName string
	podCmd        string
	handler       *streamHandler
	err           error
)

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

	//fmt.Println("read msg:", string(msg.Data), msg.MessageType)

	if err = json.Unmarshal(msg.Data, &xtermMsg); err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

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

	//fmt.Println("web端输出:", string(p))
	// 产生副本
	copyData = make([]byte, len(p))
	copy(copyData, p)
	size = len(p)
	err = handler.wsConn.WsWrite(websocket.TextMessage, copyData) //写到channel
	return
}

func startProcess(podName string, podNs string, containerName string, cmd []string, userName string) error {

	// restConf, err := utils.GetRestConf()

	// if err != nil {
	// 	return err
	// }

	// clientset, err := kubernetes.NewForConfig(restConf)

	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(userName))

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
	if executor, err = remotecommand.NewSPDYExecutor(global.GetK8sConfig(userName), "POST", sshReq.URL()); err != nil {
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
		global.GVA_LOG.Error(err.Error())
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


func WsHandler(resp http.ResponseWriter, req *http.Request) {
	// 解析GET参数
	if err = req.ParseForm(); err != nil {
		return
	}

	vars := req.URL.Query()
	podNs = vars.Get("podNs")
	podName = vars.Get("podName")
	containerName = vars.Get("containerName")
	userName = vars.Get("userName")
	podCmd = fmt.Sprintf("%s/%s", "/bin", vars.Get("podCmd"))

	// 得到websocket长连接
	if wsConn, err = utils.InitWebsocket(resp, req); err != nil {
		wsConn.WsClose()
		return
	}

	validbashs := []string{"/bin/bash", "/bin/sh"}
	if isValidBash(validbashs, "") {
		cmds := []string{""}
		err = startProcess(podName, podNs, containerName, cmds, userName)
	} else {
		for _, testShell := range validbashs {
			if testShell == podCmd {
				cmd := []string{testShell}
				if err = startProcess(podName, podNs, containerName, cmd, userName); err != nil {
					global.GVA_LOG.Error(err.Error())
				}
			}

		}
	}

}


func WsGetBuildHistory(resp http.ResponseWriter, req *http.Request) {

	// 得到websocket长连接
	if wsConn, err = utils.InitWebsocket(resp, req); err != nil {
		wsConn.WsClose()
		return
	}
	vars := req.URL.Query()
	userName = vars.Get("userName")

	handler = &streamHandler{wsConn: wsConn, resizeEvent: make(chan remotecommand.TerminalSize)}
	config := global.GetK8sConfig(userName)
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		return
	}
	//fmt.Println(1)
	watcher, err := clientset.TektonV1alpha1().PipelineRuns("tekton-pipelines").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return
	}
	var pipelineRuns model.TektonPipelineRunsList
	//listResult := make([]interface{}, 0)
	for event := range watcher.ResultChan() {
		pipelineRun := event.Object.(*v1alpha1.PipelineRun)
		//for _, item := range pipeline.Items {
		//	fmt.Printf("%s/%s status:%s\n", item.Namespace, item.Namespace, item.Status.Conditions[0])
		//}
		//fmt.Printf("%s/%s status:%s\n", pipelineRun.Namespace, pipelineRun.Name, pipelineRun.Status.Conditions)
		pipelineRuns.Name = pipelineRun.Name
		pipelineRuns.Namespace = pipelineRun.Namespace
		pipelineRuns.Pipeline = pipelineRun.Spec.PipelineRef.Name
		pipelineRuns.TriggersEventId = pipelineRun.Labels["triggers.tekton.dev/triggers-eventid"]
		pipelineRuns.Message = pipelineRun.Status.Conditions[0].Message
		pipelineRuns.Reason = pipelineRun.Status.Conditions[0].Reason
		pipelineRuns.Status = string(pipelineRun.Status.Conditions[0].Status)
		pipelineRuns.CompletionTime = pipelineRun.Status.CompletionTime
		if pipelineRun.Status.CompletionTime == nil {
			pipelineRuns.CompletionTime = &metav1.Time{Time: time.Now()}
		}
		//pipelineRuns.CompletionTime = pipelineRun.Status.CompletionTime
		pipelineRuns.LastTransitionTime = pipelineRun.Status.Conditions[0].LastTransitionTime.Inner
		pipelineRuns.StartTime = pipelineRun.Status.StartTime
		pipelineRuns.CreateTime = pipelineRun.CreationTimestamp.Time
		d, _ := json.Marshal(pipelineRuns)
		handler.Write(d)
	}
}



func WsGetTaskRun(resp http.ResponseWriter, req *http.Request) {

	// 得到websocket长连接
	if wsConn, err = utils.InitWebsocket(resp, req); err != nil {
		wsConn.WsClose()
		return
	}
	vars := req.URL.Query()
	userName = vars.Get("userName")
	pipelineRunName = vars.Get("pipelineRunName")

	handler = &streamHandler{wsConn: wsConn, resizeEvent: make(chan remotecommand.TerminalSize)}
	config := global.GetK8sConfig(userName)
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		return
	}
	watcher, err := clientset.TektonV1alpha1().TaskRuns("tekton-pipelines").Watch(context.TODO(), metav1.ListOptions{
		LabelSelector: "tekton.dev/pipelineRun="+pipelineRunName,
	})
	if err != nil {
		return
	}
	//var pipelineRuns model.TektonPipelineRunsList
	//listResult := make([]interface{}, 0)
	var taskRun model.TektonTaskRunList
	var taskStep model.TaskStep
	for event := range watcher.ResultChan() {
		t := event.Object.(*v1alpha1.TaskRun)
		taskRun.Name = t.Name
		taskRun.Namespace = t.Namespace
		taskRun.PodName = t.Status.PodName
		//taskRun.TaskRefName = t.Spec.TaskRef.Name
		if t.Status.Conditions != nil {
			taskRun.Status = string(t.Status.Conditions[0].Status)
			taskRun.Reason = t.Status.Conditions[0].Reason
			taskRun.Message = t.Status.Conditions[0].Message
		} else {
			taskRun.Status = "Unknown"
			taskRun.Reason = "Running"
			taskRun.Message = ""
		}
		if taskRun.Reason == "Pending" {
			continue
		}

		taskRun.TaskRefName = t.Labels["tekton.dev/pipelineTask"]
		taskRun.PipelineRunName = pipelineRunName
		var taskStepList []model.TaskStep
		stepStatusFlag := false
		for _, step := range t.Status.Steps {
			taskStep.StepName = step.ContainerName
			taskStep.StepContainerName = step.ContainerName
			if step.Terminated != nil {
				taskStep.StartAt = step.Terminated.StartedAt
				taskStep.FinishAt = step.Terminated.FinishedAt
				taskStep.Reason = step.Terminated.Reason
			} else {
				if stepStatusFlag == false {
					taskStep.Reason = "Running"
					stepStatusFlag = true
				} else {
					taskStep.Reason = "Waiting"
				}
				taskStep.StartAt = step.Running.StartedAt
				taskStep.FinishAt = metav1.Now()

			}

			taskStepList = append(taskStepList, taskStep)

		}
		taskRun.Steps = taskStepList
		d, _ := json.Marshal(taskRun)
		handler.Write(d)
	}
}


