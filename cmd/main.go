package main

import (
	controller "topup-api/internal/controllers"
	"topup-api/internal/db"
	"topup-api/internal/repositories"
	"topup-api/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	dbConnection, err := db.NewPostgresDB()
	if err != nil {
		panic(err)
	}

	ClientRepository := repositories.NewClientRepository(dbConnection)
	ClientUsecase := usecases.NewClientUsercase(ClientRepository)
	ClientController := controller.NewClientController(ClientUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	server.GET("/clients", ClientController.GetClients)

	server.Run(":8000")
}
