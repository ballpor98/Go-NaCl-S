package main

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)


func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"timestamp": time.Now()})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
