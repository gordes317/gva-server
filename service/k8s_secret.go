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
//@function: CreateSecret
//@description: 创建Secret记录
//@param: Secret model.Secret
//@return: err error

func CreateSecret(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 SecretsGetter 接口方法 Secrets 返回 SecretInterface
	// SecretInterface 接口拥有操作 Secret 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.CoreV1().Secrets("")

	namespace := &apiv1.Secret{}

	_, err = secretClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSecret
//@description: 删除Secret记录
//@param: Secret model.Secret
//@return: err error

func DeleteSecret(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 SecretsGetter 接口方法 Secrets 返回 SecretInterface
	// SecretInterface 接口拥有操作 Secret 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.CoreV1().Secrets("")

	deletePolicy := metav1.DeletePropagationForeground
	if err := secretClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSecretByIds
//@description: 批量删除Secret记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSecretByIds(names request.NamesReqUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(names.UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 SecretsGetter 接口方法 Secrets 返回 SecretInterface
	// SecretInterface 接口拥有操作 Secret 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.CoreV1().Secrets("")

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		if err := secretClient.Delete(context.TODO(), value, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSecret
//@description: 更新Secret记录
//@param: Secret *model.Secret
//@return: err error

func UpdateSecret(Secret model.Secret, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &apiv1.Secret{}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 SecretsGetter 接口方法 Secrets 返回 SecretInterface
	// SecretInterface 接口拥有操作 Secret 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.CoreV1().Secrets("")
	_, err = secretClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSecret
//@description: 根据name获取Secret记录
//@param: name uint
//@return: err error, Secret model.Secret

func GetSecret(name string, namespace string, UserName string) (err error, Secret model.Secret) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 SecretsGetter 接口方法 Secrets 返回 SecretInterface
	// SecretInterface 接口拥有操作 Secret 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.CoreV1().Secrets(namespace)

	result, err := secretClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	Secret.Name = result.Name
	Secret.Namespace = result.Namespace

	Secret.Data = result.Data
	Secret.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, Secret
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSecretInfoList
//@description: 分页获取Secret记录
//@param: info request.SecretSearch
//@return: err error, list interface{}, total int64

func GetSecretInfoList(info request.SecretSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	secretClient := clientset.CoreV1().Secrets("")

	result, err := secretClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n model.Secret
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Namespace = ser.Namespace
		n.Type = string(ser.Type)
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
//@function: ApplyYamlSecret
//@description: 更新Secret Yaml记录
//@param: Secret *model.Secret
//@return: err error

func ApplyYamlSecret(Secret model.SecretUser) (err error) {

	//动态client
	dynamicClient, err := dynamic.NewForConfig(global.GetK8sConfig(Secret.UserName))

	if err != nil {
		log.Fatal(err)
	}

	var data []byte = []byte(Secret.YamlData) // string to byte

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(global.GetK8sConfig(Secret.UserName))

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
//@function: GetSecret
//@description: 根据name获取Secret记录
//@param: name uint
//@return: err error, Secret model.Secret

func ReadYamlSecret(name string, namespace string, UserName string) (err error, Secret model.Secret) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientset.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	//for _, deploy := range deployment.Items {
	y, err := ghodssyaml.Marshal(deployment)
	if err != nil {
		panic(err.Error)
	}
	//}

	yamldata := fmt.Sprintf("apiVersion: v1\nkind: Secret\n%s", string(y))

	Secret.Name = name
	Secret.Namespace = namespace
	Secret.YamlData = yamldata

	return err, Secret
}
