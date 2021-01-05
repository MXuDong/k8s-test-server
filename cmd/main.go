package main

import (
	"fmt"
	"os"
)

var Version = ""
var BuildPlatform = ""
var BuildStamp = ""

func main() {
	if len(os.Args) > 1 {
		fmt.Println(Version)
	}
	fmt.Println("Complete")
}

func showVersion() {
	fmt.Println("The version is :", Version)
	fmt.Println("The build from :", BuildPlatform)
	fmt.Println("The build stamp:", BuildStamp)
}
