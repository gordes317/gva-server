package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	mcs "k8s.io/metrics/pkg/client/clientset/versioned"
)

type Namespace struct {
	Name       string               `json:"namespace" `
	Status     apiv1.NamespacePhase `json:"status" `
	CreateTime metav1.Time          `json:"createTime" `
}

var config *rest.Config

func init() {

	var kubeconfig *string
	var clientErr error
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	//	k8sconfig := flag.String("k8sconfig", "/Users/wuchengjiang/Service/golang/go/packages/src/gin-vue-admin/server/kubeconfig/.k8sconfig", "kubernetes config file path")

	flag.Parse()

	// use the current context in kubeconfig
	config, clientErr = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if clientErr != nil {
		panic(clientErr.Error())
	} else {
		fmt.Println("connect k8s success")
	}

}

func testGetNode() {
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	fmt.Println(nodes.Items[0].Name)
	fmt.Println(nodes.Items[0].CreationTimestamp) //加入集群时间
	fmt.Println(nodes.Items[0].Status.NodeInfo)
	fmt.Println(nodes.Items[0].Status.Conditions[len(nodes.Items[0].Status.Conditions)-1].Type)
	fmt.Println(nodes.Items[0].Status.Allocatable.Memory().String())

	fmt.Println(nodes.Items[0].Status.Allocatable.Cpu().AsInt64())
	fmt.Println(nodes.Items[0].Status.Allocatable.Pods().AsInt64())

	// for _, node := range nodes.Items {
	// 	fmt.Printf(" Name: %s\n, Status:%s\n, nodeinfo: %s\n,nodeaddress: %s\n,capacity: %s\n,allocatable: %s\n", node.GetName(), node.Status.Conditions[len(nodes.Items[0].Status.Conditions)-1].Type, node.Status.NodeInfo, node.Status.Addresses, node.Status.Capacity.Memory().String(), node.Status.Allocatable.Memory().String())
	// }

	//获取node详情
	srv, err := clientset.CoreV1().Nodes().Get(context.TODO(), "172.16.25.15", metav1.GetOptions{})

	fmt.Printf("%v", srv)

}

func testMetrics1() {

	//podMetrics
	mc, err := mcs.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	nmList, err := mc.MetricsV1beta1().PodMetricses("").List(context.TODO(), metav1.ListOptions{})

	fmt.Println(nmList)

}

func testGetPod() {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// namespace不指定获取所有Pod 列表
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	//测试watcher

	// resourceVersion := pods.ListMeta.ResourceVersion

	// watcher, err := clientset.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{ResourceVersion: resourceVersion})

	// ch := watcher.ResultChan()

	// // LISTEN TO CHANNEL
	// for {
	// 	event := <-ch

	// 	pod, ok := event.Object.(*apiv1.Pod)
	// 	if !ok {
	// 		panic("Could not cast to Endpoint")
	// 	}
	// 	fmt.Printf("%v\n", pod.ObjectMeta.Name)

	// }

	//

	fmt.Printf("There are %v pods :\n", pods.Items[0])

	//打印一个pod的详细信息
	// fmt.Println(pods.Items[1].Name)
	// fmt.Println(pods.Items[1].CreationTimestamp)
	//fmt.Println("label:", pods.Items[1].Labels)
	// fmt.Println(pods.Items[1].Namespace)
	// fmt.Println(pods.Items[1].Status.HostIP)
	// fmt.Println(pods.Items[1].Status.PodIP)
	// fmt.Println(pods.Items[1].Status.StartTime)
	//fmt.Println("phase:", pods.Items[1].Status.Phase)
	// fmt.Println(pods.Items[1].Status.ContainerStatuses[0].RestartCount) //重启次数
	//fmt.Println("lastterminationstate:", pods.Items[1].Status.ContainerStatuses[0].LastTerminationState)
	//fmt.Println("image:", pods.Items[1].Status.ContainerStatuses[0].Image) //image

	//打印所有pod
	// for _, pod := range pods.Items {
	// 	fmt.Printf("Namespaces:%s, Name: %s, Status: %s, CreateTime: %s\n", pod.ObjectMeta.Namespace, pod.ObjectMeta.Name, pod.Status.Phase, pod.ObjectMeta.CreationTimestamp)
	// }

	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	namespace := "kube-system"
	pod := pods.Items[0].ObjectMeta.Name
	_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
			pod, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	}
	//fmt.Printf(" Name: %s\n, Kind:%s,Label:%s\n, ObjectMeta: %s\n,Spec: %s\n,podip:%s\n", podDetail.Name, podDetail.Kind, podDetail.Labels, podDetail.ObjectMeta, podDetail.Spec, podDetail.Status.PodIP)

}

