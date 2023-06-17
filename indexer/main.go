package main

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

/*
IndexFunc: 索引器函数，用于计算一个资源对象的索引值列表，可以根据需求定义其他的，比如根据 Label 标签，Annotation 等属性来生成索引值列表。
Index: 存储数据，要查找某个命名空间下面的 pod, 那就要让pod 按照其命名空间进行索引，对应的Index类型就是 map[namespace]sets.pod。
Indexers: 存储索引器，key 为索引器名称，value 为索引器的实现寒素，例如 map["namespace"]MetaNamespaceIndexFunc。
indices: 存储索引器，key 为索引名称，value 为缓存的数据，例如： map["namespace"]map[namespace]sets.pod。
*/

/*
 实现两个索引器函数，分别基于 Namespace、NodeName， 资源对象Pod
*/

func NamespaceIndexFunc(obj interface{}) (result []string, err error) {
	pod, ok := obj.(*v1.Pod)
	if !ok {
		return nil, fmt.Errorf("类型错误， %v", err)
	}

	result = []string{pod.Namespace}
	return
}

func NodeNameIndexFunc(obj interface{}) (result []string, err error) {
	pod, ok := obj.(*v1.Pod)
	if !ok {
		return nil, fmt.Errorf("类型错误， %v", err)
	}
	// 索引值列表
	result = []string{pod.Spec.NodeName}
	return
}

func main() {
	// 实现一个Indexer 对象
	index := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{
		"namespace": NamespaceIndexFunc,
		"nodeName":  NodeNameIndexFunc,
	})

	// 模拟数据
	pod1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "index-pod-1",
			Namespace: "default",
		},
		Spec: v1.PodSpec{
			NodeName: "node1",
		},
	}

	pod2 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "index-pod-2",
			Namespace: "default",
		},
		Spec: v1.PodSpec{
			NodeName: "node2",
		},
	}

	pod3 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "index-pod-3",
			Namespace: "kube-system",
		},
		Spec: v1.PodSpec{
			NodeName: "node2",
		},
	}

	// 将数据写入到Indexer中
	_ = index.Add(pod1)
	_ = index.Add(pod2)
	_ = index.Add(pod3)

	// 通过索引器函数查询一下数据
	fmt.Println("---------namespace 查询----------")
	pods, err := index.ByIndex("namespace", "default")
	if err != nil {
		panic(err)
	}

	for _, pod := range pods {
		fmt.Println(pod.(*v1.Pod).Name)
	}

	fmt.Println("---------nodeName 查询----------")
	pods, err = index.ByIndex("nodeName", "node2")
	if err != nil {
		panic(err)
	}

	for _, pod := range pods {
		fmt.Println(pod.(*v1.Pod).Name)
	}
}
