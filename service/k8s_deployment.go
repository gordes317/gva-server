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
	"time"

	"k8s.io/apimachinery/pkg/types"

	ghodssyaml "github.com/ghodss/yaml"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDeployment
//@description: 创建Deployment记录
//@param: Deployment model.Deployment
//@return: err error

func CreateDeployment(name string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 DeploymentsGetter 接口方法 Deployments 返回 DeploymentInterface
	// DeploymentInterface 接口拥有操作 Deployment 资源的方法，例如 Create、Update、Get、List 等方法
	deploymentsClient := clientset.AppsV1().Deployments("")

	//参数传递变量替换
	deploymentSpec := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: func() *int32 { i := int32(2); return &i }(),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	_, err = deploymentsClient.Create(context.TODO(), deploymentSpec, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDeployment
//@description: 删除Deployment记录
//@param: Deployment model.Deployment
//@return: err error

func DeleteDeployment(name string, namespace string, UserName string) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 DeploymentsGetter 接口方法 Deployments 返回 DeploymentInterface
	// DeploymentInterface 接口拥有操作 Deployment 资源的方法，例如 Create、Update、Get、List 等方法
	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDeploymentByIds
//@description: 批量删除Deployment记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDeploymentByIds(names request.NamesReqUser) (err error) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(names.UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 DeploymentsGetter 接口方法 Deployments 返回 DeploymentInterface
	// DeploymentInterface 接口拥有操作 Deployment 资源的方法，例如 Create、Update、Get、List 等方法
	//deploymentsClient := clientset.AppsV1().Deployments("")

	deletePolicy := metav1.DeletePropagationForeground

	for _, value := range names.Names {
		name := strings.Split(value, ";")[0]
		namespace := strings.Split(value, ";")[1]
		deploymentsClient := clientset.AppsV1().Deployments(namespace)
		if err := deploymentsClient.Delete(context.TODO(), name, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	return err
}

//@function: UpdateDeployment
//@description: 更新Deployment记录
//@param: Deployment *model.Deployment
//@return: err error

func UpdateDeployment(Deployment model.DeploymentUser) (err error) {
	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(Deployment.UserName))
	if err != nil {
		panic(err.Error())
	}

	s, err := clientset.AppsV1().
		Deployments(Deployment.Namespace).
		GetScale(context.TODO(), Deployment.Name, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	sc := *s
	sc.Spec.Replicas = Deployment.Replicas

	us, err := clientset.AppsV1().
		Deployments(Deployment.Namespace).
		UpdateScale(context.TODO(),
			Deployment.Name, &sc, metav1.UpdateOptions{})
	if err != nil {
		log.Fatal(err)
	}

	rt := *&us.Spec

	log := fmt.Sprintf("scale deployment to :%v success", rt)

	global.GVA_LOG.Info(log)

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDeployment
//@description: 根据name获取Deployment记录
//@param: name uint
//@return: err error, Deployment model.Deployment

func GetDeployment(name string, namespace string, UserName string) (err error, Deployment model.Deployment) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	// 通过实现 clientset 的 AppsV1Interface 接口列表中的 DeploymentsGetter 接口方法 Deployments 返回 DeploymentInterface
	// DeploymentInterface 接口拥有操作 Deployment 资源的方法，例如 Create、Update、Get、List 等方法
	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	result, err := deploymentsClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	Deployment.Name = result.Name
	Deployment.Namespace = result.Namespace
	Deployment.Replicas = result.Status.Replicas

	return err, Deployment
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDeploymentInfoList
//@description: 分页获取Deployment记录
//@param: info request.DeploymentSearch
//@return: err error, list interface{}, total int64

func GetDeploymentInfoList(info request.DeploymentSearchUser) (err error, list interface{}, total int64) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(info.UserName))
	if err != nil {
		panic(err.Error())
	}

	//DeploymentMetrics

	deploymentsClient := clientset.AppsV1().Deployments("")

	result, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	listResult := make([]interface{}, 0)

	var n model.Deployment
	for _, ser := range result.Items {
		n.Name = ser.Name
		n.Namespace = ser.Namespace
		n.Replicas = ser.Status.Replicas
		n.ReadyReplicas = ser.Status.ReadyReplicas
		n.UpdateReplicas = ser.Status.UpdatedReplicas
		n.AvailableReplicas = ser.Status.AvailableReplicas
		//n.Label = utils.MapToJson(ser.Spec.Selector.MatchLabels) //Selector

		n.Label = utils.MapToJson(ser.Spec.Template.Labels)
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
//@function: ApplyYamlDeployment
//@description: 更新Deployment Yaml记录
//@param: Deployment *model.Deployment
//@return: err error

func ApplyYamlDeployment(Deployment model.DeploymentUser) (err error) {

	//动态client
	dynamicClient, err := dynamic.NewForConfig(global.GetK8sConfig(Deployment.UserName))

	if err != nil {
		log.Fatal(err)
	}

	var data []byte = []byte(Deployment.YamlData) // string to byte

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(global.GetK8sConfig(Deployment.UserName))

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
//@function: GetDeployment
//@description: 根据name获取Deployment记录
//@param: name uint
//@return: err error, Deployment model.Deployment

func ReadYamlDeployment(name string, namespace string, UserName string) (err error, Deployment model.Deployment) {

	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(UserName))
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	//for _, deploy := range deployment.Items {
	y, err := ghodssyaml.Marshal(deployment)
	if err != nil {
		panic(err.Error)
	}
	//}

	yamldata := fmt.Sprintf("apiVersion: apps/v1\nkind: Deployment\n%s", string(y))

	Deployment.Name = name
	Deployment.Namespace = namespace
	Deployment.YamlData = yamldata

	return err, Deployment
}

//@function: RestartDeployment
//@description: 重启Deployment
//@param: Deployment *model.Deployment
//@return: err error
func RestartDeployment(Deployment model.DeploymentUser) (err error) {
	// create the clientset
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(Deployment.UserName))
	if err != nil {
		panic(err.Error())
	}

	deploymentsClient := clientset.AppsV1().Deployments(Deployment.Namespace)
	data := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"kubectl.kubernetes.io/restartedAt":"%s"}}}}}`, time.Now().String())
	_, err = deploymentsClient.Patch(context.Background(), Deployment.Name, types.StrategicMergePatchType, []byte(data), metav1.PatchOptions{FieldManager: "kubectl-rollout"})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
