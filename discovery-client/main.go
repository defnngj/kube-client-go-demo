package main

import (
	"fmt"
	"kube-client-go-demo/pkg/kubeconfig"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	fmt.Println("DiscoveryClient use!")

	kubeconfig := kubeconfig.ConfigPath()

	// 1. 加载配置文件，生成config 对象
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("connect  error")
	}

	// 2. 实例化 DiscoveryClient 对象
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3. 发送请求，获取GVR 数据
	_, apiResources, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err.Error())
	}

	for _, list := range apiResources {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err.Error())
		}
		for _, resource := range list.APIResources {
			fmt.Printf("NAME: %v GROUP: %v \t VERSION: %v \n", resource.Name, gv.Group, gv.Version)
		}

	}

}
