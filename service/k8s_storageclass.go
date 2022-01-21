package service

import (
	"context"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"strings"

	stgv1 "k8s.io/api/storage/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateStorageClass
//@description: 创建StorageClass记录
//@param: StorageClass model.StorageClass
//@return: err error

func CreateStorageClass(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 StorageClasssGetter 接口方法 StorageClasss 返回 StorageClassInterface
	// StorageClassInterface 接口拥有操作 StorageClass 资源的方法，例如 Create、Update、Get、List 等方法
	storageclassClient := clientset.StorageV1().StorageClasses()

	namespace := &stgv1.StorageClass{}

	_, err = storageclassClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStorageClass
//@description: 删除StorageClass记录
//@param: StorageClass model.StorageClass
//@return: err error

func DeleteStorageClass(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 StorageClasssGetter 接口方法 StorageClasss 返回 StorageClassInterface
	// StorageClassInterface 接口拥有操作 StorageClass 资源的方法，例如 Create、Update、Get、List 等方法
	storageclassClient := clientset.StorageV1().StorageClasses()

	deletePolicy := metav1.DeletePropagationForeground
	if err := storageclassClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStorageClassByIds
//@description: 批量删除StorageClass记录
//@param: ids request.IdsReq
//@return: err error

func DeleteStorageClassByIds(names request.NamesReqUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(names.UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 StorageClasssGetter 接口方法 StorageClasss 返回 StorageClassInterface
	// StorageClassInterface 接口拥有操作 StorageClass 资源的方法，例如 Create、Update、Get、List 等方法
	storageclassClient := clientset.StorageV1().StorageClasses()

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		if err := storageclassClient.Delete(context.TODO(), value, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateStorageClass
//@description: 更新StorageClass记录
//@param: StorageClass *model.StorageClass
//@return: err error

func UpdateStorageClass(StorageClass model.StorageClassUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(StorageClass.UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &stgv1.StorageClass{}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 StorageClasssGetter 接口方法 StorageClasss 返回 StorageClassInterface
	// StorageClassInterface 接口拥有操作 StorageClass 资源的方法，例如 Create、Update、Get、List 等方法
	storageclassClient := clientset.StorageV1().StorageClasses()
	_, err = storageclassClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStorageClass
//@description: 根据name获取StorageClass记录
//@param: name uint
//@return: err error, StorageClass model.StorageClass

func GetStorageClass(name string, UserName string) (err error, StorageClass model.StorageClass) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 StorageClasssGetter 接口方法 StorageClasss 返回 StorageClassInterface
	// StorageClassInterface 接口拥有操作 StorageClass 资源的方法，例如 Create、Update、Get、List 等方法
	storageclassClient := clientset.StorageV1().StorageClasses()

	result, err := storageclassClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	StorageClass.Name = result.Name
	StorageClass.Provisioner = result.Provisioner
	StorageClass.ReclaimPolicy = string(*result.ReclaimPolicy)
	StorageClass.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, StorageClass
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStorageClassInfoList
//@description: 分页获取StorageClass记录
//@param: info request.StorageClassSearch
//@return: err error, list interface{}, total int64

func GetStorageClassInfoList(info request.StorageClassSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	storageclassClient := clientset.StorageV1().StorageClasses()

	result, err := storageclassClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n model.StorageClass
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Provisioner = ser.Provisioner
		n.ReclaimPolicy = string(*ser.ReclaimPolicy)
		n.VolumeBindingMode = string(*ser.VolumeBindingMode)
		//n.AllowVolumeExpansion = *ser.AllowVolumeExpansion

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
