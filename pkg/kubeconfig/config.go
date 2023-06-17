package kubeconfig

import (
	"flag"
	"fmt"
	"path/filepath"
)

// HomeDir 用户根目录，-需修改
func HomeDir() string {
	return "/Users/fnngj/"
}

// apiServer 地址 查看 ~/.kube/config 文件
func KubeHost() string {
	return "127.0.0.1:6443"
}

// 返回 kube 配置文件路径
func ConfigPath() string {

	var kubeconfig *string

	// 默认会从~/.kube/config路径下获取配置文件
	if home := HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional)absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	fmt.Println("kube配置地址:", *kubeconfig)

	return *kubeconfig
}
