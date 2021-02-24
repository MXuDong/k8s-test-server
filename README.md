# k8s测试服务-后端

![kts](./docs/kts.png)


对k8s特性进行测试的服务后端。

## Params - 应用启动参数说明
```
Flags:
      --config_path string              the application run config file path
      --enable_cache_http               whether use cache http handle, default is true (default true)
      --enable_common_http              whether use common http handle, default is true (default true)
      --enable_kubernetes_feature       whether enable kubernetes feature, default is false
      --enable_service_mesh             whether the enable service mesh
  -h, --help                            help for k8s-test-server
      --is_in_kubernetes                whether the application in kubernetes cluster as the pods
      --kubernetes_config_path string   the config path of kubernetes (default "/Users/mengxudong/.kube/config")
      --log_path string                 the file of log output (default "log.log")
      --mode string                     the application run mode, in 'debug', 'release', 'test' (default "debug")
      --port string                     the application start port, default is :3000 (default ":3000")
  -v, --version                         show the version of application
```

## 配置文件
支持 yaml、toml、json等主流配置文件类型

[示例文件 k8s-test-server-config.yaml](k8s-test-server-config.yaml)

## ENVs - 环境变量声明

所有参数均可以转换为环境变量，大写后，前缀为`KTS_`，具体参数参配置文件即可

参数优先级：控制台启动指令>环境变量>配置文件>默认启动参数


### KTS_CONFIG_PATH
配置文件路径

### KTS_ENV_SERVICE_IP
应用服务地址 

    docker环境内为 in-docker
### KTS_ENV_SERVICE_NAME
应用服务名称

    docker环境内为 in-docker 
### KTS_ENV_SERVICE_NAMESPACE
k8s环境内命名空间

    docker环境内为 in-docker
    
## k8s 编排文件

[示例文件 k8s-deploy.yaml](k8s-deploy.yaml)

## 编译
### 源码直接编译
```
$ go get -d -v ./...
$ go build -o app cmd/main.go
```

### Makefile编译镜像
```
$  sh LocalBuild.sh
```

### Makefile 编译为本地执行包
```
$ make build
```

### Docker 直接打包镜像
```
$ docker build -t A-Donga/k8s-common-test:0.0.1 .
```