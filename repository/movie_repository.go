package repository

import (
	db "backend/db/sqlc"
)

type MovieRepository interface {
	GetMovieById(id string) (*db.GetMovieByIdRow, error)
	GetGenreByMovieId(id string) (*[]db.GetGenresByMovieIdRow, error)
}
