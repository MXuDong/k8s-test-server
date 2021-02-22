package server

import (
	"github.com/gin-gonic/gin"
	"k8s-test-backend/conf"
)

func KubeBaseInfo(ctx *gin.Context) {
	if conf.ApplicationConfig.UseKubernetesFeature {
		ctx.String(200, "Kube can use")
	} else {
		ctx.String(200, "Kube can't use now")
	}
}

// todo dynamic model for kubernetes resource
//func GetKubeResource(ctx *gin.Context) {
//
//	apiPath := ctx.Param("apiPath")
//	version := ctx.Param("version")
//	group := ctx.Param("group")
//	resourceName := ctx.Param("name")
//
//	kubeConfig := conf.ApplicationConfig.KubeClientConf
//	kubeConfig.APIPath = apiPath
//	kubeConfig.GroupVersion = &schema.GroupVersion{
//		Group:   group,
//		Version: version,
//	}
//	client, err := rest.RESTClientFor(kubeConfig)
//
//	if err != nil {
//		ctx.JSON(400, err)
//		return
//	}
//
//	var result runtime.Object
//
//	err = client.Get().Name(resourceName).Do(ctx).Into(result)
//	if err != nil {
//		ctx.JSON(400, err)
//		return
//	}
//	ctx.JSON(200, result)
//
//}
