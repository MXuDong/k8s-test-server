package k8s_feature

import (
	"github.com/gin-gonic/gin"
	"k8s-test-backend/conf"
	client2 "k8s-test-backend/pkg/client"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListNamespace(ctx *gin.Context) {
	client := conf.ApplicationConfig.KubeClientSet
	res, err := client.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	result := make([]string, 0)
	for _, item := range res.Items {
		result = append(result, item.Name)
	}
	client2.BaseResponse(200, ctx, result)
}

func GetNamespace(ctx *gin.Context) {
	name := ctx.Param("name")
	client := conf.ApplicationConfig.KubeClientSet
	res, err := client.CoreV1().Namespaces().Get(ctx, name, v1.GetOptions{})
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	client2.BaseResponse(200, ctx, res)
}
