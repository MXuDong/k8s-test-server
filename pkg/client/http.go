package client

import (
	"github.com/gin-gonic/gin"
	"k8s-test-backend/conf"
)

type CommonRequestResponse struct {
	Method          string      `json:"method"`
	ApplicationName string      `json:"application_name"`
	Object          interface{} `json:"object"`
}

type HeaderResponse struct {
	Header interface{} `json:"header"`
	Body   interface{} `json:"body"`
}

func BaseResponse(code int, ctx *gin.Context, obj interface{}) {
	ctx.Header("APP_NAME", conf.ApplicationConfig.ApplicationRunName)
	ctx.Header("APP_VERSION", conf.ApplicationConfig.Version)
	ctx.JSON(code, CommonRequestResponse{
		Method:          ctx.Request.Method,
		Object:          obj,
		ApplicationName: conf.ApplicationConfig.ApplicationRunName,
	})
}

func GetResponse(ctx *gin.Context, obj interface{}) {
	BaseResponse(200, ctx, obj)
}

func PostResponse(ctx *gin.Context, obj interface{}) {
	BaseResponse(200, ctx, obj)
}

func PutResponse(ctx *gin.Context, obj interface{}) {
	BaseResponse(200, ctx, obj)
}

func PatchResponse(ctx *gin.Context, obj interface{}) {
	BaseResponse(200, ctx, obj)
}

func DeleteResponse(ctx *gin.Context) {
	ctx.Header("APP_NAME", conf.ApplicationConfig.ApplicationRunName)
	ctx.Header("APP_VERSION", conf.ApplicationConfig.Version)
	ctx.JSON(204, nil)
}
