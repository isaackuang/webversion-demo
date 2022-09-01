package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"
)

func main() {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.GET("/healthz", func(c *gin.Context) {
		hostname, _ := os.Hostname()
		c.JSON(200, gin.H{"version": "v1.0", "hostname": hostname})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
