package model

import (
	db "backend/db/sqlc"
	"time"
)

type Genre struct {
	Id        string    `json:"id"`
	Genre     string    `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ConvertGenreFromRepository(g db.Genre) (*Genre, error) {
	return &Genre {
		Id: g.ID.String(),
		Genre: g.Genre,
	}, nil
}