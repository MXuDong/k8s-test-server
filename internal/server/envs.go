package server

import (
	"github.com/gin-gonic/gin"
	"k8s-test-backend/conf"
	"k8s-test-backend/pkg/common"
	"os"
)

// Define the service info
type podEnvInfo struct {
	ServiceIp        string `json:"service_ip"`
	ServiceName      string `json:"service_name"`
	ServiceNamespace string `json:"service_namespace"`
}

var podEnvInfoInstance *podEnvInfo = nil

// GetEnvInfo will output some env info
func PodEnv(c *gin.Context) {
	if podEnvInfoInstance == nil {
		podEnvInfoInstance = &podEnvInfo{
			ServiceIp:        conf.ApplicationConfig.ServiceIp,
			ServiceName:      conf.ApplicationConfig.ServiceName,
			ServiceNamespace: conf.ApplicationConfig.ServiceNamespace,
		}
	}
	common.Success(c, podEnvInfoInstance)
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
	common.Success(c, res)
}

// GetEnv will return target env value, if not exits, it will return empty value
func GetEnv(c *gin.Context) {
	common.SuccessString(c, os.Getenv(c.Param("env")))
}
