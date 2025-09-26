package main

import (
	"topup-api/internal/controllers"
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

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	ClientRepository := repositories.NewClientRepository(dbConnection)
	ClientUsecase := usecases.NewClientUsercase(ClientRepository)
	ClientController := controller.NewClientController(ClientUsecase)

	PurchaseRepository := repositories.NewPurchaseRepository(dbConnection)
	PurchaseUsecase := usecases.NewPurchaseUsecase(PurchaseRepository)
	PurchaseController := controller.NewPurchaseController(PurchaseUsecase)

	OfferRepository := repositories.NewOfferRepository(dbConnection)
	OfferUsecase := usecases.NewOfferUsecase(OfferRepository)
	OfferController := controllers.NewOfferController(OfferUsecase)

	server.GET("/clients", ClientController.GetClients)
	server.GET("/offers", OfferController.GetOffers)
	server.GET("/client/:clientId", ClientController.GetClientById)

	server.POST("/purchase", PurchaseController.CreateProduct)

	server.Run(":8000")
}
