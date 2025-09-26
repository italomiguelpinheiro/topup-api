package controllers

import (
	"net/http"
	"strconv"
	models "topup-api/internal/domains"
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

func (c *clientController) GetClientById(ctx *gin.Context) {

	id := ctx.Param("clientId")

	if id == "" {
		response := models.Response{
			Message: "Id do client nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	clientId, err := strconv.Atoi(id)

	if err != nil {
		response := models.Response{
			Message: "Id do client precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	client, err := c.clientUsecase.GetClientById(clientId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if client == nil {
		response := models.Response{
			Message: "Cliente nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, client)
}
