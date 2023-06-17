package main

import (
	"context"
	"fmt"
	"kube-client-go-demo/pkg/kubeconfig"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// HomeDir 项目跟目录
func HomeDir() string {
	return "/Users/fnngj/"
}

func main() {
	fmt.Println("RestClient use!")

	kubeconfig := kubeconfig.ConfigPath()

	// 1. 加载配置文件，生成config 对象
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("connect  error")
	}

	// 2. 配置API路径
	config.APIPath = "api"
	// 3. 配置分组版本
	config.GroupVersion = &corev1.SchemeGroupVersion
	// 4. 配置数据的编解码工具
	config.NegotiatedSerializer = scheme.Codecs

	// 5. 实例化 RESTClient 对象
	restClient, err := rest.RESTClientFor(config)

	// 定义接受返回值的变量
	result := &corev1.PodList{}
	// 跟 阿皮server交互
	err = restClient.
		Get(). // get 请求方式
		Namespace("kube-system").
		Resource("pods").                                                        // 制定要访问的资源
		VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec). // 参数以及参数序列化
		Do(context.TODO()).                                                      //触发请求
		Into(result)                                                             // 写入返回结果

	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(result)
	// 打印所有获取到的pods资源，输出到标准输出
	for _, d := range result.Items {
		fmt.Printf("NAMESPACE: %v NAME: %v \t STATUS: %v \n", d.Namespace, d.Name, d.Status.Phase)
	}
}
