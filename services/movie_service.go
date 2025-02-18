package services

import (
	"backend/model"
	"backend/repository"
)

func NewMovieService(repository repository.MovieRepository) *MovieService {
	return &MovieService{
		repository: repository,
	}
}

type MovieService struct {
	repository repository.MovieRepository
}

func (s *MovieService) GetMovieById(uuid string) *model.Movie {
	movie_data, err := s.repository.GetMovieById(uuid)
	movie_genre, err := s.repository.GetGenreByMovieId(uuid)

	if err != nil {
		return nil //TODO error checking
	}

	converted_movie, err := model.ConvertMovieFromRepository(*&movie_data.Movie)
	converted_publisher, err := model.ConvertPublisherFromRepository(*&movie_data.Publisher)
	converted_country, err := model.ConvertCountryFromRepository(*&movie_data.Country)

	converted_movie.Publisher = converted_publisher
	converted_movie.Publisher.Country = converted_country

	for _, v := range *movie_genre {
		converted_genre, _ := model.ConvertGenreFromRepository(v.Genre)

		converted_movie.Genre = append(converted_movie.Genre, converted_genre)
	}

	return &converted_movie
}

func (s *MovieService) GetListAllMovies() []*model.Movie {
	movies, _ := s.repository.GetAllMovies();
	var converted_movies []*model.Movie

	for _, v := range *movies {
		movie, _ := model.ConvertMovieFromRepository(v.Movie)
		converted_publisher, _ := model.ConvertPublisherFromRepository(v.Publisher)
	
		movie.Publisher = converted_publisher
		converted_movies = append(converted_movies, &movie)
	}

	return converted_movies
}

