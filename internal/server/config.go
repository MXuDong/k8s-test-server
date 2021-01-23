package server

import "k8s.io/client-go/kubernetes"

var Config = globalConfig{
	UseKubeFeature:  false,
	IsInSideCluster: false,
	KubeClientSet:   nil,
}

// the config center
type globalConfig struct {
	UseKubeFeature  bool
	IsInSideCluster bool
	KubeClientSet   *kubernetes.Clientset

	logPath  string
	logLevel string
}
