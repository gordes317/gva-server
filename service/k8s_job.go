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
	apiv1 "k8s.io/api/batch/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateJob
//@description: 创建Job记录
//@param: Job model.Job
//@return: err error

func CreateJob(name string, ns string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 BatchV1Interface 接口列表中的 JobsGetter 接口方法 Jobs 返回 JobInterface
	// JobInterface 接口拥有操作 Job 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.BatchV1().Jobs(ns)

	namespace := &apiv1.Job{}

	_, err = secretClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err.Error())
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteJob
//@description: 删除Job记录
//@param: Job model.Job
//@return: err error

func DeleteJob(name, namespace, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 BatchV1Interface 接口列表中的 JobsGetter 接口方法 Jobs 返回 JobInterface
	// JobInterface 接口拥有操作 Job 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.BatchV1().Jobs(namespace)

	deletePolicy := metav1.DeletePropagationForeground
	if err := secretClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err.Error())
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteJobByIds
//@description: 批量删除Job记录
//@param: ids request.IdsReq
//@return: err error

func DeleteJobByIds(names request.NamesReqUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(names.UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 BatchV1Interface 接口列表中的 JobsGetter 接口方法 Jobs 返回 JobInterface
	// JobInterface 接口拥有操作 Job 资源的方法，例如 Create、Update、Get、List 等方法
	//secretClient := clientset.BatchV1().Jobs("")

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		name := strings.Split(value, ";")[0]
		namespace := strings.Split(value, ";")[1]
		secretClient := clientset.BatchV1().Jobs(namespace)
		if err := secretClient.Delete(context.TODO(), name, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err.Error())
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateJob
//@description: 更新Job记录
//@param: Job *model.Job
//@return: err error

func UpdateJob(Job model.JobUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(Job.UserName))
	if err != nil {
		panic(err.Error())
	}

	namespace := &apiv1.Job{}

	// 通过实现 clientset 的 BatchV1Interface 接口列表中的 JobsGetter 接口方法 Jobs 返回 JobInterface
	// JobInterface 接口拥有操作 Job 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.BatchV1().Jobs(Job.Namespace)
	_, err = secretClient.Update(context.TODO(), namespace, metav1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetJob
//@description: 根据name获取Job记录
//@param: name uint
//@return: err error, Job model.Job

func GetJob(name string, namespace string, UserName string) (err error, Job model.Job) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 BatchV1Interface 接口列表中的 JobsGetter 接口方法 Jobs 返回 JobInterface
	// JobInterface 接口拥有操作 Job 资源的方法，例如 Create、Update、Get、List 等方法
	secretClient := clientset.BatchV1().Jobs(namespace)

	result, err := secretClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	Job.Name = result.Name
	Job.Namespace = result.Namespace

	//Job.Data = result.Data
	Job.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, Job
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetJobInfoList
//@description: 分页获取Job记录
//@param: info request.JobSearch
//@return: err error, list interface{}, total int64

func GetJobInfoList(info request.JobSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	secretClient := clientset.BatchV1().Jobs("")

	result, err := secretClient.List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	listResult := make([]interface{}, 0)

	var n model.Job
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Namespace = ser.Namespace
		n.Completions = *ser.Spec.Completions
		n.Successed = ser.Status.Succeeded
		n.Failed = ser.Status.Failed

		//fmt.Println("====", ser.Status.CompletionTime.Time)
		//n.CompletionTime = ser.Status.CompletionTime.Time
		n.StartTime = ser.Status.StartTime.Time
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
//@function: ApplyYamlJob
//@description: 更新Job Yaml记录
//@param: Job *model.Job
//@return: err error

func ApplyYamlJob(Job model.JobUser) (err error) {

	//动态client
	dynamicClient, err := dynamic.NewForConfig(global.GetK8sConfig(Job.UserName))

	if err != nil {
		log.Fatal(err.Error())
	}

	var data []byte = []byte(Job.YamlData) // string to byte

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(global.GetK8sConfig(Job.UserName))

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
//@function: GetJob
//@description: 根据name获取Job记录
//@param: name uint
//@return: err error, Job model.Job

func ReadYamlJob(name string, namespace string, UserName string) (err error, Job model.Job) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientset.BatchV1().Jobs(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	//for _, deploy := range deployment.Items {
	y, err := ghodssyaml.Marshal(deployment)
	if err != nil {
		panic(err.Error())
	}
	//}

	yamldata := fmt.Sprintf("apiVersion: batch/v1\nkind: Job\n%s", string(y))

	Job.Name = name
	Job.Namespace = namespace
	Job.YamlData = yamldata

	return err, Job
}
