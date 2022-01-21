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
	apiv1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateConfigMap
//@description: 创建ConfigMap记录
//@param: ConfigMap model.ConfigMap
//@return: err error

func CreateConfigMap(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 ConfigMapsGetter 接口方法 ConfigMaps 返回 ConfigMapInterface
	// ConfigMapInterface 接口拥有操作 ConfigMap 资源的方法，例如 Create、Update、Get、List 等方法
	configmapClient := clientset.CoreV1().ConfigMaps("")

	namespace := &apiv1.ConfigMap{}

	_, err = configmapClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteConfigMap
//@description: 删除ConfigMap记录
//@param: ConfigMap model.ConfigMap
//@return: err error

func DeleteConfigMap(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 ConfigMapsGetter 接口方法 ConfigMaps 返回 ConfigMapInterface
	// ConfigMapInterface 接口拥有操作 ConfigMap 资源的方法，例如 Create、Update、Get、List 等方法
	configmapClient := clientset.CoreV1().ConfigMaps("")

	deletePolicy := metav1.DeletePropagationForeground
	if err := configmapClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteConfigMapByIds
//@description: 批量删除ConfigMap记录
//@param: ids request.IdsReq
//@return: err error

func DeleteConfigMapByIds(names request.NamesReq, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 ConfigMapsGetter 接口方法 ConfigMaps 返回 ConfigMapInterface
	// ConfigMapInterface 接口拥有操作 ConfigMap 资源的方法，例如 Create、Update、Get、List 等方法
	configmapClient := clientset.CoreV1().ConfigMaps("")

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		if err := configmapClient.Delete(context.TODO(), value, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateConfigMap
//@description: 更新ConfigMap记录
//@param: ConfigMap *model.ConfigMap
//@return: err error

func UpdateConfigMap(ConfigMap model.ConfigMap, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &apiv1.ConfigMap{}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 ConfigMapsGetter 接口方法 ConfigMaps 返回 ConfigMapInterface
	// ConfigMapInterface 接口拥有操作 ConfigMap 资源的方法，例如 Create、Update、Get、List 等方法
	configmapClient := clientset.CoreV1().ConfigMaps("")
	_, err = configmapClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetConfigMap
//@description: 根据name获取ConfigMap记录
//@param: name uint
//@return: err error, ConfigMap model.ConfigMap

func GetConfigMap(name string, namespace string, UserName string) (err error, ConfigMap model.ConfigMap) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 ConfigMapsGetter 接口方法 ConfigMaps 返回 ConfigMapInterface
	// ConfigMapInterface 接口拥有操作 ConfigMap 资源的方法，例如 Create、Update、Get、List 等方法
	configmapClient := clientset.CoreV1().ConfigMaps(namespace)

	result, err := configmapClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	ConfigMap.Name = result.Name
	ConfigMap.Namespace = result.Namespace
	ConfigMap.Data = result.Data
	ConfigMap.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, ConfigMap
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetConfigMapInfoList
//@description: 分页获取ConfigMap记录
//@param: info request.ConfigMapSearch
//@return: err error, list interface{}, total int64

func GetConfigMapInfoList(info request.ConfigMapSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	configmapClient := clientset.CoreV1().ConfigMaps("")

	result, err := configmapClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n model.ConfigMap
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Namespace = ser.Namespace

		n.CreateTime = ser.ObjectMeta.CreationTimestamp.Time

		//按照条件搜索判断条件
		if strings.Contains(n.Name, info.Name) && strings.Contains(n.Namespace, info.Namespace) {
			listResult = append(listResult, n)
		}
	}

	total = int64(len(listResult))

	// 分页方法
	list, err = utils.Paginator(info.Page, info.PageSize, listResult)

	return err, list, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ApplyYamlConfigMap
//@description: 更新ConfigMap Yaml记录
//@param: ConfigMap *model.ConfigMap
//@return: err error

func ApplyYamlConfigMap(ConfigMap model.ConfigMap, UserName string) (err error) {

	//动态client
	dynamicClient, err := dynamic.NewForConfig(global.GetK8sConfig(UserName))

	if err != nil {
		log.Fatal(err)
	}

	var data []byte = []byte(ConfigMap.YamlData) // string to byte

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
//@function: GetConfigMap
//@description: 根据name获取ConfigMap记录
//@param: name uint
//@return: err error, ConfigMap model.ConfigMap

func ReadYamlConfigMap(name string, namespace string, UserName string) (err error, ConfigMap model.ConfigMap) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	//for _, deploy := range deployment.Items {
	y, err := ghodssyaml.Marshal(deployment)
	if err != nil {
		panic(err.Error)
	}
	//}

	yamldata := fmt.Sprintf("apiVersion: v1\nkind: ConfigMap\n%s", string(y))

	ConfigMap.Name = name
	ConfigMap.Namespace = namespace
	ConfigMap.YamlData = yamldata

	return err, ConfigMap
}
