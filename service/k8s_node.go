package service

import (
	"context"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"strings"

	apiv1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"

	"k8s.io/client-go/kubernetes"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateNode
//@description: 创建Node记录
//@param: Node model.Node
//@return: err error

func CreateNode(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NodesGetter 接口方法 Nodes 返回 NodeInterface
	// NodeInterface 接口拥有操作 Node 资源的方法，例如 Create、Update、Get、List 等方法
	podsClient := clientset.CoreV1().Nodes()

	node := &apiv1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: apiv1.NodeStatus{
			Phase: apiv1.NodeRunning,
		},
	}

	_, err = podsClient.Create(context.TODO(), node, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNode
//@description: 删除Node记录
//@param: Node model.Node
//@return: err error

func DeleteNode(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NodesGetter 接口方法 Nodes 返回 NodeInterface
	// NodeInterface 接口拥有操作 Node 资源的方法，例如 Create、Update、Get、List 等方法
	podsClient := clientset.CoreV1().Nodes()

	deletePolicy := metav1.DeletePropagationForeground
	if err := podsClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNodeByIds
//@description: 批量删除Node记录
//@param: ids request.IdsReq
//@return: err error

func DeleteNodeByIds(names request.NamesReqUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(names.UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NodesGetter 接口方法 Nodes 返回 NodeInterface
	// NodeInterface 接口拥有操作 Node 资源的方法，例如 Create、Update、Get、List 等方法
	podsClient := clientset.CoreV1().Nodes()

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		if err := podsClient.Delete(context.TODO(), value, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateNode
//@description: 更新Node记录
//@param: Node *model.Node
//@return: err error

func UpdateNode(Node model.NodeUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(Node.UserName))
	if err != nil {
		panic(err.Error())
	}

	node := &apiv1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: Node.Name,
		},
		Status: apiv1.NodeStatus{
			Phase: apiv1.NodePhase(Node.Status),
		},
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NodesGetter 接口方法 Nodes 返回 NodeInterface
	// NodeInterface 接口拥有操作 Node 资源的方法，例如 Create、Update、Get、List 等方法
	podsClient := clientset.CoreV1().Nodes()
	_, err = podsClient.Update(context.TODO(), node, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNode
//@description: 根据name获取Node记录
//@param: name uint
//@return: err error, Node model.Node

func GetNode(name string, userName string) (err error, Node model.Node) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(userName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NodesGetter 接口方法 Nodes 返回 NodeInterface
	// NodeInterface 接口拥有操作 Node 资源的方法，例如 Create、Update、Get、List 等方法
	podsClient := clientset.CoreV1().Nodes()

	result, err := podsClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	Node.Name = result.Name
	Node.Status = string(result.Status.Phase)
	Node.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, Node
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNodeInfoList
//@description: 分页获取Node记录
//@param: info request.NodeSearch
//@return: err error, list interface{}, total int64

func GetNodeInfoList(info request.NodeSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	podsClient := clientset.CoreV1().Nodes()

	result, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	//get node Metrics usage

	var mcm model.MetricsCpuMem

	mc, err := metrics.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err)
	}
	nodeMetrics, err := mc.MetricsV1beta1().NodeMetricses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	nodeMetricList := make([]interface{}, 0)

	for _, nodeMetric := range nodeMetrics.Items {

		mcm.Name = nodeMetric.ObjectMeta.Name

		cpuQuantity := nodeMetric.Usage.Cpu().ToDec().AsApproximateFloat64()
		memQuantity, ok := nodeMetric.Usage.Memory().AsInt64()

		if !ok {
			return
		}
		mcm.Cpu = float32(cpuQuantity)
		mcm.Mem = float32(memQuantity / 1024 / 1024)

		nodeMetricList = append(nodeMetricList, mcm)

	}

	//

	listResult := make([]interface{}, 0)

	var n model.Node
	for _, ser := range result.Items {
		n.Name = ser.Name

		var ips string
		for _, val := range ser.Status.Addresses {
			if val.Type == "InternalIP" {
				ips = val.Address
			}
		}
		n.IP = ips

		n.Status = string(ser.Status.Conditions[len(result.Items[0].Status.Conditions)-1].Type)
		n.KernelVersion = ser.Status.NodeInfo.KernelVersion
		n.NodeSystem = ser.Status.NodeInfo.OSImage
		n.Version = ser.Status.NodeInfo.KubeletVersion

		n.RunTime = ser.Status.NodeInfo.ContainerRuntimeVersion

		n.Role = ser.ObjectMeta.GetLabels()["kubernetes.io/role"]

		// podnum遍历 每个pod 默认 110个pod
		podNumMap := getPodNumPerNode(info.UserName)
		for k, v := range podNumMap {
			if k == n.IP {
				n.PodsNumber = v
			}
		}

		n.CpuTotal = float32(ser.Status.Capacity.Cpu().Value())

		n.MemTotal = float32(ser.Status.Capacity.Memory().Value() / 1024 / 1024)

		//遍历node MetricList 获取 node usage cpu mem
		for _, mapVal := range nodeMetricList {

			metricsMap := utils.StructToMap(mapVal)

			if metricsMap["Name"] == ser.Name {
				n.CpuAllocatable = metricsMap["Cpu"].(float32)
				n.MemAllocatable = metricsMap["Mem"].(float32)
			}

		}
		//node metrics end
		m := make(map[string]string)
		for _, val := range ser.Spec.Taints {
			m[val.Key] = val.Value

		}
		n.Taint = utils.MapToJson(m)

		n.CreateTime = ser.ObjectMeta.CreationTimestamp.Time

		//按照条件搜索判断条件
		if strings.Contains(n.IP, info.IP) {
			listResult = append(listResult, n)
		}

	}

	total = int64(len(listResult))

	// 分页方法
	list, err = utils.Paginator(info.Page, info.PageSize, listResult)

	return err, list, total
}

func getPodNumPerNode(userName string) map[string]int32 {

	m := make(map[string]int32)

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(userName))
	if err != nil {
		panic(err.Error())
	}

	podsClient := clientset.CoreV1().Pods("")

	result, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	k := make([]string, 0)
	for _, ser := range result.Items {
		k = append(k, ser.Status.HostIP)
	}

	for _, v := range k {
		m[v] += 1
	}

	return m
}
