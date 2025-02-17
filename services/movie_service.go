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

func (s *MovieService) GetById(uuid string) *model.Movie {
	movie_result, err := s.repository.GetMovieById(uuid)
	movie_genre, err := s.repository.GetGenreByMovieId(uuid)

	if err != nil {
		return nil //TODO error checking
	}

	converted_movie, err := model.ConvertMovieFromRepository(*&movie_result.Movie)
	converted_publisher, err := model.ConvertPublisherFromRepository(*&movie_result.Publisher)
	converted_country, err := model.ConvertCountryFromRepository(*&movie_result.Country)

	converted_movie.Publisher = *converted_publisher
	converted_movie.Publisher.Country = *converted_country

	for _, v := range *movie_genre {
		converted_genre, _ := model.ConvertGenreFromRepository(v.Genre)

		converted_movie.Genre = append(converted_movie.Genre, *converted_genre)
	}

	return converted_movie
}
