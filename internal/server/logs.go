package server

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"k8s-test-backend/conf"
)

func GetLogs(ctx *gin.Context) {
	readResult, err := ioutil.ReadFile(conf.ApplicationConfig.LogPath)
	if err != nil {
		ctx.JSON(500, err)
	} else {
		ctx.String(200, string(readResult))
	}
}
