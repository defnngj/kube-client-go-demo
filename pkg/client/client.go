package client

import (
	"fmt"
	"kube-client-go-demo/pkg/kubeconfig"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var config *rest.Config

type Clients struct {
	clientSet kubernetes.Interface
}

func NewClients() (clients Clients) {

	var (
		err error
	)

	kubeconfig := kubeconfig.ConfigPath()
	// 1. 加载配置，生成配置文件对象
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Printf("读取配置错误，%+v", err)
		return
	}

	// 2. 实例化各客户端
	clients.clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("实例化客户端错误, %v", err)
		return
	}

	return
}

func (c *Clients) ClientSet() kubernetes.Interface {
	return c.clientSet
}

func GetConfig() *rest.Config {
	return config
}
