package repository

import "backend/model"

type MovieRepository interface {
	GetById(id int) (*model.Movie, error)
}