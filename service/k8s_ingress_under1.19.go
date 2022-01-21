package service

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	//apiv1 "k8s.io/api/networking/v1"
	"strings"

	ghodssyaml "github.com/ghodss/yaml"
	//apiv1 "k8s.io/api/networking/v1"
	apiv1beta1 "k8s.io/api/networking/v1beta1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateIngress
//@description: 创建Ingress记录
//@param: Ingress model.Ingress
//@return: err error

func CreateIngressUnder119(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 NetworkingV1Interface 接口列表中的 IngresssGetter 接口方法 Ingresss 返回 IngressInterface
	// IngressInterface 接口拥有操作 Ingress 资源的方法，例如 Create、Update、Get、List 等方法
	//secretClient := clientset.NetworkingV1().Ingresses("")
	secretClient := clientset.NetworkingV1beta1().Ingresses("")

	namespace := &apiv1beta1.Ingress{}

	_, err = secretClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteIngress
//@description: 删除Ingress记录
//@param: Ingress model.Ingress
//@return: err error

func DeleteIngressUnder119(name, namespace, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 NetworkingV1Interface 接口列表中的 IngresssGetter 接口方法 Ingresss 返回 IngressInterface
	// IngressInterface 接口拥有操作 Ingress 资源的方法，例如 Create、Update、Get、List 等方法
	//secretClient := clientset.NetworkingV1().Ingresses(namespace)
	secretClient := clientset.NetworkingV1beta1().Ingresses(namespace)

	deletePolicy := metav1.DeletePropagationForeground
	if err := secretClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteIngressByIds
//@description: 批量删除Ingress记录
//@param: ids request.IdsReq
//@return: err error

func DeleteIngressByIdsUnder119(names request.NamesReqUser, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 NetworkingV1Interface 接口列表中的 IngresssGetter 接口方法 Ingresss 返回 IngressInterface
	// IngressInterface 接口拥有操作 Ingress 资源的方法，例如 Create、Update、Get、List 等方法
	//secretClient := clientset.NetworkingV1().Ingresses("")

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		name := strings.Split(value, ";")[0]
		namespace := strings.Split(value, ";")[1]
		//secretClient := clientset.NetworkingV1().Ingresses(namespace)
		secretClient := clientset.NetworkingV1beta1().Ingresses(namespace)
		if err := secretClient.Delete(context.TODO(), name, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateIngress
//@description: 更新Ingress记录
//@param: Ingress *model.Ingress
//@return: err error

func UpdateIngressUnder119(Ingress model.IngressUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(Ingress.UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &apiv1beta1.Ingress{}

	// 通过实现 clientset 的 NetworkingV1Interface 接口列表中的 IngresssGetter 接口方法 Ingresss 返回 IngressInterface
	// IngressInterface 接口拥有操作 Ingress 资源的方法，例如 Create、Update、Get、List 等方法
	//secretClient := clientset.NetworkingV1().Ingresses("")
	secretClient := clientset.NetworkingV1beta1().Ingresses("")
	_, err = secretClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetIngress
//@description: 根据name获取Ingress记录
//@param: name uint
//@return: err error, Ingress model.Ingress

func GetIngressUnder119(name string, namespace string, UserName string) (err error, Ingress model.Ingress) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 NetworkingV1Interface 接口列表中的 IngresssGetter 接口方法 Ingresss 返回 IngressInterface
	// IngressInterface 接口拥有操作 Ingress 资源的方法，例如 Create、Update、Get、List 等方法
	//secretClient := clientset.NetworkingV1().Ingresses(namespace)
	secretClient := clientset.NetworkingV1beta1().Ingresses(namespace)

	result, err := secretClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	Ingress.Name = result.Name
	Ingress.Namespace = result.Namespace

	//Ingress.Data = result.Data
	Ingress.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, Ingress
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetIngressInfoList
//@description: 分页获取Ingress记录
//@param: info request.IngressSearch
//@return: err error, list interface{}, total int64

func GetIngressInfoListUnder119(info request.IngressSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	//secretClient := clientset
	secretClient := clientset.NetworkingV1beta1().Ingresses("")

	result, err := secretClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n model.Ingress
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Namespace = ser.Namespace

		n.Host = ser.Spec.Rules[0].Host

		//slice to string
		p, err := json.Marshal(ser.Spec.Rules[0].HTTP.Paths)
		if err != nil {
			panic(err)
		}
		n.Paths = string(p)

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
//@function: GetIngress
//@description: 根据name获取Ingress记录
//@param: name uint
//@return: err error, Ingress model.Ingress

func ReadYamlIngressUnder119(name string, namespace string, UserName string) (err error, Ingress model.Ingress) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	//deployment, err := clientset.NetworkingV1().Ingresses(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	deployment, err := clientset.NetworkingV1beta1().Ingresses(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	//for _, deploy := range deployment.Items {
	y, err := ghodssyaml.Marshal(deployment)
	if err != nil {
		panic(err.Error)
	}
	//}

	yamldata := fmt.Sprintf("apiVersion: networking/v1\nkind: Ingresses\n%s", string(y))

	Ingress.Name = name
	Ingress.Namespace = namespace
	Ingress.YamlData = yamldata

	return err, Ingress
}
