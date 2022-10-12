package repository

const GiftCardTableName = "gift_cards"
const (
	InsertGiftCard                       = `INSERT INTO ` + GiftCardTableName + ` SET price = ?, created_at = NOW(), updated_at = NOW()`
	UpdateGiftCard                       = `UPDATE ` + GiftCardTableName + ` SET sender = ?, receiver = ?, updated_at = NOW() WHERE id = ?`
	UpdateGiftCardStatus                 = `UPDATE ` + GiftCardTableName + ` SET status = ?, updated_at = NOW() WHERE id = ?`
	GetReceivedGiftCardsWithStatusFilter = `SELECT id, price, sender, receiver, status, created_at, updated_at FROM ` + GiftCardTableName + ` WHERE receiver = ? and status = ?`
	GetReceivedGiftCards                 = `SELECT id, price, sender, receiver, status, created_at, updated_at FROM ` + GiftCardTableName + ` WHERE receiver = ?`
)
