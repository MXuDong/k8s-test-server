package mesh_feature

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"k8s-test-backend/conf"
	"net/http"
)

func RegisterRoute(r *gin.RouterGroup, name, host string) {
	r.GET("/"+name, func(context *gin.Context) {
		// todo trans request copy request info(request params, request body and request query)
		res, err := http.Get(host)
		if err != nil {
			context.JSON(400, err)
			return
		}
		robots, err := ioutil.ReadAll(res.Body)
		if err != nil {
			context.JSON(400, err)
			return
		}
		_ = res.Body.Close()

		_, _ = context.Writer.Write(robots)
		context.Writer.WriteHeader(200)
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
