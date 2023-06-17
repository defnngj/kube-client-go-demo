package main

import (
	"context"
	"fmt"
	"kube-client-go-demo/pkg/kubeconfig"

	runtime "k8s.io/apimachinery/pkg/runtime"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	fmt.Println("DynameicClient use!")

	kubeconfig := kubeconfig.ConfigPath()

	// 1. 加载配置文件，生成config 对象
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("connect  error")
	}

	// 2. 实例化 DynameicClient 对象
	dynameicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3. 配置 GVR
	gvr := schema.GroupVersionResource{
		Group:    "", // 不需要写，因为无名资源组，也就是core 资源组
		Version:  "v1",
		Resource: "pods",
	}

	// 4. 发送请求，得到返回结果
	unStructData, err := dynameicClient.Resource(gvr).Namespace("kube-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// 5. unStructData 转化为结构化的数据
	podList := &corev1.PodList{}
	if err = runtime.DefaultUnstructuredConverter.FromUnstructured(unStructData.UnstructuredContent(), podList); err != nil {
		fmt.Println("结构化失败", err)
	}

	// 打印所有获取到的pods资源，输出到标准输出
	for _, d := range podList.Items {
		fmt.Printf("NAMESPACE: %v NAME: %v \t STATUS: %v \n", d.Namespace, d.Name, d.Status.Phase)
	}
}
