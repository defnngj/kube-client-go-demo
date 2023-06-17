package kubeconfig

import (
	"flag"
	"fmt"
	"path/filepath"
)

// HomeDir 项目跟目录
func HomeDir() string {
	return "/Users/fnngj/"
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
