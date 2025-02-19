package model

import (
	db "backend/db/sqlc"
	"time"
)

type Movie struct {
	Id          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty" binding:"required"`
	Score       float64   `json:"score,omitempty" binding:"required"`
	Picture     string    `json:"picture,omitempty" binding:"required"`
	ReleaseDate time.Time `json:"release_date,omitempty" binding:"required"`
	Synopsis    string    `json:"synopsis,omitempty" binding:"required"`
	Publisher   Publisher `json:"publisher,omitempty" binding:"required"`
	Genre       []Genre   `json:"genre,omitempty" binding:"required"`
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
