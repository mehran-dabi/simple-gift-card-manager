package controller

type AddGiftCard struct {
	Price int64 `json:"price"`
}

type SendGiftCard struct {
	Sender     int64 `json:"sender"`
	Receiver   int64 `json:"receiver"`
	GiftCardID int64 `json:"gift_card_id"`
}
