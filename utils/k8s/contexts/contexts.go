package main

import (
	"context"
	"errors"
	"fmt"
	"gin-vue-admin/utils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"path/filepath"

	"k8s.io/client-go/util/homedir"

	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	clientcmdapilatest "k8s.io/client-go/tools/clientcmd/api/latest"
)

type Client struct {
	ClientSet          *kubernetes.Clientset
	ClientConfig       clientcmd.ClientConfig
	ClientApiExtension *apiextension.Clientset
	//ClientV1Alpha1     *v1alpha1.V1Alpha1Client
}

func NewClient(kubeConfigPath string) (*Client, error) {
	var kubeconfig string
	if kubeConfigPath != "" {
		kubeconfig = kubeConfigPath
	} else {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		} else {
			return nil, errors.New("no kubeconfig available")
		}
	}
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{})

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("kubernetes config file: %s", err)
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	clientExtension, err := apiextension.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	c := &Client{
		ClientSet:          clientset,
		ClientConfig:       clientConfig,
		ClientApiExtension: clientExtension,
		//	ClientV1Alpha1:     clientV1Alpha1,
	}

	return c, nil
}

func (c *Client) GetConfigClusters() (map[string]*clientcmdapi.Cluster, error) {
	config, err := c.ClientConfig.ConfigAccess().GetStartingConfig()

	if err != nil {
		return nil, err
	}
	return config.Clusters, nil
}

func (c *Client) GetConfigContext() (map[string]*clientcmdapi.Context, string, error) {
	config, err := c.ClientConfig.ConfigAccess().GetStartingConfig()
	if err != nil {
		return nil, "", err
	}
	return config.Contexts, config.CurrentContext, nil
}

func (c *Client) GetCurrentContext() (string, error) {
	config, err := c.ClientConfig.ConfigAccess().GetStartingConfig()
	if err != nil {
		return "", err
	}
	return config.CurrentContext, nil
}

func (c *Client) SwitchContext(contextName string) error {
	config, err := c.ClientConfig.ConfigAccess().GetStartingConfig()
	if err != nil {
		return err
	}
	config.CurrentContext = contextName
	err = clientcmd.ModifyConfig(c.ClientConfig.ConfigAccess(), *config, true)
	return err
}

type MergeOptions struct {
	overwrite   bool
	backup      bool
	kubeconfigs []string
	path        string
	output      string
}

func (o *MergeOptions) encodeConfig(config *clientcmdapi.Config) ([]byte, error) {
	var err error
	var output []byte

	encode, err := runtime.Encode(clientcmdapilatest.Codec, config)
	if err != nil {
		return nil, err
	}

	const (
		configYAML = "yaml"
		configJSON = "json"
	)

	switch o.output {
	case configYAML:
		output, err = yaml.JSONToYAML(encode)
	case configJSON:
		output, err = yaml.YAMLToJSON(encode)
		output, err = utils.PrettyJson(output)
	default:
		err = fmt.Errorf("unsupported output type only save as yaml or json")
	}
	return output, err
}

func (o *MergeOptions) Merge() error {
	rules := clientcmd.ClientConfigLoadingRules{
		Precedence: o.kubeconfigs,
	}

	mergedConfig, err := rules.Load()
	if err != nil {
		return err
	}

	output, err := o.encodeConfig(mergedConfig)
	if err != nil {
		return err
	}

	if !o.overwrite {
		validations := map[string]int{
			"Contexts":  len(mergedConfig.Contexts),
			"Clusters":  len(mergedConfig.Clusters),
			"AuthInfos": len(mergedConfig.AuthInfos),
		}
		for k, v := range validations {
			if len(o.kubeconfigs) != v {
				return fmt.Errorf("merged config has conflict in %s. Use --overwrite to force merge", k)
			}
		}
	}

	if err := utils.WriteFile(o.path, output, 0664); err != nil {
		return err
	}
	return nil
}

func main() {

	kubestring := "/Users/wuchengjiang/Service/golang/go/packages/src/gin-vue-admin/server/kubeconfig/.k8sconfig.mul"

	t, _ := NewClient(kubestring)

	fmt.Println(t.GetCurrentContext())
	fmt.Println(t.GetConfigClusters())
	fmt.Println(t.GetConfigContext())

	// fmt.Println(string(ll["cluster1"].CertificateAuthorityData))
	// fmt.Println(ll["default"])

	t.SwitchContext("context-cluster1-admin")
	//t.SwitchContext("default")

	t, _ = NewClient(kubestring)

	fmt.Println("-----------------")
	fmt.Println(t.GetCurrentContext())
	fmt.Println(t.GetConfigClusters())
	fmt.Println(t.GetConfigContext())

	// namespace不指定获取所有Pod 列表
	pods, err := t.ClientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	//打印所有pod
	for _, pod := range pods.Items {
		fmt.Printf("Namespaces:%s, Name: %s, Status: %s, CreateTime: %s\n", pod.ObjectMeta.Namespace, pod.ObjectMeta.Name, pod.Status.Phase, pod.ObjectMeta.CreationTimestamp)
	}

	//merge config

	confighomepath := "/Users/wuchengjiang/Service/golang/go/packages/src/gin-vue-admin/server/kubeconfig/"

	var m MergeOptions
	m.backup = true

	var configpath []string
	configpath = []string{confighomepath + "config1", confighomepath + "config2"}

	m.kubeconfigs = configpath
	m.output = "yaml"
	m.overwrite = true
	m.path = confighomepath + "config3"

	err = m.Merge()
	if err != nil {
		fmt.Println(err.Error())
	}

}
