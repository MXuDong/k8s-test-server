package server

import (
	"github.com/gin-gonic/gin"
)

func KubeBaseInfo(ctx *gin.Context) {
	if Config.UseKubeFeature {
		ctx.String(200, "Kube can use")
	} else {
		ctx.String(200, "Kube can't use now")
	}
}

// todo dynamic model for kubernetes resource
//func GetKubeResource(ctx *gin.Context) {
//
//	version := ctx.Param("version")
//	group := ctx.Param("group")
//
//	kubeConfig := *(Config.KubeConfig)
//	kubeConfig.APIPath = ""
//	kubeConfig.GroupVersion = &schema.GroupVersion{
//		Group:   version,
//		Version: group,
//	}
//}
