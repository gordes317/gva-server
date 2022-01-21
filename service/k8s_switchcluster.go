package service

import (
	"fmt"
	"gin-vue-admin/global"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

func SwitchCluster(clustername, userName string) (err error) {

	//switch cluster

	global.GVA_K8SCurrent = clustername

	//获取当前集群 只能从客户端得到

	var clientErr error
	var config string

	rows, err := global.GVA_DB.Raw("select config from multicluster where name = ?", clustername).Rows()

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&config)
	}
	fmt.Println("UserName:", userName)

	global.GVA_K8SConfig, clientErr = clientcmd.RESTConfigFromKubeConfig([]byte(config))

	//修改全局 gloal config

	if clientErr != nil {
		panic(clientErr.Error())
	} else {
		fmt.Printf("切换K8S集群: %s  success\n", clustername)
	}
	global.GVA_K8SConfigMap[userName] = global.GVA_K8SConfig
	for s, _ := range global.GVA_K8SConfigMap {
		fmt.Println("userName:", s)
	}

	return nil

}

