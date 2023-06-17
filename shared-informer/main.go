package main

import (
	"fmt"
	"kube-client-go-demo/pkg/kubeconfig"
	"time"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	fmt.Println("SharedInformer use!")

	kubeconfig := kubeconfig.ConfigPath()

	// 1. 加载配置文件，生成config 对象
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("connect  error")
	}

	// 2. 实例化 ClientSet 对象
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 初始化 SharedInformerFactory
	sharedInformerFactory := informers.NewSharedInformerFactory(clientSet, 0)

	// 查询Pod数据，生成 PodInformer 对象
	podInformer := sharedInformerFactory.Core().V1().Pods()

	// 生成一个 indexer, 便于数据的查询
	indexer := podInformer.Lister()

	// 启动 informer
	sharedInformerFactory.Start(nil)

	// 等待数据同步完成
	sharedInformerFactory.WaitForCacheSync(nil)

	// 利用 indexer 获取数据
	for i := range time.Tick(time.Second * 5) {
		fmt.Println("nginx-心跳:", i)
		pods, err := indexer.List(labels.Everything())
		if err != nil {
			panic(err.Error())
		}

		// 打印所有获取到的pods资源，输出到标准输出
		for _, item := range pods {
			fmt.Printf("NAMESPACE: %v NAME: %v \t STATUS: %v \n", item.Namespace, item.Name, item.Status.Phase)
		}
	}

}

// 运行过滤 nginx 关键字
// go run main.go | grep "nginx"
