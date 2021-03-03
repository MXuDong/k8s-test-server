package server

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"k8s-test-backend/conf"
	"k8s-test-backend/pkg/common"
)

func GetLogs(ctx *gin.Context) {
	readResult, err := ioutil.ReadFile(conf.ApplicationConfig.LogPath)
	if err != nil {
		common.Error(500, ctx, err)
	} else {
		common.SuccessString(ctx, string(readResult))
	}
}
