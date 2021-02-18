package main

import (
	"github.com/sirupsen/logrus"
)

var Version = ""
var BuildPlatform = ""
var BuildStamp = ""

func main() {
	cmd := InitCmd()
	if cmd == nil {
		return
	}

	if err := cmd.Execute(); err != nil {
		logrus.Error(err)
	}
}
