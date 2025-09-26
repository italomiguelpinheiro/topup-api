package controllers

import (
	"net/http"
	"topup-api/internal/usecases"

	"github.com/gin-gonic/gin"
)

type OfferController struct {
	usecase usecases.OfferUsecase
}

func NewOfferController(usecase usecases.OfferUsecase) OfferController {
	return OfferController{
		usecase: usecase,
	}
}

func (oc *OfferController) GetOffers(ctx *gin.Context) {
	offers, err := oc.usecase.GetOffers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, offers)
}
