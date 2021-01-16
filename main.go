package main

import (
	"os"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ballpor98/Go-NaCl-S/routers"
)

func readConfig() gin.HandlerFunc {
	AESbase64 := os.Getenv("AES_SECRET")
	if len(AESbase64) == 0 {
		panic("can't read AES_SECRET")
	}

	return func(c *gin.Context) {
		c.Set("AES_SECRET", AESbase64)
	}
}

func setupRouter(e *gin.Engine) {
	e.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"timestamp": time.Now()})
	})

	e.POST("/encrypt", routers.Encrypt)
	e.POST("/decrypt", routers.Decrypt)
	
	return
}

func main() {
	e := gin.Default()
	e.Use(readConfig())
	setupRouter(e)
	e.Run(":8080")
}
