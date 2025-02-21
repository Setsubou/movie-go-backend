package services

import (
	"backend/errors"
	"backend/model"
	"backend/repository"
)

type GenreService struct {
	repository repository.GenreRepository
}

func (s *GenreService) GetAllGenres() ([]*model.Genre, error) {
	genre_repository, err := s.repository.GetAllGenres()

	if err != nil {
		return nil, errors.ErrNotFound.SetMessage("genre resource not found")
	}

	var genres []*model.Genre
	for _, value := range *genre_repository {
		converted_genre := model.ConvertGenreFromDTO(value.Genre)

		genres = append(genres, &converted_genre)
	}

	return genres, nil
}

func NewGenreService(repository repository.GenreRepository) *GenreService {
	return &GenreService{
		repository: repository,
	}
}
