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
	"k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateService
//@description: 创建Service记录
//@param: Service model.Service
//@return: err error

func CreateService(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 ServicesGetter 接口方法 Services 返回 ServiceInterface
	// ServiceInterface 接口拥有操作 Service 资源的方法，例如 Create、Update、Get、List 等方法
	servicesClient := clientset.CoreV1().Services("")

	//service 变量传递
	serviceSpec := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-service",
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"app": "demo",
			},
			Type: "LoadBalancer",
			Ports: []apiv1.ServicePort{
				{
					Port: 80,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 80,
					},
				},
			},
		},
	}

	_, err = servicesClient.Create(context.TODO(), serviceSpec, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteService
//@description: 删除Service记录
//@param: Service model.Service
//@return: err error

func DeleteService(name string, Namespace string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 ServicesGetter 接口方法 Services 返回 ServiceInterface
	// ServiceInterface 接口拥有操作 Service 资源的方法，例如 Create、Update、Get、List 等方法
	servicesClient := clientset.CoreV1().Services(Namespace)

	deletePolicy := metav1.DeletePropagationForeground
	if err := servicesClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteServiceByIds
//@description: 批量删除Service记录
//@param: ids request.IdsReq
//@return: err error

func DeleteServiceByIds(names request.NamesReqUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(names.UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 ServicesGetter 接口方法 Services 返回 ServiceInterface
	// ServiceInterface 接口拥有操作 Service 资源的方法，例如 Create、Update、Get、List 等方法
	//servicesClient := clientset.CoreV1().Services("")

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		name := strings.Split(value, ";")[0]
		namespace := strings.Split(value, ";")[1]
		servicesClient := clientset.CoreV1().Services(namespace)
		if err := servicesClient.Delete(context.TODO(), name, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateService
//@description: 更新Service记录
//@param: Service *model.Service
//@return: err error

func UpdateService(Service model.ServiceUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(Service.UserName))
	if err != nil {
		panic(err.Error())
	}

	//service 变量传递
	serviceSpec := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-service",
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"app": "demo",
			},
			Type: "LoadBalancer",
			Ports: []apiv1.ServicePort{
				{
					Port: 80,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 80,
					},
				},
			},
		},
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 ServicesGetter 接口方法 Services 返回 ServiceInterface
	// ServiceInterface 接口拥有操作 Service 资源的方法，例如 Create、Update、Get、List 等方法
	servicesClient := clientset.CoreV1().Services("")
	_, err = servicesClient.Update(context.TODO(), serviceSpec, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetService
//@description: 根据name获取Service记录
//@param: name uint
//@return: err error, Service model.Service

func GetService(name string, UserName string) (err error, Service model.Service) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 ServicesGetter 接口方法 Services 返回 ServiceInterface
	// ServiceInterface 接口拥有操作 Service 资源的方法，例如 Create、Update、Get、List 等方法
	servicesClient := clientset.CoreV1().Services("")

	result, err := servicesClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	Service.Name = result.Name
	Service.CreateTime = result.ObjectMeta.CreationTimestamp.Time

	return err, Service
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetServiceInfoList
//@description: 分页获取Service记录
//@param: info request.ServiceSearch
//@return: err error, list interface{}, total int64

func GetServiceInfoList(info request.ServiceSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	//ServiceMetrics

	servicesClient := clientset.CoreV1().Services("")

	result, err := servicesClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n model.Service
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Namespace = ser.Namespace
		n.Type = string(ser.Spec.Type)

		var port []string
		for _, value := range ser.Spec.Ports {

			portStr := fmt.Sprintf("%d:%d/%v", value.Port, value.NodePort, value.Protocol)
			port = append(port, portStr)

		}

		n.Ports = strings.Join(port, ",")

		n.ClusterIP = ser.Spec.ClusterIP

		n.ExternalIPs = strings.Join(ser.Spec.ExternalIPs, ",")
		n.Selector = utils.MapToJson(ser.Spec.Selector)

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
//@function: ApplyYamlService
//@description: 更新Service Yaml记录
//@param: Service *model.Service
//@return: err error

func ApplyYamlService(Service model.ServiceUser) (err error) {

	//动态client
	dynamicClient, err := dynamic.NewForConfig(global.GetK8sConfig(Service.UserName))

	if err != nil {
		log.Fatal(err)
	}

	var data []byte = []byte(Service.YamlData) // string to byte

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(global.GetK8sConfig(Service.UserName))

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
//@function: GetService
//@description: 根据name获取Service记录
//@param: name uint
//@return: err error, Service model.Service

func ReadYamlService(name string, namespace string, UserName string) (err error, Service model.Service) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientset.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	//for _, deploy := range deployment.Items {
	y, err := ghodssyaml.Marshal(deployment)
	if err != nil {
		panic(err.Error)
	}
	//}

	yamldata := fmt.Sprintf("apiVersion: v1\nkind: Service\n%s", string(y))

	Service.Name = name
	Service.Namespace = namespace
	Service.YamlData = yamldata

	return err, Service
}
