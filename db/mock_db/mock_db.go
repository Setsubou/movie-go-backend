package db

import (
	db "backend/db/sqlc"
)

type Mock_db struct{}

func NewMockDb() *Mock_db {
	return &Mock_db{}
}

func (m *Mock_db) GetById(id string) (*db.GetMovieByIdRow, error) {
	return &db.GetMovieByIdRow{
		Movie: db.Movie{
			Title:       "The Matrix",
			Synopsis:    "A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.",
		},
	}, nil
}
