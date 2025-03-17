package postgres

import "github.com/jmoiron/sqlx"

type AuthStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(db *sqlx.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}
