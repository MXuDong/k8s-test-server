package client

import (
	"k8s-test-backend/internal/server"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

const (
	EnvIsInCluster = "IS_IN_CLUSTER"

	InCluster = "true"
)

//InitClient will return kubectl client, the return is client set, is out side of cluster and init error
func InitClient() (*kubernetes.Clientset, bool, error) {
	isInCluster := IsInKubernetes()
	if isInCluster {
		clientSet, err := insideMode()
		return clientSet, isInCluster, err
	} else {
		clientSet, err := outsideMode()
		return clientSet, isInCluster, err
	}
}

// IsInKubernetes will return the application is run in kubernetes, it get value from env : EnvIsInCluster
func IsInKubernetes() bool {
	return os.Getenv(EnvIsInCluster) == InCluster
}

//outsideMode will init kubernetes client out side of cluster
func outsideMode() (*kubernetes.Clientset, error) {

	config, err := clientcmd.BuildConfigFromFlags("", server.Config.KubeConfig)
	if err != nil {
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientSet, nil
}

// insideMode will init kubernetes client in side of cluster
func insideMode() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientSet, nil
}
