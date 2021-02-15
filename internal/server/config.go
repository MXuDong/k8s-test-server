package server

import (
	"github.com/gin-gonic/gin"
	"k8s-test-backend/conf"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var Config = globalConfig{
	UseKubeFeature:  false,
	IsInSideCluster: false,
	KubeClientSet:   nil,
	LogPath:         conf.LogFilePath,
	GinMode:         gin.DebugMode,
	ApplicationPort: conf.ServicePort,
}

// the config center
type globalConfig struct {
	UseKubeFeature  bool                  // is can use kube feature
	IsInSideCluster bool                  // the mode of application
	KubeConfigPath  string                // the kube config path
	KubeClientSet   *kubernetes.Clientset // the kube client, if `UseKubeFeature` is false, it will nil
	KubeConfig      *rest.Config          // the kube config

	LogPath string // the log path

	GinMode         string // the gin mode, support: release, debug, test
	ApplicationPort string // the application listen port, it will set to gin server
}
