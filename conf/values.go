package conf

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	ApplicationName = "k8s-test-server" // the application name.
	version         = "0.0.1"           // the application version.

	// paths

	// config-file
	DefaultConfigFile = "temp.yaml" // default config name is `k8s-test-server.yaml`, and search it in config path.

	// envs
	EnvPreFix = "KTS" // the environment's name, short of k8s-test-server
)

var ApplicationConfig = config{
	// the config or command can change of config struct.
	Port:                 ":3000",
	Mode:                 "debug",
	LogPath:              "log.log",
	UseCommonHttp:        true,
	UseCacheHttp:         true,
	UseKubernetesFeature: false,
	KubernetesConfigPath: "",
	IsInCluster:          false,

	// the config or command can't change of config struct field.
	ServiceIp:        "",
	ServiceName:      "",
	ServiceNamespace: "",

	// the application constant value
	Version:       version,
	BuildPlatform: "",
	BuildStamp:    "",

	// program create
	KubeClientSet: nil,
}

type config struct {
	// server
	Port    string // the application port(default should be :3000).
	Mode    string // the application run mode(for gin, default is debug).
	LogPath string // the log file path(default should be ./log.log).

	// common http bin
	UseCommonHttp bool // whether to enable common http bin handle, default should be true.
	UseCacheHttp  bool // whether to enable cache http bin handle, default should be true.

	// build args
	Version       string // the version of application
	BuildPlatform string // the build platform of application
	BuildStamp    string // the build timestamp of application

	// kubernetes feature
	UseKubernetesFeature bool                  // whether to enable k8s feature service.
	KubernetesConfigPath string                // the kubernetes config.
	IsInCluster          bool                  // whether application in k8s cluster(as a pod)
	ServiceIp            string                // the pod in kubernetes ip, it will get from k8s.
	ServiceName          string                // the pod name.
	ServiceNamespace     string                // the pod's namespace.
	KubeClientSet        *kubernetes.Clientset // the kubernetes client set
	KubeClientConf       *rest.Config          // the kubernetes config
}
