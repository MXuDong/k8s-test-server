package internal

import (
	"github.com/gin-gonic/gin"
	"time"
)

// SleepHandle will return a handler, the handler will auto sleep by target duration.
// Param sleepTime for get sleep duration.
func SleepHandle(sleepTime func() time.Duration) func(*gin.Context) {
	return func(ctx *gin.Context) {
		time.Sleep(sleepTime())
		ctx.JSON(204, nil)
	}
}

func ErrorHandler(errorGenerator func() (code int, err error, msg string)) func(*gin.Context) {
	code, err, msg := errorGenerator()
	return func(ctx *gin.Context) {
		ctx.JSON(code, struct {
			Error error  `json:"error"`
			Msg   string `json:"msg"`
		}{
			Error: err,
			Msg:   msg,
		})
	}
}
