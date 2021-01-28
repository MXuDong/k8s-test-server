package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"io"
	"k8s-test-backend/conf"
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

const (
	EnvUseKubeFeature = "USE_KUBE_FEATURE"
	UseKubeFeature    = "true"
)

// main will start application
func main() {

	// set params
	flag.StringVar(&server.Config.LogPath, "logPath", "log.log", "The log file path.")
	showVersionFlag := flag.Bool("v", false, "Show version info, if true, it will not start server.")
	flag.Parse()

	if *showVersionFlag {
		showVersion()
	} else {
		// all the argument with application will output version info.
		if len(os.Args) == 1 {
			if ClusterSet == nil {
				server.Config.UseKubeFeature = false
				server.Config.IsInSideCluster = false
				server.Config.KubeClientSet = nil
			} else {
				server.Config.UseKubeFeature = true
				server.Config.IsInSideCluster = IsInCluster
				server.Config.KubeClientSet = ClusterSet
			}
			server.Start(conf.ServicePort)
		}
	}
}

// showVersion will print version info
func showVersion() {
	logrus.Infoln("The version is :", Version)
	logrus.Infoln("The build from :", BuildPlatform)
	logrus.Infoln("The build stamp:", BuildStamp)
}

func init() {

	server.Config.LogPath = conf.LogFilePath

	if os.Getenv(EnvUseKubeFeature) == UseKubeFeature {
		logrus.Infoln("Use kube feature mode")
		clientItem, isInCluster, err := client.InitClient()
		if err != nil {
			logrus.Error(err)
			logrus.Infoln("Change mode to disable kube feature mode")
			return
		}
		IsInCluster = isInCluster
		ClusterSet = clientItem
		logrus.Infoln("Init kubernetes cluster success, the mode is:(false : out side of cluster, true: in side of cluster) ", isInCluster)
		server.Config.UseKubeFeature = true
	} else {
		logrus.Infoln("Disable kube feature mode")
	}

	// What the mean of 0666?
	file, err := os.OpenFile(server.Config.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		logrus.SetOutput(fileAndStdoutWriter)
	} else {
		logrus.Infoln("fail to log to file")
	}

	// init the log file

}
