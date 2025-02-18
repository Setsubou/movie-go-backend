package model

import (
	db "backend/db/sqlc"
	"time"
)

type Publisher struct {
	Id             string    `json:"id,omitempty"`
	Publisher_name string    `json:"publisher_name,omitempty"`
	Year_Founded   int       `json:"year_founded,omitempty"`
	Country        Country   `json:"country,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

func ConvertPublisherFromRepository(p db.Publisher) (Publisher, error) {
	return Publisher{
		Id: p.ID.String(),
		Publisher_name: p.PublisherName,
		Year_Founded: int(p.YearFounded),
	}, nil
}