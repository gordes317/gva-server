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
//@function: CreatePVC
//@description: 创建PVC记录
//@param: PVC model.PVC
//@return: err error

func CreatePVC(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PVCsGetter 接口方法 PVCs 返回 PVCInterface
	// PVCInterface 接口拥有操作 PVC 资源的方法，例如 Create、Update、Get、List 等方法
	pvcClient := clientset.CoreV1().PersistentVolumeClaims("")

	namespace := &apiv1.PersistentVolumeClaim{}

	_, err = pvcClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePVC
//@description: 删除PVC记录
//@param: PVC model.PVC
//@return: err error

func DeletePVC(name, namespace, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PVCsGetter 接口方法 PVCs 返回 PVCInterface
	// PVCInterface 接口拥有操作 PVC 资源的方法，例如 Create、Update、Get、List 等方法
	pvcClient := clientset.CoreV1().PersistentVolumeClaims(namespace)

	deletePolicy := metav1.DeletePropagationForeground
	if err := pvcClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePVCByIds
//@description: 批量删除PVC记录
//@param: ids request.IdsReq
//@return: err error

func DeletePVCByIds(names request.NamesReqUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(names.UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PVCsGetter 接口方法 PVCs 返回 PVCInterface
	// PVCInterface 接口拥有操作 PVC 资源的方法，例如 Create、Update、Get、List 等方法
	//pvcClient := clientset.CoreV1().PersistentVolumeClaims("")

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		name := strings.Split(value, ";")[0]
		namespace := strings.Split(value, ";")[1]
		pvcClient := clientset.CoreV1().PersistentVolumeClaims(namespace)
		if err := pvcClient.Delete(context.TODO(), name, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePVC
//@description: 更新PVC记录
//@param: PVC *model.PVC
//@return: err error

func UpdatePVC(PVC model.PVCUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(PVC.UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &apiv1.PersistentVolumeClaim{}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PVCsGetter 接口方法 PVCs 返回 PVCInterface
	// PVCInterface 接口拥有操作 PVC 资源的方法，例如 Create、Update、Get、List 等方法
	pvcClient := clientset.CoreV1().PersistentVolumeClaims(PVC.Namespace)
	_, err = pvcClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPVC
//@description: 根据name获取PVC记录
//@param: name uint
//@return: err error, PVC model.PVC

func GetPVC(name string, namespace string, UserName string) (err error, PVC model.PVC) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PVCsGetter 接口方法 PVCs 返回 PVCInterface
	// PVCInterface 接口拥有操作 PVC 资源的方法，例如 Create、Update、Get、List 等方法
	pvcClient := clientset.CoreV1().PersistentVolumeClaims(namespace)

	result, err := pvcClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	PVC.Name = result.Name
	PVC.Status = string(result.Status.Phase)
	PVC.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, PVC
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPVCInfoList
//@description: 分页获取PVC记录
//@param: info request.PVCSearch
//@return: err error, list interface{}, total int64

func GetPVCInfoList(info request.PVCSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	pvcClient := clientset.CoreV1().PersistentVolumeClaims("")

	result, err := pvcClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n model.PVC
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Status = string(ser.Status.Phase)
		n.Namespace = ser.Namespace
		n.Capacity = ser.Status.Capacity.Storage().String()

		n.Volume = string(ser.Spec.VolumeName)

		n.AccessMode = string(ser.Spec.AccessModes[0])
		//	n.StorageClass = string(*ser.Spec.StorageClassName)

		n.VolumeMode = string(*ser.Spec.VolumeMode)
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
