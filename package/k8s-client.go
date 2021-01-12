package client

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

//InitClient will return kubectl client, the return is client set, is out side of cluster and init error
func InitClient() (*kubernetes.Clientset, bool, error) {
	// todo check is in or out side of cluster
	clientSet, err := outsideMode()
	return clientSet, true, err
}

//outsideMode will init kubernetes client out side of cluster
func outsideMode() (*kubernetes.Clientset, error) {
	var kubeConfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
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
