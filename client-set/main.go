package main

import (
	"context"
	"fmt"
	"kube-client-go-demo/pkg/kubeconfig"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	fmt.Println("ClientSet use!")

	kubeconfig := kubeconfig.ConfigPath()

	// 1. 加载配置文件，生成config 对象
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("connect  error")
	}

	// 2. 实例化 ClientSet 对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.
		CoreV1().                                  // 返回 CoreV1Client 实例
		Pods("kube-system").                       // 指定查询的资源以及指定的 namaspace, 如果 namespace 为空，则表示插叙所有人的 namespace
		List(context.TODO(), metav1.ListOptions{}) // 在这里就表示查询Pod列表
	if err != nil {
		panic(err.Error())
	}

	// 打印所有获取到的pods资源，输出到标准输出
	for _, d := range pods.Items {
		fmt.Printf("NAMESPACE: %v NAME: %v \t STATUS: %v \n", d.Namespace, d.Name, d.Status.Phase)
	}
}
