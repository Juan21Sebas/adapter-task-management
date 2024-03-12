package repository

import (
	"database/sql"
)

type BDRepository struct {
	db *sql.DB
}

func NewBdRepository(db *sql.DB) *BDRepository {
	return &BDRepository{
		db: db,
	}
}
