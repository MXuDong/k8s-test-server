package server

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func GetLogs(ctx *gin.Context) {
	readResult, err := ioutil.ReadFile(Config.LogPath)
	if err != nil {
		ctx.JSON(500, err)
	} else {
		ctx.String(200, string(readResult))
	}
}
