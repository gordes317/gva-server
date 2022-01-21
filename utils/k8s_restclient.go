package utils

import (
	"io/ioutil"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

//k8s restclient for exec terminal

func GetRestConf() (restConf *rest.Config, err error) {

	var kubeconfig []byte
	var clientErr error
	var restconfig *rest.Config

	if home := homedir.HomeDir(); home != "" {
		if kubeconfig, clientErr = ioutil.ReadFile(filepath.Join(home, ".kube", "config")); clientErr != nil {
			return
		}

		restconfig, clientErr = clientcmd.RESTConfigFromKubeConfig(kubeconfig)

		if clientErr != nil {
			panic(clientErr)
		}

	}
	return restconfig, err
}
