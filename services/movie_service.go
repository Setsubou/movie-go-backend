package movie_service

import (
	"backend/model"
	"backend/repository"
)

type MovieService struct {
	repository repository.MovieRepository
}

func NewMovieService(repository repository.MovieRepository) *MovieService {
	return &MovieService{
		repository: repository,
	}
}

func (s *MovieService) GetById(id int) *model.Movie {
	result, err := s.repository.GetById(id);

	if err != nil {
		return nil; //TODO error checking
	}
	return result;
}