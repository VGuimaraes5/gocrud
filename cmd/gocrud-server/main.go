package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "OK",
				})
			})
		}
	}

	router.Run(":8080")
}
