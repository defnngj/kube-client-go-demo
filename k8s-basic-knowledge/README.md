# K8s 基础知识


参考：https://blog.csdn.net/yhj_911/article/details/124702851

参考：https://kubernetes.io/zh-cn/docs/tasks/run-application/run-stateless-application-deployment/

### 命名空间 Namespace

Namespace 表示命名空间

```bash
kubectl get namespace   # 查看所有命名
kubectl describe namespace nginx  # 查看命名空间详情
```

* 创建命名空间

```bash
kubectl create -f nginx-namespace.yaml  # 创建命名空间
```

### 部署 deployment

deplyment 表示pod发布

```bash
kubectl get deploy  # 查看部署
kubectl get deploy -n nginx  # 指定命名空间
```

* 创建deploy

```bash
kubectl create -f nginx-deployment.yaml  # 创建deployment
kubectl apply  -f nginx-deployment.yaml  # 修改后重新执行
kubectl edit deploy nginx-deployment1 -n nginx  # 直接修改
kubectl delete deployment nginx-deployment1 -n nginx  #删除
```


### service

Service 表示多个pod做为一组的集合对外通过服务的表示

查看服务

```bash
kubectl get service 
```

* 创建 service

```bash
kubectl create -f nginx-service.yaml
```

### pod 

查看pod

```bash
kubectl get pod -n nginx  # 查看pod
kubectl delete pod -n nginx <pod名>   # 删除pod
kubectl -n nginx get pod -o yaml <pod名>  # 查看pod 日志
kubectl get deploy nginx-deployment1 -o yaml -n nginx | grep resourceVersion   # 查看版本
```


## docker 代理配置

参考：https://zhuanlan.zhihu.com/p/146876547

```json
{
  "builder": {
    "gc": {
      "defaultKeepStorage": "20GB",
      "enabled": true
    }
  },
  "experimental": false,
  "features": {
    "buildkit": true
  }
}
```

替换为

```json
{
  "debug": true,
  "experimental": true,
  "registry-mirrors": [
    "http://hub-mirror.c.163.com"
  ]
}
```