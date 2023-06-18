## kubebuilder 使用

Kubebuilder 是对 k8s client-go 提供的再封装，并且提供了一个脚手架项目。


官方文档：https://book.kubebuilder.io/quick-start.html

中文文档：https://cloudnative.to/kubebuilder/

### CRD 编程


熟悉 k8s 的同学都知道，k8s 中定义了一系列的资源对象，比如 pod，deployment,svc 等等，使用 `kubectl api-resources` 可以看到当前版本的 k8s 所支持的全部资源以及分组等信息

```bash
kubectl api-resources
```

在这些资源中有一种 custom resource definitions （自定义资源定义） 的对象简称 CRD。`kubectl get crd -A` 命令查看 CRD 资源。

* 查看 CRD 资源

```bash
kubectl get crd -A
```


### 安装

* 查看go 环境

```bash
go env

GOOS="darwin"
GOARCH="amd64"
...
```

* 根据go 环境下载对应的版本

```bash
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$darwin/$amd64
```

* 增加执行权限，移到指定 `bin/` 目录。

```bash
chmod +x kubebuilder
mv kubebuilder /usr/local/bin/
```

### 创建项目

#### 1.初始化项目

```bash
kubebuilder init --domain=kruise.io
kubebuilder init --domain my.domain --repo my.demain/guestbook
```

* tree

```shell
.
├── Dockerfile
├── Makefile
├── PROJECT
├── README.md
├── cmd
│   └── main.go
├── config
│   ├── default
│   │   ├── kustomization.yaml
│   │   ├── manager_auth_proxy_patch.yaml
│   │   └── manager_config_patch.yaml
│   ├── manager
│   │   ├── kustomization.yaml
│   │   └── manager.yaml
│   ├── prometheus
│   │   ├── kustomization.yaml
│   │   └── monitor.yaml
│   └── rbac
│       ├── auth_proxy_client_clusterrole.yaml
│       ├── auth_proxy_role.yaml
│       ├── auth_proxy_role_binding.yaml
│       ├── auth_proxy_service.yaml
│       ├── kustomization.yaml
│       ├── leader_election_role.yaml
│       ├── leader_election_role_binding.yaml
│       ├── role_binding.yaml
│       └── service_account.yaml
├── go.mod
├── go.sum
└── hack
    └── boilerplate.go.txt
```

#### 2.创建API

```bash
kubebuilder create api --group apps --version v1alpha1 --kind SidecarSet --namespaced=false
kubebuilder create api --group webapp --version v1 --kind Guestbook
Create Resource [y/n]
y
Create Controller [y/n]
y
...


```

__参数说明__

* group + 前面的 domain, 为此CRD的 group: apps.kruise.io
* version 一般为三种，按社区标准：
  * v1alpha1: 不稳定版本
  * v1beta1: 已稳定，特性可能调整
  * v1: 稳定版本
* kind: 此 CRD 的类型，类似于社区原生的 Service 的概念
* namespaced: 此 CRD 是全局唯一还是namespace唯一，类似node和pod

__效果解读__

* 生成了 CRD 和 controller 的框架，后面需要手工填充代码

__目录树__

* tree

```bash
.
├── Dockerfile
├── Makefile
├── PROJECT
├── README.md
├── api
│   └── v1alpha1
│       ├── groupversion_info.go
│       ├── sidecarset_types.go
│       └── zz_generated.deepcopy.go
├── bin
│   ├── controller-gen
│   └── kustomize
├── cmd
│   └── main.go
├── config
│   ├── crd
│   │   ├── bases
│   │   │   └── apps.kruise.io_sidecarsets.yaml
│   │   ├── kustomization.yaml
│   │   ├── kustomizeconfig.yaml
│   │   └── patches
│   │       ├── cainjection_in_sidecarsets.yaml
│   │       └── webhook_in_sidecarsets.yaml
│   ├── default
│   │   ├── kustomization.yaml
│   │   ├── manager_auth_proxy_patch.yaml
│   │   └── manager_config_patch.yaml
│   ├── manager
│   │   ├── kustomization.yaml
│   │   └── manager.yaml
│   ├── prometheus
│   │   ├── kustomization.yaml
│   │   └── monitor.yaml
│   ├── rbac
│   │   ├── auth_proxy_client_clusterrole.yaml
│   │   ├── auth_proxy_role.yaml
│   │   ├── auth_proxy_role_binding.yaml
│   │   ├── auth_proxy_service.yaml
│   │   ├── kustomization.yaml
│   │   ├── leader_election_role.yaml
│   │   ├── leader_election_role_binding.yaml
│   │   ├── role.yaml
│   │   ├── role_binding.yaml
│   │   ├── service_account.yaml
│   │   ├── sidecarset_editor_role.yaml
│   │   └── sidecarset_viewer_role.yaml
│   └── samples
│       ├── apps_v1alpha1_sidecarset.yaml
│       └── kustomization.yaml
├── go.mod
├── go.sum
├── hack
│   └── boilerplate.go.txt
└── internal
    └── controller
        ├── sidecarset_controller.go
        └── suite_test.go
```

### 3. 测试一下

* 将CRD 安装到集群中

```bash
make install
```

* 运行你的控制器

```bash
make run
```


### 4. 接下来

* 安装 CR 实例

```bash
kubectl apply -f config/samples/
```

* 构建并推送

```bash
make docker-build docker-push IMG=harbor.dmcca.loc:10000/zxl-test/guestbook-sample:v1
```

* 根据 IMG 指定的景象将控制器部署到集群中。

```bash
make deploy IMG=harbor.dmcca.loc:10000/zxl-test/guestbook-sample:v1
``` 

* 卸载

```bash
make uninstall   # 从集群中卸载CRD
make undeploy    # 从集群中卸载 控制器
```

## 帮助文档

文章：
https://www.cnblogs.com/alisystemsoftware/p/11580202.html

视频
https://www.kubesphere.io/zh/live/uisee0923-live/
https://www.bilibili.com/video/BV1x84y157UX/?spm_id_from=autoNext&vd_source=3a823453217bd1790e9e4293cb86b6df

示例：
https://github.com/schwarzeni/kubebuilder-imoocpod
https://github.com/a772304419/31-kubebuilder-mysql-operator
