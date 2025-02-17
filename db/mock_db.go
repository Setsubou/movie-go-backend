package db

import (
	"backend/model"
	"time"
)

type Mock_db struct{}

func NewMockDb() *Mock_db {
	return &Mock_db{}
}

func (m *Mock_db) GetById(id int) (*model.Movie, error) {
	return &model.Movie{
		Id:          id,
		Title:       "The Matrix",
		Score:       8.0,
		Picture:     []byte{},
		ReleaseDate: time.Now(),
		Synopsis:    "A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.",
		Publisher:   "Lana Wachowski, Lilly Wachowski",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
