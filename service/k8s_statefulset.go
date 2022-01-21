package service

import (
	"context"
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
//@function: CreateStatefulSet
//@description: 创建StatefulSet记录
//@param: StatefulSet model.StatefulSet
//@return: err error

func CreateStatefulSet(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 StatefulSetsGetter 接口方法 StatefulSets 返回 StatefulSetInterface
	// StatefulSetInterface 接口拥有操作 StatefulSet 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.AppsV1().StatefulSets("")

	namespace := &apiv1.StatefulSet{}

	_, err = secretClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStatefulSet
//@description: 删除StatefulSet记录
//@param: StatefulSet model.StatefulSet
//@return: err error

func DeleteStatefulSet(name, namespace, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 StatefulSetsGetter 接口方法 StatefulSets 返回 StatefulSetInterface
	// StatefulSetInterface 接口拥有操作 StatefulSet 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.AppsV1().StatefulSets(namespace)

	deletePolicy := metav1.DeletePropagationForeground
	if err := secretClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStatefulSetByIds
//@description: 批量删除StatefulSet记录
//@param: ids request.IdsReq
//@return: err error

func DeleteStatefulSetByIds(names request.NamesReqUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(names.UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 StatefulSetsGetter 接口方法 StatefulSets 返回 StatefulSetInterface
	// StatefulSetInterface 接口拥有操作 StatefulSet 资源的方法，例如 Create、Update、Get、List 等方法
	//secretClient := clientset.AppsV1().StatefulSets("")

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		name := strings.Split(value, ";")[0]
		namespace := strings.Split(value, ";")[1]
		secretClient := clientset.AppsV1().StatefulSets(namespace)
		if err := secretClient.Delete(context.TODO(), name, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateStatefulSet
//@description: 更新StatefulSet记录
//@param: StatefulSet *model.StatefulSet
//@return: err error

func UpdateStatefulSet(StatefulSet model.StatefulSetUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(StatefulSet.UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &apiv1.StatefulSet{}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 StatefulSetsGetter 接口方法 StatefulSets 返回 StatefulSetInterface
	// StatefulSetInterface 接口拥有操作 StatefulSet 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.AppsV1().StatefulSets("")
	_, err = secretClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStatefulSet
//@description: 根据name获取StatefulSet记录
//@param: name uint
//@return: err error, StatefulSet model.StatefulSet

func GetStatefulSet(name string, namespace string, UserName string) (err error, StatefulSet model.StatefulSet) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 StatefulSetsGetter 接口方法 StatefulSets 返回 StatefulSetInterface
	// StatefulSetInterface 接口拥有操作 StatefulSet 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.AppsV1().StatefulSets(namespace)

	result, err := secretClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	StatefulSet.Name = result.Name
	StatefulSet.Namespace = result.Namespace

	StatefulSet.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, StatefulSet
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStatefulSetInfoList
//@description: 分页获取StatefulSet记录
//@param: info request.StatefulSetSearch
//@return: err error, list interface{}, total int64

func GetStatefulSetInfoList(info request.StatefulSetSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))

	if err != nil {
		panic(err.Error())
	}

	secretClient := clientset.AppsV1().StatefulSets("")

	result, err := secretClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n model.StatefulSet
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Namespace = ser.Namespace
		n.Replicas = ser.Status.Replicas
		n.Ready = ser.Status.ReadyReplicas
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
//@function: ApplyYamlStatefulSet
//@description: 更新StatefulSet Yaml记录
//@param: StatefulSet *model.StatefulSet
//@return: err error

func ApplyYamlStatefulSet(StatefulSet model.StatefulSetUser) (err error) {

	//动态client
	dynamicClient, err := dynamic.NewForConfig(global.GetK8sConfig(StatefulSet.UserName))

	if err != nil {
		log.Fatal(err)
	}

	var data []byte = []byte(StatefulSet.YamlData) // string to byte

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(global.GetK8sConfig(StatefulSet.UserName))

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
//@function: GetStatefulSet
//@description: 根据name获取StatefulSet记录
//@param: name uint
//@return: err error, StatefulSet model.StatefulSet

func ReadYamlStatefulSet(name string, namespace string, UserName string) (err error, StatefulSet model.StatefulSet) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientset.AppsV1().StatefulSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	//for _, deploy := range deployment.Items {
	y, err := ghodssyaml.Marshal(deployment)
	if err != nil {
		panic(err.Error)
	}
	//}

	yamldata := fmt.Sprintf("apiVersion: apps/v1\nkind: StatefulSet\n%s", string(y))

	StatefulSet.Name = name
	StatefulSet.Namespace = namespace
	StatefulSet.YamlData = yamldata

	return err, StatefulSet
}
