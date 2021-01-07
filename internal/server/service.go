package server

import (
	"github.com/gin-gonic/gin"
	"os"
)

// Define the service info
type PodEnvInfo struct {
	ServiceIp   string `json:"service_ip"`
	ServiceName string `json:"service_name"`
	ServiceNamespace string `json:"service_namespace"`
}

// Start will start net work
func Start(port string) {

	e := &PodEnvInfo{
		ServiceIp:   os.Getenv("SERVICE_IP"),
		ServiceName: os.Getenv("SERVICE_NAME"),
		ServiceNamespace: os.Getenv("SERVICE_NAMESPACE"),
	}

	r := gin.Default()
	r.GET("/hello", HelloWord)
	r.GET("/podenv", GetEnvInfo(e))

	// run and listen
	_ = r.Run(port)
}

// HelloWord is the func of Get
func HelloWord(c *gin.Context) {
	c.String(200, "hello world")
}

// GetEnvInfo will output some env info
func GetEnvInfo(e *PodEnvInfo) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, e)
	}
}
