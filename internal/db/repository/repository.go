package repository

import (
	"database/sql"

	database "github.com/treboc/virtus/internal/db/generated"
)

type Repository struct {
	db *sql.DB
	q  *database.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
		q:  database.New(db),
	}
}
