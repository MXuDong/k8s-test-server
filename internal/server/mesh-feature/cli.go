package mesh_feature

import (
	"github.com/gin-gonic/gin"
	"k8s-test-backend/conf"
)

func RegisterRoute(r *gin.RouterGroup, name, host string) {
	r.GET("/"+name, func(context *gin.Context) {
		// todo trans request to host, copy all request info(request params, request body and request query)
		context.JSON(200, host)
	})
}

func MeshInfo(ctx *gin.Context) {
	if !conf.ApplicationConfig.EnableServerFeature {
		ctx.JSON(200, "the service mesh is disable now")
		return
	}

	ctx.JSON(200, struct {
		Status bool        `json:"mesh_status"`
		List   interface{} `json:"list"`
	}{
		Status: conf.ApplicationConfig.EnableServerFeature,
		List:   conf.ApplicationConfig.ServiceMeshMapper,
	})
}
