package common

import "github.com/gin-gonic/gin"

func SuccessEmpty(ctx *gin.Context) {
	ctx.JSON(204, nil)
}

func SuccessString(ctx *gin.Context, value string) {
	CodeString(ctx, 200, value)
}

func SuccessList(ctx *gin.Context, value []interface{}) {
	ctx.JSON(200, struct {
		Items []interface{} `json:"items"`
	}{
		Items: value,
	})
}

func Success(ctx *gin.Context, value interface{}) {
	ctx.JSON(200, value)
}

func Error(code int, ctx *gin.Context, err error) {
	ctx.JSON(code, err)
}

func CodeString(ctx *gin.Context, code int, value string) {
	ctx.JSON(code, struct {
		Value string `json:"value"`
	}{
		Value: value,
	})
}
