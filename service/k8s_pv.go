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
//@function: CreatePV
//@description: 创建PV记录
//@param: PV model.PV
//@return: err error

func CreatePV(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PVsGetter 接口方法 PVs 返回 PVInterface
	// PVInterface 接口拥有操作 PV 资源的方法，例如 Create、Update、Get、List 等方法
	pvClient := clientset.CoreV1().PersistentVolumes()

	namespace := &apiv1.PersistentVolume{}

	_, err = pvClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePV
//@description: 删除PV记录
//@param: PV model.PV
//@return: err error

func DeletePV(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PVsGetter 接口方法 PVs 返回 PVInterface
	// PVInterface 接口拥有操作 PV 资源的方法，例如 Create、Update、Get、List 等方法
	pvClient := clientset.CoreV1().PersistentVolumes()

	deletePolicy := metav1.DeletePropagationForeground
	if err := pvClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePVByIds
//@description: 批量删除PV记录
//@param: ids request.IdsReq
//@return: err error

func DeletePVByIds(names request.NamesReqUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(names.UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PVsGetter 接口方法 PVs 返回 PVInterface
	// PVInterface 接口拥有操作 PV 资源的方法，例如 Create、Update、Get、List 等方法
	pvClient := clientset.CoreV1().PersistentVolumes()

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		if err := pvClient.Delete(context.TODO(), value, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePV
//@description: 更新PV记录
//@param: PV *model.PV
//@return: err error

func UpdatePV(PV model.PVUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(PV.UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &apiv1.PersistentVolume{}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PVsGetter 接口方法 PVs 返回 PVInterface
	// PVInterface 接口拥有操作 PV 资源的方法，例如 Create、Update、Get、List 等方法
	pvClient := clientset.CoreV1().PersistentVolumes()
	_, err = pvClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPV
//@description: 根据name获取PV记录
//@param: name uint
//@return: err error, PV model.PV

func GetPV(name string, UserName string) (err error, PV model.PV) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PVsGetter 接口方法 PVs 返回 PVInterface
	// PVInterface 接口拥有操作 PV 资源的方法，例如 Create、Update、Get、List 等方法
	pvClient := clientset.CoreV1().PersistentVolumes()

	result, err := pvClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	PV.Name = result.Name
	PV.Status = string(result.Status.Phase)
	PV.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, PV
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPVInfoList
//@description: 分页获取PV记录
//@param: info request.PVSearch
//@return: err error, list interface{}, total int64

func GetPVInfoList(info request.PVSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	pvClient := clientset.CoreV1().PersistentVolumes()

	result, err := pvClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n model.PV
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Status = string(ser.Status.Phase)
		n.Capacity = string(ser.Spec.Capacity.Storage().String())
		n.AccessMode = string(ser.Spec.AccessModes[0])
		n.ReclaimPolicy = string(ser.Spec.PersistentVolumeReclaimPolicy)
		n.Claim = ser.Spec.ClaimRef.Name
		n.StorageClass = string(ser.Spec.StorageClassName)
		n.Reason = ser.Status.Reason
		n.VolumeMode = string(*ser.Spec.VolumeMode)
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
