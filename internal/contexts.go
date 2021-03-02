package internal

import "github.com/gin-gonic/gin"

func CommonSuccessEmpty(ctx *gin.Context) {
	ctx.JSON(204, nil)
}

func CommonSuccessString(ctx *gin.Context, value string) {
	ctx.JSON(200, struct {
		Value string `json:"value"`
	}{
		Value: value,
	})
}

func CommonSuccessList(ctx *gin.Context, value []interface{}) {
	ctx.JSON(200, struct {
		Items []interface{} `json:"items"`
	}{
		Items: value,
	})
}

func CommonSuccess(ctx *gin.Context, value interface{}) {
	ctx.JSON(200, value)
}
