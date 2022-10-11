package controller

import (
	"dono/domain/giftcard/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IGiftCardController interface {
	AddGiftCard(ctx *gin.Context)
	SendGiftCard(ctx *gin.Context)
}

type GiftCardController struct {
	service service.IGiftCardService
}

func NewGiftCardController(service service.IGiftCardService) *GiftCardController {
	return &GiftCardController{service: service}
}

func (g *GiftCardController) AddGiftCard(ctx *gin.Context) {
	var request AddGiftCard

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	giftCardDTO, err := g.service.AddGiftCard(ctx, request.Price)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, giftCardDTO)
}

func (g *GiftCardController) SendGiftCard(ctx *gin.Context) {
	var request SendGiftCard

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := g.service.SendGiftCard(ctx, request.Sender, request.Receiver, request.GiftCardID); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
