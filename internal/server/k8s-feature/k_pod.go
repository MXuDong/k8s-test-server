package k8s_feature

import (
	"github.com/gin-gonic/gin"
	"k8s-test-backend/conf"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListPodByNamespace(ctx *gin.Context) {
	namespace := ctx.Param("namespace")

	client := conf.ApplicationConfig.KubeClientSet
	res, err := client.CoreV1().Pods(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	result := make([]string, 0)
	for _, item := range res.Items {
		result = append(result, item.Name)
	}

	ctx.JSON(200, result)
}

func GetPodByName(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")

	client := conf.ApplicationConfig.KubeClientSet

	res, err := client.CoreV1().Pods(namespace).Get(ctx, name, v1.GetOptions{})
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, res)
}
