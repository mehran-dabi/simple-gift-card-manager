package service

import (
	"context"
	"dono/domain/giftcard/dto"
	"dono/domain/giftcard/repository"
)

type IGiftCardService interface {
	AddGiftCard(ctx context.Context, price int64) (giftCard *dto.GiftCard, err error)
	SendGiftCard(ctx context.Context, price, sender, receiver int64) (err error)
	UpdateGiftCardStatus(ctx context.Context, ID int64, status string) (err error)
	GetReceivedGiftCards(ctx context.Context, receiverID int64, statusFilter string) (giftCards []*dto.GiftCard, err error)
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

func (g *GiftCardService) SendGiftCard(ctx context.Context, price, sender, receiver int64) (err error) {
	giftCardEntity, err := g.repository.AddGiftCard(ctx, price)
	if err != nil {
		return err
	}

	if err := g.repository.SendGiftCard(ctx, sender, receiver, giftCardEntity.ID); err != nil {
		return err
	}

	return nil
}

func (g *GiftCardService) UpdateGiftCardStatus(ctx context.Context, ID int64, status string) (err error) {
	if err := g.repository.UpdateGiftCardStatus(ctx, ID, status); err != nil {
		return err
	}

	return nil
}

func (g *GiftCardService) GetReceivedGiftCards(ctx context.Context, receiverID int64, statusFilter string) (giftCards []*dto.GiftCard, err error) {
	giftCardsEntities, err := g.repository.GetReceivedGiftCards(ctx, receiverID, statusFilter)
	if err != nil {
		return nil, err
	}

	giftCards = make([]*dto.GiftCard, len(giftCardsEntities))
	for i, giftCardsEntity := range giftCardsEntities {
		giftCards[i] = dto.GiftCardDTOFromEntity(giftCardsEntity)
	}

	return giftCards, nil
}