func testCRUDNamespace() {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NamespacesGetter 接口方法 Namespaces 返回 NamespaceInterface
	// NamespaceInterface 接口拥有操作 Namespace 资源的方法，例如 Create、Update、Get、List 等方法
	name := "client-go-test"
	namespacesClient := clientset.CoreV1().Namespaces()

	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: apiv1.NamespaceStatus{
			Phase: apiv1.NamespaceActive,
		},
	}

	// 创建一个新的 Namespaces
	fmt.Println("Creating Namespaces...")
	result, err := namespacesClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Namespaces %s on %s\n", result.ObjectMeta.Name, result.ObjectMeta.CreationTimestamp)

	// 获取指定名称的 Namespaces 信息
	fmt.Println("Getting Namespaces...")
	result, err = namespacesClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s, Status: %s, selfLink: %s, uid: %s\n",
		result.ObjectMeta.Name, result.Status.Phase, result.ObjectMeta.SelfLink, result.ObjectMeta.UID)

	// 删除指定名称的 Namespaces 信息
	fmt.Println("Deleting Namespaces...")
	deletePolicy := metav1.DeletePropagationForeground
	if err := namespacesClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Printf("Deleted Namespaces %s\n", name)
}

func testGetListNamespace() {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespacesClient := clientset.CoreV1().Namespaces()

	result, err := namespacesClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n Namespace
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Status = ser.Status.Phase
		n.CreateTime = ser.ObjectMeta.CreationTimestamp
		listResult = append(listResult, n)

	}

	for k, v := range listResult {
		fmt.Println(k, v)
	}

}

func testCRUDDeployment() {
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientset.AppsV1().Deployments("madhouse").List(context.TODO(), metav1.ListOptions{})
	//打印所有pod
	for _, dep := range deployment.Items {
		fmt.Printf("Namespaces:%s, deployment Name: %s, replicas: %d,Ready Replicas:%d ,UpdateReplicas :%d,Available Replicas: %d, Condition: %v,CreateTime: %s\n", dep.ObjectMeta.Namespace, dep.ObjectMeta.Name, dep.Status.Replicas, dep.Status.ReadyReplicas, dep.Status.UpdatedReplicas, dep.Status.AvailableReplicas, dep.ObjectMeta.CreationTimestamp)
		fmt.Println(dep.Spec.Selector.MatchLabels)
	}

	dep, err := clientset.AppsV1().Deployments("madhouse").Get(context.TODO(), deployment.Items[0].ObjectMeta.Name, metav1.GetOptions{})

	fmt.Println(dep)

}

func testCRUDService() {
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	service, err := clientset.CoreV1().Services("madhouse").List(context.TODO(), metav1.ListOptions{})
	//打印所有service
	for _, ser := range service.Items {
		fmt.Printf("Namespaces:%s, service Name: %s, CreateTime: %s\n", ser.ObjectMeta.Namespace, ser.ObjectMeta.Name, ser.ObjectMeta.CreationTimestamp)
		fmt.Println(ser.Spec.Type, ser.Spec.Ports, ser.Spec.ClusterIP, ser.Spec.ExternalIPs, ser.Spec.LoadBalancerIP, ser.Spec.Selector)
	}
	//获取service详情
	srv, err := clientset.CoreV1().Services("madhouse").Get(context.TODO(), service.Items[0].ObjectMeta.Name, metav1.GetOptions{})

	fmt.Println(srv)
}

func testCRUDIngress() {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ing, err := clientset.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})

	//打印所有ingress
	for _, ing := range ing.Items {
		fmt.Printf("Namespaces:%s, service Name: %s, rule :%v,CreateTime: %s\n", ing.ObjectMeta.Namespace, ing.ObjectMeta.Name, ing.Spec.Rules, ing.ObjectMeta.CreationTimestamp)
		fmt.Println(ing.Spec.IngressClassName)
	}

	//打印ingress详情
	ingr, err := clientset.NetworkingV1().Ingresses("kube-system").Get(context.TODO(), ing.Items[0].ObjectMeta.Name, metav1.GetOptions{})

	fmt.Println(ingr)

}

