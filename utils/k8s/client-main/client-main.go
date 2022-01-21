package main

import "gin-vue-admin/service"

func main() {

	clustername := "k3s"

	service.SwitchCluster(clustername, "jun.huang")
	// t, _ = NewClient(kubestring)

	// // namespace不指定获取所有Pod 列表
	// pods, err := t.ClientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

	// if err != nil {
	// 	panic(err.Error())
	// }

	// //打印所有pod
	// for _, pod := range pods.Items {
	// 	fmt.Printf("Namespaces:%s, Name: %s, Status: %s, CreateTime: %s\n", pod.ObjectMeta.Namespace, pod.ObjectMeta.Name, pod.Status.Phase, pod.ObjectMeta.CreationTimestamp)
	// }

}
