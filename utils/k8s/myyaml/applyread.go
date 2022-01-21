package main

// // 	"bytes"
// "context"
// "flag"
// "fmt"

// // 	"gin-vue-admin/utils"
// // 	"io"
// // 	"io/ioutil"
// // 	"log"
// "path/filepath"

// ghodssyaml "github.com/ghodss/yaml"

// // 	"k8s.io/apimachinery/pkg/api/meta"
// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// // 	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
// // 	"k8s.io/apimachinery/pkg/runtime"
// // 	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
// //yamlutil "k8s.io/apimachinery/pkg/util/yaml"
// // 	"k8s.io/client-go/discovery"
// // "k8s.io/client-go/dynamic"
// "k8s.io/client-go/kubernetes"
// "k8s.io/client-go/rest"

// // 	"k8s.io/client-go/restmapper"
// "k8s.io/client-go/tools/clientcmd"
// "k8s.io/client-go/util/homedir"

// var config *rest.Config

// func init() {

// 	var kubeconfig *string
// 	var clientErr error
// 	if home := homedir.HomeDir(); home != "" {
// 		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
// 	} else {
// 		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
// 	}

// 	flag.Parse()

// 	// use the current context in kubeconfig
// 	config, clientErr = clientcmd.BuildConfigFromFlags("", *kubeconfig)
// 	if clientErr != nil {
// 		panic(clientErr.Error())
// 	} else {
// 		fmt.Println("connect k8s success")
// 	}

// }

// func create() {

// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	//使用上面的配置获取连接
// 	dd, err := dynamic.NewForConfig(config)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var data []byte

// 	if data, err = ioutil.ReadFile("/Users/wuchengjiang/Service/golang/go/packages/src/gin-vue-admin/server/utils/k8s/yaml/k8s-test.yaml"); err != nil {
// 		fmt.Print(err)
// 	}

// 	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewBuffer(data), 100)

// 	//循环每个yaml体
// 	for {
// 		var rawObj runtime.RawExtension
// 		if err = decoder.Decode(&rawObj); err != nil {
// 			break
// 		}

// 		obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)

// 		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}

// 		//获取支持的资源类型列表
// 		gr, err := restmapper.GetAPIGroupResources(clientset.Discovery())
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		//获取查询的资源类型
// 		mapper := restmapper.NewDiscoveryRESTMapper(gr)

// 		//查找gvk的REST映射
// 		mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		var dri dynamic.ResourceInterface

// 		//需要为 namespace 范围内的资源提供不同的接口
// 		if mapping.Scope.Name() == meta.RESTScopeNameNamespace {

// 			if unstructuredObj.GetNamespace() == "" {
// 				unstructuredObj.SetNamespace("default")
// 			}

// 			dri = dd.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())

// 		} else {

// 			dri = dd.Resource(mapping.Resource)
// 		}

// 		//创建对象
// 		obj2, err := dri.Create(context.Background(), unstructuredObj, metav1.CreateOptions{})

// 		//删除对象
// 		//err = dri.Delete(context.Background(), unstructuredObj.GetName(), metav1.DeleteOptions{})

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Printf("%s/%s created", obj2.GetKind(), obj2.GetName())

// 	}

// 	if err != io.EOF {
// 		log.Fatal("eof ", err)
// 	}

// 	return

// }

// func apply() {

// 	//使用上面的配置获取连接
// 	dynamicClient, err := dynamic.NewForConfig(config)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var data []byte

// 	if data, err = ioutil.ReadFile("/Users/wuchengjiang/Service/golang/go/packages/src/gin-vue-admin/server/utils/k8s/yaml/k8s-test.yaml"); err != nil {
// 		fmt.Print(err)
// 	}

// 	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)

// 	if err != nil {
// 		log.Fatal(discoveryClient)
// 	}
// 	applyOptions := utils.NewApplyOptions(dynamicClient, discoveryClient)

// 	if err := applyOptions.Apply(context.TODO(), data); err != nil {
// 		log.Fatalf("apply error: %v", err)
// 	}

// }

// func read() {

// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	nameSpace := "k8s-test"
// 	deployment, err := clientset.AppsV1().Deployments(nameSpace).List(context.TODO(), metav1.ListOptions{})

// 	for _, deploy := range deployment.Items {
// 		y, err := ghodssyaml.Marshal(deploy)
// 		if err != nil {
// 			panic(err.Error)
// 		}
// 		fmt.Println("deployment print in yaml :\n", string(y))
// 	}
// 	//apiVersion: v1
// 	//kind: Service
// 	fmt.Println("---------------")
// 	//打印service
// 	service, err := clientset.CoreV1().Services(nameSpace).List(context.TODO(), metav1.ListOptions{})

// 	for _, srv := range service.Items {
// 		y, err := ghodssyaml.Marshal(srv)
// 		if err != nil {
// 			panic(err.Error)
// 		}
// 		fmt.Println("service print in yaml :\n", string(y))
// 	}
// }

// func delete() {

// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	//使用上面的配置获取连接
// 	dd, err := dynamic.NewForConfig(config)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var data []byte

// 	if data, err = ioutil.ReadFile("/Users/wuchengjiang/Service/golang/go/packages/src/gin-vue-admin/server/utils/k8s/yaml/k8s-test.yaml"); err != nil {
// 		fmt.Print(err)
// 	}

// 	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewBuffer(data), 100)

// 	for {
// 		var rawObj runtime.RawExtension
// 		if err = decoder.Decode(&rawObj); err != nil {
// 			break
// 		}

// 		obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)

// 		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}

// 		//获取支持的资源类型列表
// 		gr, err := restmapper.GetAPIGroupResources(clientset.Discovery())
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		//获取查询的资源类型
// 		mapper := restmapper.NewDiscoveryRESTMapper(gr)

// 		//查找gvk的REST映射
// 		mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		var dri dynamic.ResourceInterface

// 		//需要为 namespace 范围内的资源提供不同的接口
// 		if mapping.Scope.Name() == meta.RESTScopeNameNamespace {

// 			if unstructuredObj.GetNamespace() == "" {
// 				unstructuredObj.SetNamespace("default")
// 			}

// 			dri = dd.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())

// 		} else {

// 			dri = dd.Resource(mapping.Resource)
// 		}

// 		//删除对象
// 		err = dri.Delete(context.Background(), unstructuredObj.GetName(), metav1.DeleteOptions{})

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 	}

// 	if err != io.EOF {
// 		log.Fatal("eof ", err)
// 	}

// 	return

// }
func main() {
	//	create()
	//	apply()
	//read()

	//	delete()
}
