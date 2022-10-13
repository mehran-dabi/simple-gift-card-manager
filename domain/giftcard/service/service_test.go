package service

import (
	"context"
	"dono/domain/giftcard/entity"
	repositoryMocks "dono/mocks/domain/giftcard/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
}

func (s *ServiceTestSuite) TestAddGiftCard() {
	giftCardEntity := &entity.GiftCard{
		ID:        1,
		Price:     200,
		Sender:    0,
		Receiver:  0,
		Status:    "",
		CreatedAt: "",
		UpdatedAt: "",
	}
	repositoryMock := repositoryMocks.IGiftCardRepository{}
	repositoryMock.On("AddGiftCard", mock.Anything, giftCardEntity.Price).Return(giftCardEntity, nil)

	giftCardService := NewGiftCardService(&repositoryMock)
	giftCardDTO, err := giftCardService.AddGiftCard(context.Background(), 200)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), giftCardDTO)
}

func (s *ServiceTestSuite) TestSendGiftCard() {
	giftCardEntity := &entity.GiftCard{
		ID:        1,
		Price:     200,
		Sender:    0,
		Receiver:  0,
		Status:    "",
		CreatedAt: "",
		UpdatedAt: "",
	}
	repositoryMock := repositoryMocks.IGiftCardRepository{}
	repositoryMock.On("AddGiftCard", mock.Anything, giftCardEntity.Price).Return(giftCardEntity, nil)
	repositoryMock.On("SendGiftCard", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	giftCardService := NewGiftCardService(&repositoryMock)
	err := giftCardService.SendGiftCard(context.Background(), 200, 1, 1)
	assert.NoError(s.T(), err)
}

func (s *ServiceTestSuite) TestUpdateGiftCardStatus() {
	repositoryMock := repositoryMocks.IGiftCardRepository{}
	repositoryMock.On("UpdateGiftCardStatus", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	giftCardService := NewGiftCardService(&repositoryMock)
	err := giftCardService.UpdateGiftCardStatus(context.Background(), 1, "ACCEPTED")
	assert.NoError(s.T(), err)
}

func (s *ServiceTestSuite) TestGetReceivedGiftCards() {
	giftCardEntities := []*entity.GiftCard{
		{1, 200, 3, 2, "REJECTED", "2022-10-11 16:49:04", "2022-10-12 14:11:31"},
		{2, 300, 1, 2, "ACCEPTED", "2022-10-11 16:49:04", "2022-10-12 14:11:31"},
	}
	repositoryMock := repositoryMocks.IGiftCardRepository{}
	repositoryMock.On("GetReceivedGiftCards", mock.Anything, mock.Anything, mock.Anything).Return(giftCardEntities, nil)

	giftCardService := NewGiftCardService(&repositoryMock)
	giftCardDTOs, err := giftCardService.GetReceivedGiftCards(context.Background(), 1, "")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), giftCardDTOs)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
