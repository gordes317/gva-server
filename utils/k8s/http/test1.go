package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var config *rest.Config

func init() {

	var kubeconfig *string
	var clientErr error
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	//	k8sconfig := flag.String("k8sconfig", "/Users/wuchengjiang/Service/golang/go/packages/src/gin-vue-admin/server/kubeconfig/.k8sconfig", "kubernetes config file path")

	flag.Parse()

	// use the current context in kubeconfig
	config, clientErr = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if clientErr != nil {
		panic(clientErr.Error())
	} else {
		fmt.Println("connect k8s success")
	}

}

func main() {

	fmt.Println("Prepare config object.")

	config.APIPath = "/api"
	config.GroupVersion = &corev1.SchemeGroupVersion

	config.NegotiatedSerializer = scheme.Codecs

	fmt.Println("Init RESTClient.")

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		log.Fatalln(err)
	}

	// 获取pod列表。这里只会从namespace为"kube-system"中获取指定的资源(pods)
	result := &corev1.PodList{}

	if err := restClient.
		Get().
		Namespace("kube-system").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result); err != nil {
		panic(err)
	}

	fmt.Println("Print all listed pods.")

	// 打印所有获取到的pods资源，输出到标准输出
	for _, d := range result.Items {
		fmt.Printf("NAMESPACE: %v NAME: %v \t STATUS: %v \n", d.Namespace, d.Name, d.Status.Phase)
	}

	// cache.NewListWatchFromClient()
	// cache.NewInformer()
	// workqueue.NewRateLimitingQueue()

}
