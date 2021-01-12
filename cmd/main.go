package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"k8s-test-backend/internal/server"
	client "k8s-test-backend/package"
	"os"
)

var Version = ""
var BuildPlatform = ""
var BuildStamp = ""

// main will start application
func main() {
	_, isInCluster, err := client.InitClient()
	if err != nil {
		logrus.Error(err)
	}

	logrus.Infoln("Init kubernetes cluster success, the mode is:", isInCluster)

	// all the argument with application will output version info.
	if len(os.Args) > 1 {
		fmt.Println(Version)
	} else {
		server.Start(":3000")
	}
}

// showVersion will print version info
func showVersion() {
	fmt.Println("The version is :", Version)
	fmt.Println("The build from :", BuildPlatform)
	fmt.Println("The build stamp:", BuildStamp)
}
