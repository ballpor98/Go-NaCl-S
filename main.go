package main

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ballpor98/Go-NaCl-S/routers"
)


func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"timestamp": time.Now()})
	})

	r.POST("/encrypt", routers.Encrypt)
	r.POST("/decrypt", routers.Decrypt)
	
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
