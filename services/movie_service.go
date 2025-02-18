package services

import (
	"backend/errors"
	"backend/model"
	"backend/repository"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewMovieService(repository repository.MovieRepository) *MovieService {
	return &MovieService{
		repository: repository,
	}
}

type MovieService struct {
	repository repository.MovieRepository
}

func (s *MovieService) GetMovieById(id string) (*model.Movie, error) {
	var uuid pgtype.UUID
	err := uuid.Scan(id)

	if err != nil {
		return nil, errors.ErrBadRequest.SetMessage("Invalid movie ID")
	}

	movie_data, movie_err := s.repository.GetMovieById(uuid)
	movie_genre, _ := s.repository.GetGenreByMovieId(uuid)

	if movie_err != nil {
		return nil, errors.ErrNotFound
	}

	converted_movie := model.ConvertMovieFromDTO(movie_data.Movie)
	converted_publisher := model.ConvertPublisherFromDTO(movie_data.Publisher)
	converted_country := model.ConvertCountryFromDTO(movie_data.Country)

	converted_movie.Publisher = converted_publisher
	converted_movie.Publisher.Country = converted_country

	for _, genre := range *movie_genre {
		converted_genre := model.ConvertGenreFromDTO(genre.Genre)

		converted_movie.Genre = append(converted_movie.Genre, converted_genre)
	}

	return &converted_movie, nil
}

func (s *MovieService) GetListAllMovies() []*model.Movie {
	movies, _ := s.repository.GetAllMovies()
	var converted_movies []*model.Movie

	for _, movies := range *movies {
		movie := model.ConvertMovieFromDTO(movies.Movie)
		converted_publisher := model.ConvertPublisherFromDTO(movies.Publisher)

		movie.Publisher = converted_publisher
		converted_movies = append(converted_movies, &movie)
	}

	return converted_movies
}