func testCRUDreplicaSet() {
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	repliset, err := clientset.AppsV1().ReplicaSets("").List(context.TODO(), metav1.ListOptions{})

	fmt.Println(len(repliset.Items))
	//打印所有repliset
	for _, rep := range repliset.Items {
		fmt.Printf("Namespaces:%s, repliset Name: %s, Replicas:%d, Available Replicas:%d,ReadyReplicas:%d, CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Status.Replicas, rep.Status.AvailableReplicas, rep.Status.ReadyReplicas, rep.ObjectMeta.CreationTimestamp)
	}

}

//replicationController
func testCRUDReplicationController() {
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	repliset, err := clientset.CoreV1().ReplicationControllers("").List(context.TODO(), metav1.ListOptions{})
	fmt.Println(len(repliset.Items))
	//打印所有repliset
	for _, rep := range repliset.Items {
		fmt.Printf("Namespaces:%s, repliset Name: %s,replicas: %d,ready: %d, AvailableReplicas:%d,CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Status.Replicas, rep.Status.ReadyReplicas, rep.Status.AvailableReplicas, rep.ObjectMeta.CreationTimestamp)
	}

}

func testCRUDStatefulSet() {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	statefulsetList, err := clientset.AppsV1().StatefulSets("").List(context.TODO(), metav1.ListOptions{})

	fmt.Println(len(statefulsetList.Items))
	//打印所有repliset
	for _, rep := range statefulsetList.Items {
		fmt.Printf("Namespaces:%s, statefulset Name: %s, Replicas %d ,ReadyReplicas %d ,CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Status.Replicas, rep.Status.ReadyReplicas, rep.ObjectMeta.CreationTimestamp)
	}

}

func testCRUDDaemonSet() {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	daemonset, err := clientset.AppsV1().DaemonSets("").List(context.TODO(), metav1.ListOptions{})

	fmt.Println(len(daemonset.Items))
	//打印所有repliset
	for _, rep := range daemonset.Items {
		fmt.Printf("Namespaces:%s, repliset Name: %s, CurrentNumberScheduled:%d,DesiredNumberScheduled:%d,NumberReady:%d,NumberAvailable:%d,CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Status.CurrentNumberScheduled, rep.Status.DesiredNumberScheduled, rep.Status.NumberReady, rep.Status.NumberAvailable, rep.ObjectMeta.CreationTimestamp)
	}

	//打印ingress详情
	ingr, err := clientset.AppsV1().DaemonSets("kube-system").Get(context.TODO(), "calico-node", metav1.GetOptions{})
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(ingr)

}

func testCRUDJob() {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	job, err := clientset.BatchV1().Jobs("").List(context.TODO(), metav1.ListOptions{})

	//打印所有repliset
	for _, rep := range job.Items {
		fmt.Printf("Namespaces:%s, job Name: %s,Succeeded:%d, Failed:%d, CompletionTime:%s, CompletionTime:%s, CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Status.Succeeded, rep.Status.Failed, rep.Status.CompletionTime, rep.Status.CompletionTime, rep.ObjectMeta.CreationTimestamp)
	}

}

func testCRUDCronJob() {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	cronjob, err := clientset.BatchV1beta1().CronJobs("").List(context.TODO(), metav1.ListOptions{})

	//打印所有repliset
	for _, rep := range cronjob.Items {
		fmt.Printf("Namespaces:%s, cron job Name: %s,Schedule:%s,Suspend:%v,Active:%v,LastScheduleTime:%s,CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Spec.Schedule, *rep.Spec.Suspend, rep.Status.Active, rep.Status.LastScheduleTime, rep.ObjectMeta.CreationTimestamp)

	}

}

func testCURDEndpoints() {
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	endpoint, err := clientset.CoreV1().Endpoints("").List(context.TODO(), metav1.ListOptions{})

	//打印所有repliset
	for _, rep := range endpoint.Items {
		fmt.Printf("Namespaces:%s, endpoint Name: %s,subsets:%v,CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Subsets, rep.ObjectMeta.CreationTimestamp)
	}

}

func testCRUDNetworkPolicy() {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	networkpolicy, err := clientset.NetworkingV1().NetworkPolicies("").List(context.TODO(), metav1.ListOptions{})

	//打印所有repliset
	for _, rep := range networkpolicy.Items {
		fmt.Printf("Namespaces:%s, network policy Name: %s,CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.ObjectMeta.CreationTimestamp)
	}

}

func testCRUDPV() {

	//create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pv, err := clientset.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	//打印所有storageclass
	for _, rep := range pv.Items {
		fmt.Printf("Namespaces:%s, PV Name: %s, AccessMode:%s,Capacity:%v,PersistentVolumeSource.NFS.Server:%s,PersistentVolumeReclaimPolicy:%s,Phase:%s ,StorageClass:%s,CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Spec.AccessModes, rep.Spec.Capacity, rep.Spec.PersistentVolumeSource.NFS.Server, rep.Spec.PersistentVolumeReclaimPolicy, rep.Status.Phase, rep.Spec.StorageClassName, rep.ObjectMeta.CreationTimestamp)
	}

	// pvDetail, err := clientset.CoreV1().PersistentVolumes().Get(context.TODO(), "mysql-pv", metav1.GetOptions{})
	// fmt.Print(pvDetail)

}

func testCRUDPVC() {

	// // create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pvc, err := clientset.CoreV1().PersistentVolumeClaims("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	//打印所有storageclass
	for _, rep := range pvc.Items {
		fmt.Printf("Namespaces:%s, PVC Name: %s, Status:%s,AccessMode:%s,Capacity:%v,CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Status.Phase, rep.Status.AccessModes, rep.Status.Capacity, rep.ObjectMeta.CreationTimestamp)
	}

	// pvcDetail, err := clientset.CoreV1().PersistentVolumeClaims("elk").Get(context.TODO(), "es-data-es-0", metav1.GetOptions{})
	// fmt.Print(pvcDetail)
}

func testStorageClass() {

	//create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	sc, err := clientset.StorageV1().StorageClasses().List(context.TODO(), metav1.ListOptions{})
	//打印所有storageclass
	for _, rep := range sc.Items {
		fmt.Printf("Namespaces:%s, storageClass Name: %s, Provisioner:%s,ReclaimPolicy:%v,VolumeBindingMode:%v,AllowVolumeExpansion:%s ,CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Provisioner, *rep.ReclaimPolicy, *rep.VolumeBindingMode, rep.AllowVolumeExpansion, rep.ObjectMeta.CreationTimestamp)
	}
	// 指针类型打印问题
	// scDetail, err := clientset.StorageV1().StorageClasses().Get(context.TODO(), "managed-nfs-storage", metav1.GetOptions{})
	// fmt.Print(scDetail)
}

func testCRUDConfigMap() {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	config, err := clientset.CoreV1().ConfigMaps("").List(context.TODO(), metav1.ListOptions{})

	for _, rep := range config.Items {
		fmt.Printf("Namespaces:%s, config Name: %s, data length:%d, CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, len(rep.Data), rep.ObjectMeta.CreationTimestamp)
	}

}

func testCRUDSecret() {
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	secret, err := clientset.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})

	for _, rep := range secret.Items {
		fmt.Printf("Namespaces:%s, secret Name: %s,type:%s, data length:%d, CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Type, len(rep.Data), rep.ObjectMeta.CreationTimestamp)
	}
}

func testCRUDRBAC() {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	role, err := clientset.RbacV1().Roles("").List(context.TODO(), metav1.ListOptions{})

	for _, rep := range role.Items {
		fmt.Printf("Namespaces:%s, role Name: %s,  CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.ObjectMeta.CreationTimestamp)
	}

	clusterrole, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})

	for _, rep := range clusterrole.Items {
		fmt.Printf("Namespaces:%s, clusterrole  Name: %s,  CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.ObjectMeta.CreationTimestamp)
	}

	rolebinding, err := clientset.RbacV1().RoleBindings("").List(context.TODO(), metav1.ListOptions{})

	for _, rep := range rolebinding.Items {
		fmt.Printf("Namespaces:%s, rolebinding Name: %s, RoleName:%s,  CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.RoleRef.Name, rep.ObjectMeta.CreationTimestamp)
	}

	clusterrolebinding, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})

	for _, rep := range clusterrolebinding.Items {
		fmt.Printf("Namespaces:%s, clusterrolebinding Name: %s, RoleName:%s, CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.RoleRef.Name, rep.ObjectMeta.CreationTimestamp)
	}

}

func testCRUDResourcequotas() {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	resourcequotas, err := clientset.CoreV1().ResourceQuotas("").List(context.TODO(), metav1.ListOptions{})

	for _, rep := range resourcequotas.Items {
		fmt.Printf("Namespaces:%s, config Name: %s,  CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.ObjectMeta.CreationTimestamp)
	}

}

func testCRUDEvent() {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	event, err := clientset.EventsV1().Events("").List(context.TODO(), metav1.ListOptions{})

	for _, rep := range event.Items {
		fmt.Printf("Namespaces:%s, event Name: %s, Reason:%s, Type:%s, Note Message:%s, DeprecatedLastTimestamp:%s, CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.Reason, rep.Type, rep.Note, rep.DeprecatedLastTimestamp, rep.ObjectMeta.CreationTimestamp)
	}

}

func testMetrics() {
	mc, err := mcs.NewForConfig(config)

	if err != nil {
		panic(err)
	}

	podMetrics, err := mc.MetricsV1beta1().PodMetricses("").List(context.TODO(), metav1.ListOptions{})
	//podMetrics, err := mc.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})

	for _, rep := range podMetrics.Items {
		fmt.Printf("metrics Namespaces:%s, metrics Name: %s, CreateTime: %s\n", rep.ObjectMeta.Namespace, rep.ObjectMeta.Name, rep.ObjectMeta.CreationTimestamp)
	}

}
func testLogs() {

	// var (
	// 	clientset *kubernetes.Clientset
	var req *rest.Request
	// 	err error
	// )

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	opts := &apiv1.PodLogOptions{
		Follow: true,
	}

	// 生成获取POD日志请求

	//req = clientset.CoreV1().Pods("cdi").GetLogs("cdi-apiserver-69f7f84fcf-snbpr", &apiv1.PodLogOptions{Container: "cdi-apiserver", TailLines: &tailLines})

	//req = clientset.CoreV1().Pods("cdi").GetLogs("cdi-apiserver-69f7f84fcf-snbpr", opts)

	req = clientset.CoreV1().Pods("k8s-test").GetLogs("nginx-deployment-585449566-42sdx", opts)

	// req.Stream()也可以实现Do的效果

	// 发送请求

	readCloser, err := req.Stream(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}

	defer readCloser.Close()

	// 获取结果

	r := bufio.NewReader(readCloser)

	for {
		bytes, err := r.ReadBytes('\n')
		fmt.Println(string(bytes))
		if err != nil {
			if err != io.EOF {
				goto FAIL
			}
			return
		}
	}

FAIL:
	fmt.Println(err)
	return

}

func main() {
	//测试K8s连接

	// 测试GetNode
	//testGetNode()

	//Metrics
	//testMetrics1()

	//测试pod
	//testGetPod()

	// 测试k8s namespace创建和删除
	//testCRUDNamespace()

	//testGetListNamespace()

	// //测试deployment
	//testCRUDDeployment()

	// //测试service
	//testCRUDService()

	// //测试ingress
	//testCRUDIngress()

	// //测试ReplicaSet
	//testCRUDreplicaSet()

	//测试replicationController
	//testCRUDReplicationController()

	// //测试statefulset
	//testCRUDStatefulSet()

	//测试DaemonSet
	//testCRUDDaemonSet()

	//测试configmap
	//testCRUDConfigMap()

	//测试Secret
	//testCRUDSecret()

	//测试PV
	//testCRUDPV()

	//测试PVC
	//testCRUDPVC()

	//测试StorageClass
	//testStorageClass()

	//测试Job
	//testCRUDJob()

	//测试CronJob
	//testCRUDCronJob()

	//测试endpoint
	//testCURDEndpoints()

	//测试networkpolicy
	//testCRUDNetworkPolicy()

	//测试RBAC
	//testCRUDRBAC()

	//测试Resourcequotas
	//testCRUDResourcequotas()

	//测试event
	//testCRUDEvent()

	// //
	//testMetrics()

	//testLogs()
}
