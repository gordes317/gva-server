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

	"k8s.io/client-go/kubernetes"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateNamespace
//@description: 创建Namespace记录
//@param: Namespace model.Namespace
//@return: err error

func CreateNamespace(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NamespacesGetter 接口方法 Namespaces 返回 NamespaceInterface
	// NamespaceInterface 接口拥有操作 Namespace 资源的方法，例如 Create、Update、Get、List 等方法
	namespacesClient := clientset.CoreV1().Namespaces()

	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: apiv1.NamespaceStatus{
			Phase: apiv1.NamespaceActive,
		},
	}

	_, err = namespacesClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNamespace
//@description: 删除Namespace记录
//@param: Namespace model.Namespace
//@return: err error

func DeleteNamespace(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NamespacesGetter 接口方法 Namespaces 返回 NamespaceInterface
	// NamespaceInterface 接口拥有操作 Namespace 资源的方法，例如 Create、Update、Get、List 等方法
	namespacesClient := clientset.CoreV1().Namespaces()

	deletePolicy := metav1.DeletePropagationForeground
	if err := namespacesClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNamespaceByIds
//@description: 批量删除Namespace记录
//@param: ids request.IdsReq
//@return: err error

func DeleteNamespaceByIds(names request.NamesReq, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NamespacesGetter 接口方法 Namespaces 返回 NamespaceInterface
	// NamespaceInterface 接口拥有操作 Namespace 资源的方法，例如 Create、Update、Get、List 等方法
	namespacesClient := clientset.CoreV1().Namespaces()

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		if err := namespacesClient.Delete(context.TODO(), value, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateNamespace
//@description: 更新Namespace记录
//@param: Namespace *model.Namespace
//@return: err error

func UpdateNamespace(Namespace model.NamespaceUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(Namespace.UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: Namespace.Name,
		},
		Status: apiv1.NamespaceStatus{
			Phase: apiv1.NamespacePhase(Namespace.Status),
		},
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NamespacesGetter 接口方法 Namespaces 返回 NamespaceInterface
	// NamespaceInterface 接口拥有操作 Namespace 资源的方法，例如 Create、Update、Get、List 等方法
	namespacesClient := clientset.CoreV1().Namespaces()
	_, err = namespacesClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNamespace
//@description: 根据name获取Namespace记录
//@param: name uint
//@return: err error, Namespace model.Namespace

func GetNamespace(name string, UserName string) (err error, Namespace model.Namespace) {

	// create the clientset
	//clientset, err := kubernetes.NewForConfig(global.GVA_K8SConfig)
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NamespacesGetter 接口方法 Namespaces 返回 NamespaceInterface
	// NamespaceInterface 接口拥有操作 Namespace 资源的方法，例如 Create、Update、Get、List 等方法
	namespacesClient := clientset.CoreV1().Namespaces()

	result, err := namespacesClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	Namespace.Name = result.Name
	Namespace.Status = string(result.Status.Phase)
	Namespace.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, Namespace
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNamespaceInfoList
//@description: 分页获取Namespace记录
//@param: info request.NamespaceSearch
//@return: err error, list interface{}, total int64

func GetNamespaceInfoList(info request.NamespaceList) (err error, list interface{}, total int64) {

	//测试多集群切换

	//t := make(map[string]*rest.Config)

	//fmt.Println("当前集群变量:", global.GVA_K8SCurrent)

	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	namespacesClient := clientset.CoreV1().Namespaces()

	result, err := namespacesClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n model.Namespace
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Status = string(ser.Status.Phase)
		n.CreateTime = ser.ObjectMeta.CreationTimestamp.Time

		//按照条件搜索判断条件
		if strings.Contains(n.Name, info.Name) && strings.Contains(n.Status, info.Status) {
			listResult = append(listResult, n)
		}

	}

	total = int64(len(listResult))

	// 分页方法
	list, err = utils.Paginator(info.Page, info.PageSize, listResult)

	return err, list, total
}
