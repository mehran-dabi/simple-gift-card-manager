package entity

type GiftCard struct {
	ID        int64  `json:"id"`
	Price     int64  `json:"price"`
	Sender    int64  `json:"sender"`
	Receiver  int64  `json:"receiver"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
