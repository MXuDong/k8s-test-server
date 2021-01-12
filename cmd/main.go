package main

import (
	"fmt"
	"k8s-test-backend/internal/server"
	"os"
)

var Version = ""
var BuildPlatform = ""
var BuildStamp = ""

// main will start application
func main() {

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
