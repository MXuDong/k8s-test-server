package server

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var Config = globalConfig{
	UseKubeFeature:  false,
	IsInSideCluster: false,
	KubeClientSet:   nil,
	LogPath:         "log.log",
	GinMode:         gin.DebugMode,
}

// the config center
type globalConfig struct {
	UseKubeFeature  bool                  // is can use kube feature
	IsInSideCluster bool                  // the mode of application
	KubeConfigPath  string                // the kube config path
	KubeClientSet   *kubernetes.Clientset // the kube client, if `UseKubeFeature` is false, it will nil
	KubeConfig      *rest.Config          // the kube config

	LogPath string // the log path

	GinMode string // the gin mode, support: release, debug, test
}
