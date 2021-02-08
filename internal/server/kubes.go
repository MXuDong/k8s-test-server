package server

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/runtime"
)

func KubeBaseInfo(ctx *gin.Context) {
	if Config.UseKubeFeature {
		ctx.String(200, "Kube can use")
	} else {
		ctx.String(200, "Kube can't use now")
	}
}

func GetKubeResource(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	resource := ctx.Param("resource")
	var res runtime.Object
	err := Config.KubeClientSet.RESTClient().Get().
		Namespace(namespace).
		Resource(resource).
		Do(ctx).
		Into(res)
	if err != nil {
		ctx.JSON(200, err)
		return
	}

	ctx.JSON(200, res)
}
