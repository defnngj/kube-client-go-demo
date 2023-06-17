package informer

import (
	"kube-client-go-demo/pkg/client"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
)

var sharedInformerFactory informers.SharedInformerFactory

func NewSharedInformerFactory(stopCh <-chan struct{}) (err error) {
	var (
		clients client.Clients
	)

	// 1. 加载客户端
	clients = client.NewClients()

	// 2. 实例化 sharedInformerFactory
	sharedInformerFactory = informers.NewSharedInformerFactory(clients.ClientSet(), time.Second*60)

	// 3. 启动informer
	gvrs := []schema.GroupVersionResource{
		{Group: "", Version: "v1", Resource: "pods"},
		{Group: "", Version: "v1", Resource: "services"},
		{Group: "", Version: "v1", Resource: "namespaces"},

		{Group: "apps", Version: "v1", Resource: "deployments"},
		{Group: "apps", Version: "v1", Resource: "statefulsets"},
		{Group: "apps", Version: "v1", Resource: "daemonsets"},
	}

	for _, v := range gvrs {
		// 创建 informer
		_, err = sharedInformerFactory.ForResource(v)
		if err != nil {
			return
		}
	}

	sharedInformerFactory.Start(stopCh)
	sharedInformerFactory.WaitForCacheSync(stopCh)

	return
}

func Get() informers.SharedInformerFactory {
	return sharedInformerFactory
}

// informer 初始化函数
func Setup(stopCh <-chan struct{}) (err error) {
	err = NewSharedInformerFactory(stopCh)
	if err != nil {
		panic(err)
	}
	return
}
