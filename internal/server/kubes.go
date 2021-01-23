package server

import "github.com/gin-gonic/gin"

func KubeBaseInfo(ctx *gin.Context){
	if Config.UseKubeFeature{
		ctx.String(200, "Kube can use")
	}else {
		ctx.String(200, "Kube can't use now")
	}
}