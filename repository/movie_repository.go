package repository

import (
	db "backend/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type MovieRepository interface {
	GetAllMovies() (*[]db.GetAllMoviesRow, error)
	GetMovieById(id pgtype.UUID) (*db.GetMovieByIdRow, error)
	GetGenreByMovieId(id pgtype.UUID) (*[]db.GetGenresByMovieIdRow, error)
}
