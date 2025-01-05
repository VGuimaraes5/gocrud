package server

import (
	"net/http"

	"github.com/VGuimaraes5/gocrud/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewServer(router *gin.Engine, db *gorm.DB) *Server {
	return &Server{
		router: router,
		db:     db,
	}
}

func (s *Server) Run(addr ...string) {
	api := s.router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "OK",
				})
			})

			v1.GET("/users", func(ctx *gin.Context) {
				users := []model.User{}

				if err := s.db.Find(&users).Error; err != nil {
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

			v1.POST("/user", func(ctx *gin.Context) {
				user := model.User{}

				if err := ctx.ShouldBindJSON(&user); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"message": "invalid request",
					})
					return
				}

				if err := s.db.Create(&user).Error; err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"message": "error creating user",
					})
					return
				}

				ctx.JSON(http.StatusCreated, gin.H{
					"user": user,
				})
			})
		}
	}

	s.router.Run(addr...)
}
