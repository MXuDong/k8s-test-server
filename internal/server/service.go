package server

import (
	"github.com/gin-gonic/gin"
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
	r.GET("/hello", HelloWord)
	// envs
	r.GET("/podenv", GetPodEnvInfo(e))
	r.GET("/env", GetEnvInfo)
	r.GET("/env/:env", GetEnv)

	// run and listen
	_ = r.Run(port)
}

// HelloWord is the func of Get
func HelloWord(c *gin.Context) {
	c.String(200, "hello world")
}
