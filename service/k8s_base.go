package service

import (
	"context"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"sort"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func GetSelectNamespaceList(UserName string) (err error, list interface{}) {

	// create the clientset
	var clientset *kubernetes.Clientset
	//var err error

	clientset, err = kubernetes.NewForConfig(global.GetK8sConfig(UserName))

	if err != nil {
		panic(err.Error())
	}

	namespacesClient := clientset.CoreV1().Namespaces()

	result, err := namespacesClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var sn model.SelectNamespaces
	for _, ser := range result.Items {
		sn.Value = ser.Name
		sn.Label = ser.Name
		listResult = append(listResult, sn)
	}

	return err, listResult
}

func GetSelectClusterList() (err error, list interface{}) {

	db := global.GVA_DB.Model(&model.MultiCluster{})
	var clusters []model.SelectCluster
	err = db.Find(&clusters).Error
	return err, clusters

}

func GetAllClusterInfo() (err error, list interface{}) {

	db := global.GVA_DB.Model(&model.MultiCluster{})
	var clusters []model.GetAllClusterInfo
	err = db.Find(&clusters).Error
	return err, clusters
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ApplyYamlCreate
//@description: 从 Yaml 文件Create对象
//@param: yamldata
//@return: err error

func ApplyYamlCreate(yamldata string, userName string) (err error) {

	//动态client
	dynamicClient, err := dynamic.NewForConfig(global.GetK8sConfig(userName))

	if err != nil {
		log.Fatal(err)
	}

	var data []byte = []byte(yamldata) // string to byte

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(global.GetK8sConfig(userName))

	if err != nil {
		log.Fatal(discoveryClient)
	}
	applyOptions := utils.NewApplyOptions(dynamicClient, discoveryClient)

	if err := applyOptions.Apply(context.TODO(), data); err != nil {
		log.Fatalf("apply error: %v", err)
	}
	return err
}

func GetClusterApplicationCounter(pageInfo request.PageInfoUser) (err error, clusterAppCounter map[string]int64) {

	var clusterAppCounterMap map[string]int64
	clusterAppCounterMap = make(map[string]int64)
	namespaceList := request.NamespaceList{
		User: model.User{UserName: pageInfo.UserName},
		NamespaceSearch: request.NamespaceSearch{
			PageInfo:  pageInfo.PageInfo,
			Namespace: model.Namespace{},
		},
	}
	if err, _, nameSpaceTotal := GetNamespaceInfoList(namespaceList); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
	} else {
		clusterAppCounterMap["namespace"] = nameSpaceTotal
	}
	nodeList := request.NodeSearchUser{
		NodeSearch: request.NodeSearch{
			Node:     model.Node{},
			PageInfo: pageInfo.PageInfo,
		},
		User: model.User{UserName: pageInfo.UserName},
	}
	if err, _, nodeTotal := GetNodeInfoList(nodeList); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		//response.FailWithMessage("获取失败", c)
	} else {
		clusterAppCounterMap["node"] = nodeTotal
	}
	deploymentList := request.DeploymentSearchUser{
		DeploymentSearch: request.DeploymentSearch{
			Deployment: model.Deployment{},
			PageInfo:   pageInfo.PageInfo,
		},
		User: model.User{UserName: pageInfo.UserName},
	}
	if err, _, deployTotal := GetDeploymentInfoList(deploymentList); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
	} else {
		clusterAppCounterMap["deployment"] = deployTotal
	}

	ingressList := request.IngressSearchUser{
		IngressSearch: request.IngressSearch{},
		User:          model.User{UserName: pageInfo.UserName},
	}
	kubeVersion := utils.GetKubeVersion(pageInfo.UserName)
	if kubeVersion >= 119 {
		if err, _, ingressTotal := GetIngressInfoList(ingressList); err != nil {
			global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		} else {
			clusterAppCounterMap["ingress"] = ingressTotal
		}
	} else {
		if err, _, ingressTotal := GetIngressInfoListUnder119(ingressList); err != nil {
			global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		} else {
			clusterAppCounterMap["ingress"] = ingressTotal
		}
	}

	pvcList := request.PVCSearchUser{
		PVCSearch: request.PVCSearch{},
		User:      model.User{UserName: pageInfo.UserName},
	}
	if err, _, pvcTotal := GetPVCInfoList(pvcList); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
	} else {
		clusterAppCounterMap["pvc"] = pvcTotal
	}

	pvList := request.PVSearchUser{
		PVSearch: request.PVSearch{},
		User:     model.User{UserName: pageInfo.UserName},
	}
	if err, _, pvcTotal := GetPVInfoList(pvList); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
	} else {
		clusterAppCounterMap["pv"] = pvcTotal
	}

	statefulSetList := request.StatefulSetSearchUser{
		StatefulSetSearch: request.StatefulSetSearch{},
		User:              model.User{UserName: pageInfo.UserName},
	}
	if err, _, statefulSetTotal := GetStatefulSetInfoList(statefulSetList); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
	} else {
		clusterAppCounterMap["statefulSet"] = statefulSetTotal
	}

	daemonSetList := request.DaemonSetSearchUser{
		DaemonSetSearch: request.DaemonSetSearch{},
		User:            model.User{UserName: pageInfo.UserName},
	}
	if err, _, daemonSetTotal := GetDaemonSetInfoList(daemonSetList); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
	} else {
		clusterAppCounterMap["daemonSet"] = daemonSetTotal
	}

	serviceList := request.ServiceSearchUser{
		ServiceSearch: request.ServiceSearch{},
		User:          model.User{UserName: pageInfo.UserName},
	}
	if err, _, serviceTotal := GetServiceInfoList(serviceList); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
	} else {
		clusterAppCounterMap["service"] = serviceTotal
	}

	jobList := request.JobSearchUser{
		JobSearch: request.JobSearch{},
		User:      model.User{UserName: pageInfo.UserName},
	}
	if err, _, jobTotal := GetJobInfoList(jobList); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
	} else {
		clusterAppCounterMap["job"] = jobTotal
	}

	//GetNamespaceInfoList()
	return err, clusterAppCounterMap
}

func GetClusterEvents(pageInfo request.PageInfoUser) (err error, clusterEvents interface{}, total int64) {
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(pageInfo.UserName))
	if err != nil {
		panic(err.Error())
	}
	events, _ := clientset.CoreV1().Events("").List(context.TODO(), metav1.ListOptions{})
	listResult := make([]interface{}, 0)

	var sn model.ClusterEvent
	for _, ser := range events.Items {
		sn.Reason = ser.Reason
		sn.Namespace = ser.Namespace
		sn.Message = ser.Message
		sn.Resource = ser.InvolvedObject.GetObjectKind().GroupVersionKind().Kind + "/" + ser.Name
		sn.LastEventTime = ser.CreationTimestamp
		listResult = append(listResult, sn)
	}
	sort.Sort(utils.MapSlice(listResult))
	total = int64(len(listResult))

	// 分页方法
	list, err := utils.Paginator(pageInfo.Page, pageInfo.PageSize, listResult)

	return err, list, total
}
