package conf

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
)

const (
	applicationName = "k8s-test-server" // the application name.
	version         = "0.0.1"           // the application version.

	// config-file
	defaultConfigFile = "temp.yaml" // default config name is `k8s-test-server.yaml`, and search it in config path.

	// envs
	envPrefix = "KTS" // the environment's name, short of k8s-test-server
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
	EnableServerFeature:  false,
	ApplicationRunName:   applicationName, // the application run time name, default is application name

	// the config or command can't change of config struct field. Get value from envs.
	ServiceIp:        "",
	ServiceName:      "",
	ServiceNamespace: "",

	// the application constant value
	Version:       version,
	BuildPlatform: "",
	BuildStamp:    "",

	// program create
	KubeClientSet:     nil,
	ServiceMeshMapper: make([]serviceMeshMapper, 0),

	// constant value of application
	CApplicationName:   applicationName,
	CDefaultConfigFile: defaultConfigFile,
	CEnvPrefix:         envPrefix,
}

type config struct {
	// const value here
	CApplicationName   string // the application name
	CDefaultConfigFile string // the application default config file
	CEnvPrefix         string // the env of k8s-test-server prefix

	// server
	Port               string // the application port(default should be :3000).
	Mode               string // the application run mode(for gin, default is debug).
	LogPath            string // the log file path(default should be ./log.log).
	ApplicationRunName string // the application runtime name, use to common request's response

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

	// mesh feature
	EnableServerFeature bool                // whether to enable server feature, default should be false
	ServiceMeshMapper   []serviceMeshMapper // the mesh mapper list, only get value from config
}

func InitMeshMapper(str string) serviceMeshMapper {
	return serviceMeshMapper{Str: str}
}

// serviceMeshMapper package the service route info
type serviceMeshMapper struct {
	name *string // if name is empty, skip this value
	host *string // host cloud be empty, it mean return value directly
	//methodList []string
	//model      []string
	Str string
}

func (s *serviceMeshMapper) GetName() string {
	if s.name == nil {
		equalsIndex := strings.Index(s.Str, "=")
		if equalsIndex == -1 {
			s.name = &s.Str
		} else {
			tempValue := s.Str[0:equalsIndex]
			s.name = &tempValue
		}
	}
	return *s.name
}

func (s *serviceMeshMapper) GetHost() string {
	if s.host == nil {
		equalsIndex := strings.Index(s.Str, "=")
		tempValue := ""
		if equalsIndex != -1 {
			tempValue = s.Str[equalsIndex+1:]
		}
		s.host = &tempValue
	}
	return *s.host
}
