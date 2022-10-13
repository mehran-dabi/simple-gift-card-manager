package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Repositories struct {
	db *sql.DB
}

func NewRepository(dbUser, dbPassword, dbHost, dbPort, dbName string) (*sql.DB, error) {
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", DBURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Repositories) Close() error {
	return s.db.Close()
}
