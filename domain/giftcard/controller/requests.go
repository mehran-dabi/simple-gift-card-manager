package controller

type AddGiftCard struct {
	Price int64 `json:"price"`
}

type SendGiftCard struct {
	Sender     int64 `json:"sender"`
	Receiver   int64 `json:"receiver"`
	GiftCardID int64 `json:"gift_card_id"`
}

type GiftCardStatus string

var (
	GiftCardStatusPending  GiftCardStatus = "PENDING"
	GiftCardStatusAccepted GiftCardStatus = "ACCEPTED"
	GiftCardStatusRejected GiftCardStatus = "REJECTED"
)

func (g GiftCardStatus) Verify() bool {
	switch g {
	case GiftCardStatusPending, GiftCardStatusRejected, GiftCardStatusAccepted:
		return true
	default:
		return false
	}
}

func (g GiftCardStatus) String() string {
	return g.String()
}

type UpdateGiftCardStatus struct {
	ID     int64          `json:"id"`
	Status GiftCardStatus `json:"status"`
}

type GetReceivedGiftCards struct {
	ReceiverID int64 `json:"receiver_id"`
}
