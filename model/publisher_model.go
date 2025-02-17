package model

import (
	db "backend/db/sqlc"
	"time"
)

type Publisher struct {
	Id             string    `json:"id"`
	Publisher_name string    `json:"publisher_name"`
	Year_Founded   int       `json:"year_founded"`
	Country        Country   `json:"country"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ConvertPublisherFromRepository(p db.Publisher) (*Publisher, error) {
	return &Publisher{
		Id: p.ID.String(),
		Publisher_name: p.PublisherName,
		Year_Founded: int(p.YearFounded),
	}, nil
}