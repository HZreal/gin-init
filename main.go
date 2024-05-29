package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var r *gin.Engine = gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.Run("0.0.0.0:8000")
}
