package utils

import (
	"context"
	"gin-vue-admin/global"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"strings"
)

func GetKubeVersion(userName string) int {
	clientset, err := kubernetes.NewForConfig(global.GetK8sConfig(userName))
	if err != nil {
		panic(err.Error())
	}

	podsClient := clientset.CoreV1().Nodes()
	result, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	kubeVersionString := result.Items[0].Status.NodeInfo.KubeletVersion[1:5]
	kubeVersionSlice := strings.Split(kubeVersionString, ".")
	kubeVersionString = strings.Join(kubeVersionSlice, "")
	kubeVersionInt, err := strconv.Atoi(kubeVersionString)
	if err != nil {
		panic(err)
	}

	return kubeVersionInt
}
