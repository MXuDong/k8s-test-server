# k8s测试服务-后端

对k8s特性进行测试的服务后端。

## 路由列表

#### [GET]/hello
验证服务是否正常，返回字符串`hello world`。

#### [GET]/env/pod
获取容器部分环境变量

- SERVICE_IP
- SERVICE_NAME
- SERVICE_NAMESPACE 

#### [GET]/env
获取当前环境下所有的环境变量。

## Params - 应用启动参数说明
```
-ginMode string
    The mode of gin. (default "debug")
    
-kubeconfig string
    (optional) absolute path to the kubeconfig file (default "/Users/mengxudong/.kube/config")
    
-logPath string
    The log file path. (default "log.log")
    
-v    Show version info, if true, it will not start server.
```

## ENVs - 环境变量声明
### USE_KUBE_FEATURE
是否启用应用对Kube的使用，如果启用，应用将获取k8s集群信息。

可选项：
- true: 启用
- false: 不启用

如果不是true的任何选项，包括不设置均认为不启用对接k8s集群的功能。

    docker环境内为 false
### IS_IN_CLUSTER
是否为kubernetes环境内，该环境变量将影响系统初始化`k8s-clientSet`的方式。

需要开启环境变量`USE_KUBE_FEATURE`才能正常使用该功能。

可选项：
- true: 集群内启动
- false: 集群外启动

如果不是true的任何选项，包括不设置均认为集群外启动，会尝试加载kubeConfig。
系统将会从`home/.kube/config`进行查找，或者通过通过指定`kubeconfig`参数进行查找。

    deocker环境内为 false
### SERVICE_IP
应用服务地址 

    docker环境内为 in-docker
### SERVICE_NAME
应用服务名称

    docker环境内为 in-docker 
### SERVICE_NAMESPACE
k8s环境内命名空间

    docker环境内为 in-docker
    
## k8s 编排文件

[k8s-deploy.yaml](k8s-deploy.yaml)

## 编译
```
$ go get -d -v ./...
$ go build -o app cmd/main.go
```