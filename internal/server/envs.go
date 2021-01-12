package server

import (
	"github.com/gin-gonic/gin"
	"os"
)

// GetEnvInfo will output some env info
func GetPodEnvInfo(e *PodEnvInfo) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(OK, e)
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
	c.JSON(OK, res)
}

// GetEnv will return target env value, if not exits, it will return empty value
func GetEnv(c *gin.Context) {
	c.String(OK, os.Getenv(c.Param("env")))
}
