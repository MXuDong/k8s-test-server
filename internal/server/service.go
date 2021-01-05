package server

import "github.com/gin-gonic/gin"

// Define the service info

func Start(port string) {
	r := gin.Default()
	r.GET("/hello", HelloWord)

	// run and listen
	_ = r.Run(port)
}

func HelloWord(c *gin.Context) {
	c.String(200, "hello world")
}
