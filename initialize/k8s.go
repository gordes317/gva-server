package initialize

import (
	"flag"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/service"
	"go.uber.org/zap"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

//从配置文件读取信息
func init() {

	var kubeconfig *string
	var clientErr error

	//fmt.Println(os.Getwd()/kubeconfig/.k8sconfig)

	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取程序当前路径失败:%v", err.Error())
	}
	kubeconfig = flag.String("kubeconfig", filepath.Join(currentPath, "kubeconfig", ".k8sconfig"), "(optional) absolute path to the kubeconfig file")
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//
	//	configPath := fmt.Sprintf("%s/%s/%s", currentPath, "kubeconfig", ".k8sconfig")
	//	kubeconfig = flag.String("kubeconfig", "", configPath)
	//}

	flag.Parse()

	// use the current context in kubeconfig
	global.GVA_K8SConfig, clientErr = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if clientErr != nil {
		panic(clientErr.Error())
	} else {
		fmt.Println("connect k8s  success")
	}
	global.GVA_K8SConfigMap = make(map[string]*rest.Config)
	global.GVA_K8SConfigMap["defaultConfig"] = global.GVA_K8SConfig

}

//从数据库读取配置初始化K8S

func KuberConfigInit() *rest.Config {

	if err, list := service.GetAllClusterInfo(); err != nil {
		global.GVA_LOG.Error("读取集群信息失败", zap.Any("err", err))
	} else {

		fmt.Println("获得集群列表:", list)
	}
	return nil

}
