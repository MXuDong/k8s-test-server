package server

import (
	"github.com/gin-gonic/gin"
	"os"
)

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
	r.GET("/podenv", GetPodEnvInfo(e))
	r.GET("/env", GetEnvInfo)

	// run and listen
	_ = r.Run(port)
}

// HelloWord is the func of Get
func HelloWord(c *gin.Context) {
	c.String(200, "hello world")
}

// GetEnvInfo will output some env info
func GetPodEnvInfo(e *PodEnvInfo) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, e)
	}
}

type StringArrayResult struct {
	Paths []string `json:"paths"`
}

// GetEnvInfo will return all the environment values
func GetEnvInfo(c *gin.Context) {
	envs := os.Environ()
	res := &StringArrayResult{
		Paths: envs,
	}
	c.JSON(200, res)
}
