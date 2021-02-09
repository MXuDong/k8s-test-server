package server

import (
	"github.com/gin-gonic/gin"
	"k8s-test-backend/internal/server/middleware"
	"os"
)

var OK = 200

// Define the service info
type PodEnvInfo struct {
	ServiceIp        string `json:"service_ip"`
	ServiceName      string `json:"service_name"`
	ServiceNamespace string `json:"service_namespace"`
}

// Start will start net work
func Start(port string) {

	e := &PodEnvInfo{
		ServiceIp:        os.Getenv("SERVICE_IP"),
		ServiceName:      os.Getenv("SERVICE_NAME"),
		ServiceNamespace: os.Getenv("SERVICE_NAMESPACE"),
	}

	r := gin.Default()
	gin.SetMode(Config.GinMode)

	// register mid
	r.Use(middleware.LogHandler)

	// envs
	r.GET("/env", GetEnvInfo)
	r.GET("/env-pod", GetPodEnvInfo(e))
	r.GET("/env/:env", GetEnv)

	// logs
	r.GET("/log", GetLogs)

	// switch on kube api trans if kube feature is enable
	kubeGroup := r.Group("/kube-feature")
	kubeGroup.GET("/base-info", KubeBaseInfo)
	if Config.UseKubeFeature {
		// kube feature route here
		kubeGroup.GET("/resource/:resource/namespace/:namespace", GetKubeResource)
	}

	// add base info
	r.GET("/hello", HelloWord)
	r.GET("/routes", Index(r.Routes()))
	r.GET("/", Index(r.Routes()))

	// run and listen
	_ = r.Run(port)
}

// HelloWord is the func of Get
func HelloWord(c *gin.Context) {
	c.String(200, "hello world")
}

// Index will return routes
func Index(info []gin.RouteInfo) func(ctx *gin.Context) {
	type showInfo struct {
		Method  string
		Path    string
		Handler string
	}
	var res []showInfo
	for _, item := range info {
		res = append(res, showInfo{
			Method:  item.Method,
			Path:    item.Path,
			Handler: item.Handler,
		})
	}
	return func(ctx *gin.Context) {
		ctx.JSON(200, res)
	}
}
