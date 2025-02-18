package model

import (
	db "backend/db/sqlc"
	"time"
)

type Genre struct {
	Id        string    `json:"id,omitempty"`
	Genre     string    `json:"genre,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func ConvertGenreFromDTO(g db.Genre) Genre {
	return Genre{
		Id:    g.ID.String(),
		Genre: g.Genre,
	}
}
