package main

import (
	"dono/domain/giftcard/controller"
	"dono/domain/giftcard/service"
	"dono/infrastructure"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	repositories, err := infrastructure.NewRepository(dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer func(repositories *infrastructure.Repositories) {
		_ = repositories.Close()
	}(repositories)

	giftCardService := service.NewGiftCardService(repositories.GiftCard)
	giftCardController := controller.NewGiftCardController(giftCardService)

	r := gin.Default()
	giftCard := r.Group("/gift-card")
	{
		giftCard.POST("/add", giftCardController.AddGiftCard)
		giftCard.POST("/send", giftCardController.SendGiftCard)
		giftCard.POST("/received", giftCardController.GetReceivedGiftCards)
		giftCard.POST("/update-status", giftCardController.UpdateGiftCardStatus)
	}

	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8888"
	}
	log.Fatal(r.Run(":" + appPort))
}
