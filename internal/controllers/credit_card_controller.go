package controllers

import (
	"net/http"
	"strconv"

	models "topup-api/internal/domains"
	"topup-api/internal/usecases"

	"github.com/gin-gonic/gin"
)

type CreditCardController struct {
	creditCardUsecase usecases.CreditCardUsecase
}

func NewCreditCardController(usecase usecases.CreditCardUsecase) CreditCardController {
	return CreditCardController{
		creditCardUsecase: usecase,
	}
}

func (c *CreditCardController) GetCreditCardsByClientId(ctx *gin.Context) {
	id := ctx.Param("clientId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: "Id do cliente não pode ser nulo"})
		return
	}

	clientId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: "Id do cliente precisa ser um número"})
		return
	}

	cards, err := c.creditCardUsecase.GetCreditCardsByClientId(clientId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: "Erro ao buscar cartões"})
		return
	}

	if len(cards) == 0 {
		ctx.JSON(http.StatusNotFound, models.Response{Message: "Nenhum cartão encontrado para este cliente"})
		return
	}

	ctx.JSON(http.StatusOK, cards)
}
