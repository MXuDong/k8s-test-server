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

	r.GET("/hello", HelloWord)
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
	// run and listen
	_ = r.Run(port)
}

// HelloWord is the func of Get
func HelloWord(c *gin.Context) {
	c.String(200, "hello world")
}
