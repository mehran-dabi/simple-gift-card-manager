package repository

const GiftCardTableName = "gift_cards"
const (
	InsertGiftCard       = `INSERT INTO ` + GiftCardTableName + ` SET price = ?, created_at = NOW(), updated_at = NOW()`
	UpdateGiftCard       = `UPDATE ` + GiftCardTableName + ` SET sender = ?, receiver = ?, updated_at = NOW() WHERE id = ?`
	GetReceivedGiftCards = `SELECT id, price, sender, receiver, created_at, updated_at FROM ` + GiftCardTableName + ` WHERE receiver = ?`
)
