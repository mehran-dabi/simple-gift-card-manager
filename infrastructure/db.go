package infrastructure

import (
	"database/sql"
	giftCardRepository "dono/domain/giftcard/repository"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Repositories struct {
	GiftCard giftCardRepository.IGiftCardRepository
	db       *sql.DB
}

func NewRepository(dbUser, dbPassword, dbHost, dbPort, dbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", DBURL)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		GiftCard: giftCardRepository.NewGiftCardRepository(db),
		db:       db,
	}, nil
}

func (s *Repositories) Close() error {
	return s.db.Close()
}
