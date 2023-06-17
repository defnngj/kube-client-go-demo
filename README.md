## kube-client-go-demo

client-go 学习笔记


### kube 操作

```bash
> kubectl api-resources   # 查看所有API资源
> kubectl api-versions    # 查看所有API版本
```


### GRV 缓存

```bash
> cd ~/.kube/cache/discovery/127.0.0.1_6443
> ls
admissionregistration.k8s.io authorization.k8s.io         discovery.k8s.io             policy                       v1
apiextensions.k8s.io         autoscaling                  events.k8s.io                rbac.authorization.k8s.io
apiregistration.k8s.io       batch                        flowcontrol.apiserver.k8s.io scheduling.k8s.io
apps                         certificates.k8s.io          networking.k8s.io            servergroups.json
authentication.k8s.io        coordination.k8s.io          node.k8s.io                  storage.k8s.io
```


### web服务

* 启动服务

```bash
> go run main.go

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] GET    /pod/list                 --> main.main.func2 (3 handlers)
[GIN-debug] GET    /:resource/:group/:version --> main.main.func3 (3 handlers)
GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8888
```

