package controllers

import (
	"net/http"
	"topup-api/internal/usecases"

	"github.com/gin-gonic/gin"
)

type clientController struct {
	clientUsecase usecases.ClientUsecase
}

func NewClientController(usecase usecases.ClientUsecase) clientController {
	return clientController{
		clientUsecase: usecase,
	}
}

func (c *clientController) GetClients(ctx *gin.Context) {
	clients, err := c.clientUsecase.GetClients()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, clients)
}
