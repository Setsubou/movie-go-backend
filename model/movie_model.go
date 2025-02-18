package model

import (
	db "backend/db/sqlc"
	"time"
)

type Movie struct {
	Id          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Score       float64   `json:"score,omitempty"`
	Picture     string    `json:"picture,omitempty"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Synopsis    string    `json:"synopsis,omitempty"`
	Publisher   Publisher `json:"publisher,omitempty"`
	Genre       []Genre   `json:"genre,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func ConvertMovieFromDTO(m db.Movie) Movie {
	score, _ := m.Score.Float64Value()

	return Movie{
		Id:          m.ID.String(),
		Title:       m.Title,
		Score:       score.Float64,
		Picture:     m.Picture,
		ReleaseDate: m.ReleaseDate.Time,
		Synopsis:    m.Synopsis,
	}
}
