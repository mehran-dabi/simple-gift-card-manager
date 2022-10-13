package repository

import (
	"context"
	"database/sql"
	"dono/helper"
	mocks "dono/mocks/infrastructure"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func (s *RepositoryTestSuite) TestAddGiftCard() {
	testCases := []struct {
		price int64
		id    int64
	}{
		{
			id:    1,
			price: 200,
		},
		{
			id:    2,
			price: 100,
		},
	}
	s.db, s.mock = mocks.NewDBMock()
	giftCardRepository := NewGiftCardRepository(s.db)

	for _, tc := range testCases {
		query := helper.ReplaceEscapeCharacter(InsertGiftCard)
		s.mock.ExpectExec(query).
			WithArgs(tc.price).
			WillReturnResult(sqlmock.NewResult(tc.id, 1))
		_, err := giftCardRepository.AddGiftCard(context.Background(), tc.price)
		assert.NoError(s.T(), err)
	}
}

func (s *RepositoryTestSuite) TestSendGiftCard() {
	testCases := []struct {
		senderID   int64
		receiverID int64
		giftCardID int64
	}{
		{
			senderID:   1,
			receiverID: 2,
			giftCardID: 1,
		},
		{
			senderID:   3,
			receiverID: 2,
			giftCardID: 1,
		},
	}

	s.db, s.mock = mocks.NewDBMock()
	giftCardRepository := NewGiftCardRepository(s.db)

	for _, tc := range testCases {
		query := helper.ReplaceEscapeCharacter(UpdateGiftCard)
		s.mock.ExpectExec(query).
			WithArgs(tc.senderID, tc.receiverID, tc.giftCardID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		err := giftCardRepository.SendGiftCard(context.Background(), tc.senderID, tc.receiverID, tc.giftCardID)
		assert.NoError(s.T(), err)
	}
}

func (s *RepositoryTestSuite) TestReceivedGiftCards() {
	testCases := []struct {
		receiverID   int64
		statusFilter string
	}{
		{
			receiverID:   1,
			statusFilter: "",
		},
		{
			receiverID:   1,
			statusFilter: "PENDING",
		},
		{
			receiverID:   1,
			statusFilter: "REJECTED",
		},
	}

	s.db, s.mock = mocks.NewDBMock()
	giftCardRepository := NewGiftCardRepository(s.db)

	for _, tc := range testCases {
		rows := s.mock.NewRows([]string{"id", "price", "sender", "receiver", "status", "created_at", "updated_at"}).
			AddRow(1, 200, 1, 2, "PENDING", time.Now(), time.Now())

		var query string
		if tc.statusFilter != "" {
			query = helper.ReplaceEscapeCharacter(GetReceivedGiftCardsWithStatusFilter)
			s.mock.ExpectQuery(query).
				WithArgs(tc.receiverID, tc.statusFilter).
				WillReturnRows(rows)
		} else {
			query = helper.ReplaceEscapeCharacter(GetReceivedGiftCards)
			s.mock.ExpectQuery(query).
				WithArgs(tc.receiverID).
				WillReturnRows(rows)
		}
		giftCards, err := giftCardRepository.GetReceivedGiftCards(context.Background(), tc.receiverID, tc.statusFilter)
		assert.NoError(s.T(), err)
		assert.NotNil(s.T(), giftCards)
	}
}

func (s *RepositoryTestSuite) TestUpdateGiftCardStatus() {
	testCases := []struct {
		id     int64
		status string
	}{
		{
			id:     1,
			status: "ACCEPTED",
		},
		{
			id:     1,
			status: "REJECTED",
		},
	}

	s.db, s.mock = mocks.NewDBMock()
	giftCardRepository := NewGiftCardRepository(s.db)

	for _, tc := range testCases {
		query := helper.ReplaceEscapeCharacter(UpdateGiftCardStatus)
		s.mock.ExpectExec(query).
			WithArgs(tc.status, tc.id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		err := giftCardRepository.UpdateGiftCardStatus(context.Background(), tc.id, tc.status)
		assert.NoError(s.T(), err)
	}
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
