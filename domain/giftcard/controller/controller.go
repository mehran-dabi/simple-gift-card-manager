package controller

import (
	"dono/domain/giftcard/service"
	"dono/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IGiftCardController interface {
	AddGiftCard(ctx *gin.Context)
	SendGiftCard(ctx *gin.Context)
	GetReceivedGiftCards(ctx *gin.Context)
	UpdateGiftCardStatus(ctx *gin.Context)
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

	helper.GinResponse(ctx, http.StatusOK, giftCardDTO)
}

func (g *GiftCardController) SendGiftCard(ctx *gin.Context) {
	var request SendGiftCard

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := g.service.SendGiftCard(ctx, request.Price, request.Sender, request.Receiver); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	helper.GinResponse(ctx, http.StatusOK, nil)
}

func (g *GiftCardController) GetReceivedGiftCards(ctx *gin.Context) {
	var request GetReceivedGiftCards

	filter := ctx.Query("filter")
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	receivedGiftCards, err := g.service.GetReceivedGiftCards(ctx, request.ReceiverID, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	helper.GinResponse(ctx, http.StatusOK, receivedGiftCards)
}

func (g *GiftCardController) UpdateGiftCardStatus(ctx *gin.Context) {
	var request UpdateGiftCardStatus

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if !request.Status.Verify() {
		helper.GinErrResponse(ctx, http.StatusBadRequest, fmt.Errorf("status is not valid"))
		return
	}

	if err := g.service.UpdateGiftCardStatus(ctx, request.ID, request.Status.String()); err != nil {
		helper.GinErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	helper.GinResponse(ctx, http.StatusOK, nil)
}
