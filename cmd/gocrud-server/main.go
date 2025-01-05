package main

import (
	"github.com/VGuimaraes5/gocrud/internal/server"
	"github.com/VGuimaraes5/gocrud/pkg/db/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := postgres.GetConnection()
	if err != nil {
		panic("Error initializing database")
	}

	router := gin.Default()

	server := server.NewServer(router, db)
	server.Run(":8080")
}
