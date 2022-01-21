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
	apiv1 "k8s.io/api/batch/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateCronJob
//@description: 创建CronJob记录
//@param: CronJob model.CronJob
//@return: err error

func CreateCronJob(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 BatchV1Interface 接口列表中的 CronJobsGetter 接口方法 CronJobs 返回 CronJobInterface
	// CronJobInterface 接口拥有操作 CronJob 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.BatchV1().CronJobs("")

	namespace := &apiv1.CronJob{}

	_, err = secretClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err.Error())
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteCronJob
//@description: 删除CronJob记录
//@param: CronJob model.CronJob
//@return: err error

func DeleteCronJob(name string, namespace string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 BatchV1Interface 接口列表中的 CronJobsGetter 接口方法 CronJobs 返回 CronJobInterface
	// CronJobInterface 接口拥有操作 CronJob 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.BatchV1().CronJobs(namespace)

	deletePolicy := metav1.DeletePropagationForeground
	if err := secretClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err.Error())
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteCronJobByIds
//@description: 批量删除CronJob记录
//@param: ids request.IdsReq
//@return: err error

func DeleteCronJobByIds(names request.NamesReqUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(names.UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 BatchV1Interface 接口列表中的 CronJobsGetter 接口方法 CronJobs 返回 CronJobInterface
	// CronJobInterface 接口拥有操作 CronJob 资源的方法，例如 Create、Update、Get、List 等方法
	//secretClient := clientset.BatchV1().CronJobs("")

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		name := strings.Split(value, ";")[0]
		namespace := strings.Split(value, ";")[1]
		secretClient := clientset.BatchV1().CronJobs(namespace)
		if err := secretClient.Delete(context.TODO(), name, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err.Error())
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCronJob
//@description: 更新CronJob记录
//@param: CronJob *model.CronJob
//@return: err error

func UpdateCronJob(CronJob model.CronJobUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(CronJob.UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &apiv1.CronJob{}

	// 通过实现 clientset 的 BatchV1Interface 接口列表中的 CronJobsGetter 接口方法 CronJobs 返回 CronJobInterface
	// CronJobInterface 接口拥有操作 CronJob 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.BatchV1().CronJobs(CronJob.Namespace)
	_, err = secretClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCronJob
//@description: 根据name获取CronJob记录
//@param: name uint
//@return: err error, CronJob model.CronJob

func GetCronJob(name string, namespace string, UserName string) (err error, CronJob model.CronJob) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 BatchV1Interface 接口列表中的 CronJobsGetter 接口方法 CronJobs 返回 CronJobInterface
	// CronJobInterface 接口拥有操作 CronJob 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.BatchV1beta1().CronJobs(namespace)

	result, err := secretClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	CronJob.Name = result.Name
	CronJob.Namespace = result.Namespace

	//CronJob.Data = result.Data
	CronJob.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, CronJob
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCronJobInfoList
//@description: 分页获取CronJob记录
//@param: info request.CronJobSearch
//@return: err error, list interface{}, total int64

func GetCronJobInfoList(info request.CronJobSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	secretClient := clientset.BatchV1beta1().CronJobs("") //BatchV1 如果 result 为空的话会报错

	result, err := secretClient.List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	listResult := make([]interface{}, 0)

	var n model.CronJob
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Namespace = ser.Namespace
		n.Schedule = ser.Spec.Schedule
		n.Suspend = *ser.Spec.Suspend

		//slice to string
		at, err := json.Marshal(ser.Status.Active)
		if err != nil {
			panic(err)
		}
		n.Active = string(at)
		n.LastScheduleTime = ser.Status.LastScheduleTime.Time

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
//@function: ApplyYamlCronJob
//@description: 更新CronJob Yaml记录
//@param: CronJob *model.CronJob
//@return: err error

func ApplyYamlCronJob(CronJob model.CronJobUser) (err error) {

	//动态client
	dynamicClient, err := dynamic.NewForConfig(global.GetK8sConfig(CronJob.UserName))

	if err != nil {
		log.Fatal(err)
	}

	var data []byte = []byte(CronJob.YamlData) // string to byte

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(global.GetK8sConfig(CronJob.UserName))

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
//@function: GetCronJob
//@description: 根据name获取CronJob记录
//@param: name uint
//@return: err error, CronJob model.CronJob

func ReadYamlCronJob(name string, namespace string, UserName string) (err error, CronJob model.CronJob) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientset.BatchV1().CronJobs(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	//for _, deploy := range deployment.Items {
	y, err := ghodssyaml.Marshal(deployment)
	if err != nil {
		panic(err.Error)
	}
	//}

	yamldata := fmt.Sprintf("apiVersion: batch/v1\nkind: CronJob\n%s", string(y))

	CronJob.Name = name
	CronJob.Namespace = namespace
	CronJob.YamlData = yamldata

	return err, CronJob
}
