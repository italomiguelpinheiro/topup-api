package controllers

import (
	"net/http"
	models "topup-api/internal/domains"
	"topup-api/internal/usecases"

	"github.com/gin-gonic/gin"
)

type PurchaseController struct {
	usecase usecases.PurchaseUsecase
}

func NewPurchaseController(usecase usecases.PurchaseUsecase) PurchaseController {
	return PurchaseController{
		usecase: usecase,
	}
}

func (p *PurchaseController) CreateProduct(ctx *gin.Context) {
	var order models.Order
	err := ctx.BindJSON(&order)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedOrder, err := p.usecase.CreateOrder(order)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedOrder)

}
