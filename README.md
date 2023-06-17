## kube-client-go-demo

client-go 学习笔记


## kube 操作

```bash
> kubectl api-resources   # 查看所有API资源
> kubectl api-versions    # 查看所有API版本
```


## GRV 缓存

```bash
> cd ~/.kube/cache/discovery/127.0.0.1_6443
> ls
admissionregistration.k8s.io authorization.k8s.io         discovery.k8s.io             policy                       v1
apiextensions.k8s.io         autoscaling                  events.k8s.io                rbac.authorization.k8s.io
apiregistration.k8s.io       batch                        flowcontrol.apiserver.k8s.io scheduling.k8s.io
apps                         certificates.k8s.io          networking.k8s.io            servergroups.json
authentication.k8s.io        coordination.k8s.io          node.k8s.io                  storage.k8s.io
```

