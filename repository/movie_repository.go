package repository

import (
	db "backend/db/sqlc"
)

type MovieRepository interface {
	GetAllMovies() (*[]db.GetAllMoviesRow, error)
	GetMovieById(id string) (*db.GetMovieByIdRow, error)
	GetGenreByMovieId(id string) (*[]db.GetGenresByMovieIdRow, error)
}
