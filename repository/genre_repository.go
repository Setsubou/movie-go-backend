package repository

import db "backend/db/sqlc"

type GenreRepository interface {
	GetAllGenres() (*[]db.GetAllGenresRow, error)
}