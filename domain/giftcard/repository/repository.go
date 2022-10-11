package repository

import (
	"context"
	"database/sql"
	"dono/domain"
	"dono/domain/giftcard/entity"
	"fmt"
)

type IGiftCardRepository interface {
	AddGiftCard(ctx context.Context, price int64) (giftCard *entity.GiftCard, err error)
	SendGiftCard(ctx context.Context, senderID, receiverID, giftCardID int64) (err error)
	GetReceivedGiftCards(ctx context.Context, receiverID int64) (giftCards []*entity.GiftCard, err error)
	UpdateGiftCardStatus(ctx context.Context, ID int64, status string) (err error)
}

type GiftCardRepository struct {
	db *sql.DB
}

func NewGiftCardRepository(db *sql.DB) *GiftCardRepository {
	return &GiftCardRepository{db: db}
}

func (g *GiftCardRepository) AddGiftCard(ctx context.Context, price int64) (giftCard *entity.GiftCard, err error) {
	result, err := g.db.ExecContext(ctx, InsertGiftCard, price)
	if err != nil {
		return nil, fmt.Errorf("failed to insert gift card: %w", err)
	}
	giftCard = &entity.GiftCard{Price: price}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get inserted id: %w", err)
	}
	giftCard.ID = id

	return giftCard, nil
}

func (g *GiftCardRepository) SendGiftCard(ctx context.Context, senderID, receiverID, giftCardID int64) (err error) {
	_, err = g.db.ExecContext(ctx, UpdateGiftCard, senderID, receiverID, giftCardID)
	if err != nil {
		return fmt.Errorf("failed to update gift card: %w", err)
	}
	return nil
}

func (g *GiftCardRepository) GetReceivedGiftCards(ctx context.Context, receiverID int64) (giftCards []*entity.GiftCard, err error) {
	results, err := g.db.QueryContext(ctx, GetReceivedGiftCards, receiverID)
	if err != nil {
		return nil, fmt.Errorf("failed to get received gift cards: %w", err)
	}

	for results.Next() {
		giftCard := new(entity.GiftCard)
		err = results.Scan(
			&giftCard.ID,
			&giftCard.Price,
			&giftCard.Sender,
			&giftCard.Receiver,
			&giftCard.Status,
			&giftCard.CreatedAt,
			&giftCard.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to read record from results: %w", err)
		}
		giftCards = append(giftCards, giftCard)
	}

	if len(giftCards) == 0 {
		return nil, domain.ErrNotFound
	}

	return giftCards, nil
}

func (g *GiftCardRepository) UpdateGiftCardStatus(ctx context.Context, ID int64, status string) (err error) {
	_, err = g.db.ExecContext(ctx, UpdateGiftCardStatus, status, ID)
	if err != nil {
		return fmt.Errorf("failed to update gift card: %w", err)
	}
	return nil
}
