# k8s测试服务-后端

对k8s特性进行测试的服务后端

## 路由列表

### [GET]/hello
验证服务是否正常，返回字符串`hello world`

### [GET]/podenv
获取容器部分环境变量

- SERVICE_IP
- SERVICE_NAME
- SERVICE_NAMESPACE 

### [GET]/env
获取当前环境下所有的环境变量

## ENVs - 环境变量声明
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