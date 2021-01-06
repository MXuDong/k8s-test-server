package server

import "github.com/gin-gonic/gin"

// Define the service info

// Start will start net work
func Start(port string) {
	r := gin.Default()
	r.GET("/hello", HelloWord)

	// run and listen
	_ = r.Run(port)
}

// HelloWord is the func of Get
func HelloWord(c *gin.Context) {
	c.String(200, "hello world")
}
