package main

import (
	"database/sql"
	"dono/config"
	"dono/domain/giftcard/controller"
	"dono/domain/giftcard/repository"
	"dono/domain/giftcard/service"
	"dono/infrastructure"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs := config.Init()

	db, err := infrastructure.NewRepository(
		configs.Database.User,
		configs.Database.Pass,
		configs.Database.Host,
		configs.Database.Port,
		configs.Database.Name,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	giftCardRepo := repository.NewGiftCardRepository(db)
	giftCardService := service.NewGiftCardService(giftCardRepo)
	giftCardController := controller.NewGiftCardController(giftCardService)

	r := gin.Default()
	giftCard := r.Group("/gift-card")
	{
		giftCard.POST("/add", giftCardController.AddGiftCard)
		giftCard.POST("/send", giftCardController.SendGiftCard)
		giftCard.POST("/received", giftCardController.GetReceivedGiftCards)
		giftCard.POST("/update-status", giftCardController.UpdateGiftCardStatus)
	}

	appPort := configs.Server.Port
	if appPort == "" {
		appPort = "8888"
	}
	log.Fatal(r.Run(":" + appPort))
}
