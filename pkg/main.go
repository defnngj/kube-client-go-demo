package main

import (
	"fmt"
	"kube-client-go-demo/pkg/informer"

	"k8s.io/apimachinery/pkg/labels"
)

/*
封装 cloent-go 客户端
*/
func main() {
	stopCh := make(chan struct{})
	// 进行了 client-go 的初始化
	err := informer.Setup(stopCh)
	if err != nil {
		panic(err.Error())
	}

	items, err := informer.Get().Core().V1().Pods().Lister().List(labels.Everything())
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		fmt.Printf("NAMESPACE: %v NAME: %v \t STATUS: %v \n", item.Namespace, item.Name, item.Status.Phase)
	}
}
