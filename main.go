package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "AI Infra Go CareerGO starts!",
			"status":  "ready",
		})
	})

	r.Run(":8080")
}
