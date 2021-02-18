package main

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var Version = ""
var BuildPlatform = ""
var BuildStamp = ""

var IsInCluster = false
var ClusterSet *kubernetes.Clientset = nil
var KubeConfig *rest.Config

const (
	EnvUseKubeFeature = "USE_KUBE_FEATURE"
	UseKubeFeature    = "true"
)

func main() {
	cmd := InitCmd()
	if cmd == nil {
		return
	}

	if err := cmd.Execute(); err != nil {
		logrus.Error(err)
	}
}

//// main will start application
//func main() {
//
//	// set params
//	flag.StringVar(&server.Config.LogPath, "logPath", conf.LogFilePath, "The log file path.")
//	flag.StringVar(&server.Config.GinMode, "ginMode", gin.DebugMode, "The mode of gin.")
//	flag.StringVar(&server.Config.ApplicationPort, "port", conf.ServicePort, "The port of application.")
//	showVersionFlag := flag.Bool("v", false, "Show version info, if true, it will not start server.")
//	if home := homedir.HomeDir(); home != "" {
//		server.Config.KubeConfigPath = *flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
//	} else {
//		server.Config.KubeConfigPath = *flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
//	}
//	flag.Parse()
//
//	initFunc()
//
//	if *showVersionFlag {
//		showVersion()
//	} else {
//		// all the argument with application will output version info.
//		if len(os.Args) == 1 {
//			if ClusterSet == nil {
//				server.Config.UseKubeFeature = false
//				server.Config.IsInSideCluster = false
//				server.Config.KubeClientSet = nil
//				server.Config.KubeConfig = nil
//			} else {
//				server.Config.UseKubeFeature = true
//				server.Config.IsInSideCluster = IsInCluster
//				server.Config.KubeClientSet = ClusterSet
//				server.Config.KubeConfig = KubeConfig
//			}
//			server.Start()
//		}
//	}
//}
//
// showVersion will print version info

func init() {

}

//
//func initFunc() {
//
//	// init kube if can use kube feature
//	if os.Getenv(EnvUseKubeFeature) == UseKubeFeature {
//		logrus.Infoln("Use kube feature mode")
//		clientItem, config, isInCluster, err := client2.InitClient(server.Config.KubeConfigPath)
//		if err != nil {
//			logrus.Error(err)
//			logrus.Infoln("Change mode to disable kube feature mode")
//		} else {
//			IsInCluster = isInCluster
//			ClusterSet = clientItem
//			KubeConfig = config
//			logrus.Infoln("Init kubernetes cluster success, the mode is:(false : out side of cluster, true: in side of cluster) ", isInCluster)
//			server.Config.UseKubeFeature = true
//		}
//	} else {
//		logrus.Infoln("Disable kube feature mode")
//	}
//

//}
