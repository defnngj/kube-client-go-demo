package main

import (
	"fmt"
	"kube-client-go-demo/pkg/kubeconfig"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery/cached/disk"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	fmt.Println("cache DiscoveryClient use!")

	kubeconfig := kubeconfig.ConfigPath()

	// 1. 加载配置文件，生成config 对象
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("connect  error")
	}

	// 2. 实例化客户端，将GVR 数据缓存到本地文件中
	cacheDiscoveryClient, err := disk.NewCachedDiscoveryClientForConfig(config, "./cache/discovery", "./cache/http", time.Minute*60)
	if err != nil {
		panic(err.Error())
	}

	/*
	 先从缓冲文件中招 GVR 数据，有则直接返回，没有则需要调用 APIServer
	 调用 APIServer 获取 GVR 数据。
	 将获取的 GVR 数据缓存到本地， 然后返回给客户端
	*/
	_, apiResources, err := cacheDiscoveryClient.ServerGroupsAndResources()
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
