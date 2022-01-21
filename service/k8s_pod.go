package service

import (
	"bytes"
	"context"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"io"
	"log"
	"strings"

	ghodssyaml "github.com/ghodss/yaml"
	apiv1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePod
//@description: 创建Pod记录
//@param: Pod model.Pod
//@return: err error

func CreatePod(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods 返回 PodInterface
	// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法
	podsClient := clientset.CoreV1().Pods("")

	namespace := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: apiv1.PodStatus{
			Phase: apiv1.PodSucceeded,
		},
	}

	_, err = podsClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePod
//@description: 删除Pod记录
//@param: Pod model.Pod
//@return: err error

func DeletePod(name string, namespace string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods 返回 PodInterface
	// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法
	fmt.Println("namespace:", namespace)
	podsClient := clientset.CoreV1().Pods(namespace)

	deletePolicy := metav1.DeletePropagationForeground
	if err := podsClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePodByIds
//@description: 批量删除Pod记录
//@param: ids request.IdsReq
//@return: err error

func DeletePodByIds(names request.NamesReqUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(names.UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods 返回 PodInterface
	// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法
	//podsClient := clientset.CoreV1().Pods("")

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		name := strings.Split(value, ";")[0]
		namespace := strings.Split(value, ";")[1]
		podsClient := clientset.CoreV1().Pods(namespace)
		if err := podsClient.Delete(context.TODO(), name, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePod
//@description: 更新Pod记录
//@param: Pod *model.Pod
//@return: err error

func UpdatePod(Pod model.PodUser, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: Pod.Name,
		},
		Status: apiv1.PodStatus{
			Phase: apiv1.PodPhase(Pod.Phase),
		},
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods 返回 PodInterface
	// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法
	podsClient := clientset.CoreV1().Pods("")
	_, err = podsClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPod
//@description: 根据name获取Pod记录
//@param: name uint
//@return: err error, Pod model.Pod

func GetPod(name string, UserName string) (err error, Pod model.Pod) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods 返回 PodInterface
	// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法
	podsClient := clientset.CoreV1().Pods("")

	result, err := podsClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	Pod.Name = result.Name
	Pod.Phase = string(result.Status.Phase)
	Pod.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, Pod
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPodInfoList
//@description: 分页获取Pod记录
//@param: info request.PodSearch
//@return: err error, list interface{}, total int64

func GetPodInfoList(info request.PodSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	podsClient := clientset.CoreV1().Pods("")

	result, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	//podMetrics
	mc, err := metrics.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err)
	}

	podMetrics, err := mc.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var mcm model.MetricsCpuMem

	podMetricList := make([]interface{}, 0)

	for _, podMetric := range podMetrics.Items {

		metadataname := podMetric.ObjectMeta.Name
		mcm.Name = metadataname

		podContainers := podMetric.Containers
		for _, container := range podContainers {
			cpuQuantity := container.Usage.Cpu().ToDec().AsApproximateFloat64()
			memQuantity, ok := container.Usage.Memory().AsInt64()
			if !ok {
				return
			}
			mcm.Cpu = float32(cpuQuantity)
			mcm.Mem = float32(memQuantity / 1024 / 1024)
		}

		podMetricList = append(podMetricList, mcm)

	}

	var n model.Pod

	listResult := make([]interface{}, 0)

	var container model.Container

	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Phase = string(ser.Status.Phase)
		//n.Phase = ser.Status.Phase
		n.CreateTime = ser.ObjectMeta.CreationTimestamp.Time
		n.Namespace = ser.Namespace
		n.HostIP = ser.Status.HostIP
		n.PodIP = ser.Status.PodIP
		//n.RestartCount = int32(ser.Status.ContainerStatuses[0].RestartCount)
		//n.Image = ser.Status.ContainerStatuses[0].Image
		//n.ContainerName = ser.Spec.Containers[0].Name

		//重新生成container name 和 container image的map
		containerList := make([]model.Container, 0)
		for _, ct := range ser.Spec.Containers {
			container.CName = ct.Name
			container.CImage = ct.Image

			//容器重启次数BUG修复
			for _, v := range ser.Status.ContainerStatuses {
				if ct.Name == v.Name {
					container.CRestartCount = v.RestartCount
					//container.CReady = v.Ready
					//ser.Status.Reason
				}
			}

			containerList = append(containerList, container)
		}
		n.Containers = containerList

		//遍历podMetricList 获取 pod cpu mem
		for _, mapVal := range podMetricList {

			metricsMap := utils.StructToMap(mapVal)

			if metricsMap["Name"] == ser.Name {
				n.CPU = metricsMap["Cpu"].(float32)
				n.Mem = metricsMap["Mem"].(float32)
			}

		}

		//按照条件搜索判断条件
		if strings.Contains(n.Name, info.Name) && strings.Contains(n.Namespace, info.Namespace) {
			listResult = append(listResult, n)
		}

	}

	total = int64(len(listResult))

	// 分页方法
	//list, err = utils.Paginator(info.Page, info.PageSize, listResult)

	return err, listResult, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPodLog
//@description: 获取pod日志方法
//@param: name uint
//@return: err error, Pod model.Pod

func GetPodLog(Name string, Namespace string, ContainerName string, UserName string) (err error, Pod model.Pod) {

	var req *rest.Request
	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	opts := &apiv1.PodLogOptions{
		//Follow:    true,
		Container: ContainerName,
	}

	req = clientset.CoreV1().Pods(Namespace).GetLogs(Name, opts)
	c, _ := clientset.CoreV1().Pods(Namespace).Get(context.TODO(), Name, metav1.GetOptions{})
	var reason string
	for _, ct := range c.Status.ContainerStatuses {
		if ct.Name == ContainerName {
			if ct.State.Terminated != nil {
				reason = ct.State.Terminated.Reason
			} else {
				reason = "Running"
			}
			Pod.Phase = reason
			break
		}
	}
	Pod.Status = string(c.Status.Phase)

	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}

	defer podLogs.Close()

	// 获取结果

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		fmt.Println(err.Error())
	}
	strLogs := buf.String()

	Pod.Name = Name
	Pod.Namespace = Namespace
	//Pod.ContainerName = ContainerName

	Pod.Log = strLogs
	//Pod.Ready = req.

	return err, Pod
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ApplyYamlPod
//@description: 更新Pod Yaml记录
//@param: Pod *model.Pod
//@return: err error

func ApplyYamlPod(Pod model.PodUser, UserName string) (err error) {

	//动态client
	dynamicClient, err := dynamic.NewForConfig(global.GetK8sConfig(UserName))

	if err != nil {
		log.Fatal(err)
	}

	var data []byte = []byte(Pod.YamlData) // string to byte

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(global.GetK8sConfig(UserName))

	if err != nil {
		log.Fatal(discoveryClient)
	}
	applyOptions := utils.NewApplyOptions(dynamicClient, discoveryClient)

	if err := applyOptions.Apply(context.TODO(), data); err != nil {
		log.Fatalf("apply error: %v", err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPod
//@description: 根据name获取Pod记录
//@param: name uint
//@return: err error, Pod model.Pod

func ReadYamlPod(name string, namespace string, UserName string) (err error, Pod model.Pod) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	//for _, deploy := range deployment.Items {
	y, err := ghodssyaml.Marshal(deployment)
	if err != nil {
		panic(err.Error)
	}
	//}

	yamldata := fmt.Sprintf("apiVersion: v1\nkind: Pod\n%s", string(y))

	Pod.Name = name
	Pod.Namespace = namespace
	Pod.YamlData = yamldata

	return err, Pod
}
