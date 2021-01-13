package main

import (
	"github.com/sirupsen/logrus"
	"k8s-test-backend/internal/server"
	client "k8s-test-backend/package"
	"k8s.io/client-go/kubernetes"
	"os"
)

var Version = ""
var BuildPlatform = ""
var BuildStamp = ""

var IsInCluster = false
var ClusterSet *kubernetes.Clientset = nil

// main will start application
func main() {

	// all the argument with application will output version info.
	if len(os.Args) > 1 {
		showVersion()
	} else {
		showVersion()
		server.Start(":3000")
	}
}

// showVersion will print version info
func showVersion() {
	logrus.Infoln("The version is :", Version)
	logrus.Infoln("The build from :", BuildPlatform)
	logrus.Infoln("The build stamp:", BuildStamp)
}

func init() {
	client, isInCluster, err := client.InitClient()
	if err != nil {
		logrus.Error(err)
		ClusterSet = nil
		return
	}
	IsInCluster = isInCluster
	ClusterSet = client
	logrus.Infoln("Init kubernetes cluster success, the mode is:(false : out side of cluster, true: in side of cluster) ", isInCluster)
}
