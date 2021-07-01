# 该仓库即将关闭，已经重新构造：Example.repo

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

# 服务Proxy
项目支持支持服务网格特性，
提供了构建负责服务调用关系的能力，
本质上利用服务转发的方式来描述服务之间的调用关系，
是一种模拟行为。

## 启用proxy的特性
配置文件中配置参数`mesh.enable_service_mesh=true`开启mesh特性路由。
在`mesh.server_hosts`中填写网格代理规则，系统将自动生成指定代理路由，进行模拟网格操作。

### server_host规则
`请求方法|模式|代理名称|目标主机`，
如果代理模式为`directly`，
并且请求方法不限制，
可以简写为`代理名称|目标主机`。

#### 支持的请求方法
- GET
- POST
- PUT
- PATCH
- DELETE
如果需要多个，使用逗号分隔，但是方法中不能含有空格。

例如需要支持 `post`和`get`请求:`post,get|...`
如果不填写，将默认赋值为全部请求方式。

#### 模式
- directly
- host-replace

directly模式：
在代理路由接收到请求后，直接发起目标方法到host地址，不对地址进行任何变更。
发送请求时将会复制`query`，`body`，`header`。

host-replace模式：
变更路由，提供复制path参数的能力。

如果当前路由名称为`proxy-1`，
则系统将会生成`/mesh/proxy-1`的路由，
请求`/mesh/proxy-1/test`，
将会转发请求到`/host/test`，
并且复制`query`，`body`，`header`。

以上两种方式不可混用，
如果配置文件中路由配置的模式无法识别，
将会成为空路由地址，
同样可以接受请求，
但不会产生任何作用。

#### name参数
由于项目采用了gin框架作为web服务框架，因此需要保证name可以被gin路由接受。
多个代理路由的name不可重复。
#### host参数
代理目标地址。

以`http://`或`https://`开头，且结尾处不允许存在`/`，
否则host-replace将会出现异常。

#### 返回数据说明
目标服务必须返回一个json数据，否则服务将无法解析。
返回体中将会包装代理目标的返回体和返回头。
如：
```json
{
    "method": "GET",
    "application_name": "server1",
    "object": {
        "header": {
            "App_name": [
                "k8s-test-server"
            ],
            "App_version": [
                "v0.0.45"
            ],
            "Content-Length": [
                "95"
            ],
            "Content-Type": [
                "application/json; charset=utf-8"
            ],
            "Date": [
                "Wed, 24 Feb 2021 19:25:38 GMT"
            ]
        },
        "body": {
            "application_name": "k8s-test-server",
            "method": "GET",
            "object": "2021-02-24T19:25:38.717285788Z"
        }
    }
}
```

#### 链路闭环
利用本项目提供的`common-request` 和 `common-cache`进行模式调用，
可以实现具有一定逻辑能力的服务网格，
如果需要使用istio特性，目标路由地址使用`k8s-dns`格式：
`http://serverName.namespace.svc.cluster.local`即可。

而且通过配置文件中的run_name可以提供服务名称模拟。

通过以上配置即可实现大部分的网格特性测试。 
