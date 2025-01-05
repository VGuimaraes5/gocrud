package main

import (
	"net/http"

	"github.com/VGuimaraes5/gocrud/internal/model"
	"github.com/VGuimaraes5/gocrud/pkg/gorm/db"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.Init()
	if err != nil {
		panic("Error initializing database")
	}

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

			v1.GET("/users", func(ctx *gin.Context) {
				db := db.GetConnection()
				users := []model.User{}

				if err := db.Find(&users).Error; err != nil {
					ctx.Header("Content-type", "application/json")
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"message": "error listing users",
					})
					return
				}

				ctx.Header("Content-type", "application/json")
				ctx.JSON(http.StatusOK, gin.H{
					"users": users,
				})
			})
		}
	}

	router.Run(":8080")
}
