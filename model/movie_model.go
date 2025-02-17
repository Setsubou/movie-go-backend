package model

import (
	db "backend/db/sqlc"
	"fmt"
	"time"
)

type Movie struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Score       float64   `json:"score"`
	Picture     string    `json:"picture"`
	ReleaseDate time.Time `json:"release_date"`
	Synopsis    string    `json:"synopsis"`
	Publisher   Publisher `json:"publisher"`
	Genre       []Genre   `json:"genre"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ConvertMovieFromRepository(m db.Movie) (*Movie, error) {
	score, err := m.Score.Float64Value()

	if err != nil {
		fmt.Println("Unable to convert to float")
		return nil, err
	}

	return &Movie{
		Id:          m.ID.String(),
		Title:       m.Title,
		Score:       score.Float64,
		Picture:     m.Picture,
		ReleaseDate: m.ReleaseDate.Time,
		Synopsis:    m.Synopsis,
	}, nil
}
