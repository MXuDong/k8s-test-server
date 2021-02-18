package server

import (
	"github.com/gin-gonic/gin"
	"k8s-test-backend/conf"
	"k8s-test-backend/internal/server/middleware"
)

var OK = 200

// Start will start net work
func Start() {
	gin.SetMode(conf.ApplicationConfig.Mode)
	r := gin.Default()

	// register mid
	r.Use(middleware.LogHandler)

	// envs
	r.GET("/env", GetEnvInfo)
	r.GET("/env/:env", GetEnv)

	// common handler about stander restFul request method
	commonGroup := r.Group("/common")
	{
		// only test for once
		if conf.ApplicationConfig.UseCommonHttp {
			commonGroup.GET("/resources", CommonGet)
			commonGroup.POST("/resources", CommonPost)
			commonGroup.PUT("/resources", CommonPut)
			commonGroup.PATCH("/resources", CommonPatch)
			commonGroup.DELETE("/resources", CommonDelete)
		}
		if conf.ApplicationConfig.UseCacheHttp {
			// test for cache
			commonGroup.GET("/resources-cache/key/:key/value/:value", CacheGet)
			commonGroup.GET("/resources-cache-list", CacheList)
			commonGroup.POST("/resources-cache/", CachePost)
			commonGroup.PUT("/resources-cache/key/:key/value/:value", CachePut)
			commonGroup.DELETE("/resources-cache/key/:key/value/:value", CacheDelete)
			commonGroup.DELETE("/resources-cache/", CacheClean)
			commonGroup.PATCH("/resources-cache/key/:key/value/:value", CachePatch)
		}

	}

	// logs
	r.GET("/log", GetLogs)

	// switch on kube api trans if kube feature is enable
	kubeGroup := r.Group("/kube-feature")
	kubeGroup.GET("/base-info", KubeBaseInfo)
	if conf.ApplicationConfig.UseKubernetesFeature {
		// kube feature route here
		// pods
		kubeGroup.GET("/pod-env", PodEnv)

		//================= namespace
		namespaceGroup := kubeGroup.Group("/namespace")
		namespaceGroup.GET("/", ListNamespace)
	}

	// add base info
	r.GET("/hello", HelloWord)
	r.GET("/routes", Index(r.Routes()))
	r.GET("/", Index(r.Routes()))

	// run and listen
	_ = r.Run(conf.ApplicationConfig.Port)
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
