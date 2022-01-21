package service

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"log"
	"strings"

	ghodssyaml "github.com/ghodss/yaml"
	apiv1 "k8s.io/api/apps/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDaemonSet
//@description: 创建DaemonSet记录
//@param: DaemonSet model.DaemonSet
//@return: err error

func CreateDaemonSet(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 DaemonSetsGetter 接口方法 DaemonSets 返回 DaemonSetInterface
	// DaemonSetInterface 接口拥有操作 DaemonSet 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.AppsV1().DaemonSets("")

	namespace := &apiv1.DaemonSet{}

	_, err = secretClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDaemonSet
//@description: 删除DaemonSet记录
//@param: DaemonSet model.DaemonSet
//@return: err error

func DeleteDaemonSet(name string, namespace string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 DaemonSetsGetter 接口方法 DaemonSets 返回 DaemonSetInterface
	// DaemonSetInterface 接口拥有操作 DaemonSet 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.AppsV1().DaemonSets(namespace)

	deletePolicy := metav1.DeletePropagationForeground
	if err := secretClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDaemonSetByIds
//@description: 批量删除DaemonSet记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDaemonSetByIds(names request.NamesReqUser, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 DaemonSetsGetter 接口方法 DaemonSets 返回 DaemonSetInterface
	// DaemonSetInterface 接口拥有操作 DaemonSet 资源的方法，例如 Create、Update、Get、List 等方法
	//secretClient := clientset.AppsV1().DaemonSets("")

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		name := strings.Split(value, ";")[0]
		namespace := strings.Split(value, ";")[1]
		secretClient := clientset.AppsV1().DaemonSets(namespace)
		if err := secretClient.Delete(context.TODO(), name, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDaemonSet
//@description: 更新DaemonSet记录
//@param: DaemonSet *model.DaemonSet
//@return: err error

func UpdateDaemonSet(DaemonSet model.DaemonSetUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(DaemonSet.UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &apiv1.DaemonSet{}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 DaemonSetsGetter 接口方法 DaemonSets 返回 DaemonSetInterface
	// DaemonSetInterface 接口拥有操作 DaemonSet 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.AppsV1().DaemonSets("")
	_, err = secretClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDaemonSet
//@description: 根据name获取DaemonSet记录
//@param: name uint
//@return: err error, DaemonSet model.DaemonSet

func GetDaemonSet(name string, namespace string, UserName string) (err error, DaemonSet model.DaemonSet) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 DaemonSetsGetter 接口方法 DaemonSets 返回 DaemonSetInterface
	// DaemonSetInterface 接口拥有操作 DaemonSet 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.AppsV1().DaemonSets(namespace)

	result, err := secretClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	DaemonSet.Name = result.Name
	DaemonSet.Namespace = result.Namespace

	//DaemonSet.Data = result.Data
	DaemonSet.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, DaemonSet
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDaemonSetInfoList
//@description: 分页获取DaemonSet记录
//@param: info request.DaemonSetSearch
//@return: err error, list interface{}, total int64

func GetDaemonSetInfoList(info request.DaemonSetSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	secretClient := clientset.AppsV1().DaemonSets("")

	result, err := secretClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n model.DaemonSet
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Namespace = ser.Namespace
		n.Desired = ser.Status.DesiredNumberScheduled
		n.Current = ser.Status.CurrentNumberScheduled
		n.Ready = ser.Status.NumberReady
		n.Updated = ser.Status.UpdatedNumberScheduled
		n.Available = ser.Status.NumberAvailable

		//json to string
		ns, err := json.Marshal(ser.Spec.Template.Spec.NodeSelector)
		if err != nil {
			panic(err)
		}
		n.NodeSelector = string(ns)
		n.CreateTime = ser.ObjectMeta.CreationTimestamp.Time

		//按照条件搜索判断条件
		if strings.Contains(n.Name, info.Name) {
			listResult = append(listResult, n)
		}

	}

	total = int64(len(listResult))

	// 分页方法
	list, err = utils.Paginator(info.Page, info.PageSize, listResult)

	return err, list, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ApplyYamlDaemonSet
//@description: 更新DaemonSet Yaml记录
//@param: DaemonSet *model.DaemonSet
//@return: err error

func ApplyYamlDaemonSet(DaemonSet model.DaemonSetUser) (err error) {

	//动态client
	dynamicClient, err := dynamic.NewForConfig(global.GetK8sConfig(DaemonSet.UserName))

	if err != nil {
		log.Fatal(err)
	}

	var data []byte = []byte(DaemonSet.YamlData) // string to byte

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(global.GetK8sConfig(DaemonSet.UserName))

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
//@function: GetDaemonSet
//@description: 根据name获取DaemonSet记录
//@param: name uint
//@return: err error, DaemonSet model.DaemonSet

func ReadYamlDaemonSet(name string, namespace string, UserName string) (err error, DaemonSet model.DaemonSet) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientset.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	//for _, deploy := range deployment.Items {
	y, err := ghodssyaml.Marshal(deployment)
	if err != nil {
		panic(err.Error)
	}
	//}

	yamldata := fmt.Sprintf("apiVersion: apps/v1\nkind: DaemonSet\n%s", string(y))

	DaemonSet.Name = name
	DaemonSet.Namespace = namespace
	DaemonSet.YamlData = yamldata

	return err, DaemonSet
}
