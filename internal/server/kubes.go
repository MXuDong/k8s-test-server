package server

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func KubeBaseInfo(ctx *gin.Context) {
	if Config.UseKubeFeature {
		ctx.String(200, "Kube can use")
	} else {
		ctx.String(200, "Kube can't use now")
	}
}

func GetKubeResource(ctx *gin.Context) {

	kubeConfig := *(Config.KubeConfig)
	kubeConfig.APIPath = ""
	kubeConfig.GroupVersion = &schema.GroupVersion{
		Group:   "",
		Version: "",
	}
}
