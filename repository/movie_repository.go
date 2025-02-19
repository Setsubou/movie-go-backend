package repository

import (
	db "backend/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type MovieRepository interface {
	InsertNewMovie(movie_param db.InsertNewMovieParams, genre_param []db.InsertNewMovieGenreParams) error

	GetAllMovies() (*[]db.GetAllMoviesRow, error)
	GetMovieById(id pgtype.UUID) (*db.GetMovieByIdRow, error)
	GetGenreByMovieId(id pgtype.UUID) (*[]db.GetGenresByMovieIdRow, error)
	GetMoviesByPublisherId(id pgtype.UUID)(*[]db.GetMoviesByPublisherIdRow, error)

	DeleteMovieById(id pgtype.UUID) error
}
