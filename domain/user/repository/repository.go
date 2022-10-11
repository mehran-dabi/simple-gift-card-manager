package repository

import (
	"database/sql"
)

type IUserRepository interface {
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}
