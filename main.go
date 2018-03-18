package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		msg := GetHello("nodemon 2!!!!")
		fmt.Println(msg)
		c.String(http.StatusOK, msg)
	})

	// health check
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":3000")
}

// GetHello returns Hello message to the target
func GetHello(target string) string {
	return fmt.Sprintf("Hello %v", target)
}
