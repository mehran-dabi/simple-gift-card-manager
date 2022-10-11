package service

import (
	"context"
	"dono/domain/giftcard/dto"
	"dono/domain/giftcard/repository"
)

type IGiftCardService interface {
	AddGiftCard(ctx context.Context, price int64) (giftCard *dto.GiftCard, err error)
	SendGiftCard(ctx context.Context, sender, receiver, giftCardID int64) (err error)
}

type GiftCardService struct {
	repository repository.IGiftCardRepository
}

func NewGiftCardService(repository repository.IGiftCardRepository) *GiftCardService {
	return &GiftCardService{repository: repository}
}

func (g *GiftCardService) AddGiftCard(ctx context.Context, price int64) (giftCardDTO *dto.GiftCard, err error) {
	giftCardEntity, err := g.repository.AddGiftCard(ctx, price)
	if err != nil {
		return nil, err
	}

	giftCardDTO = dto.GiftCardDTOFromEntity(giftCardEntity)

	return giftCardDTO, nil
}

func (g *GiftCardService) SendGiftCard(ctx context.Context, sender, receiver, giftCardID int64) (err error) {
	if err := g.repository.SendGiftCard(ctx, sender, receiver, giftCardID); err != nil {
		return err
	}

	return nil
}
