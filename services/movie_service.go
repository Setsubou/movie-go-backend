package services

import (
	db "backend/db/sqlc"
	"backend/errors"
	"backend/model"
	"backend/repository"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type MovieService struct {
	repository repository.MovieRepository
}

func (s *MovieService) GetMovieById(id string) (*model.Movie, error) {
	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
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

func (s *MovieService) GetListAllMovies() ([]*model.Movie, error) {
	movies, err := s.repository.GetAllMovies()
	var converted_movies []*model.Movie

	if err != nil {
		return nil, errors.ErrNotFound.SetMessage("no movies found")
	}

	for _, movies := range *movies {
		movie := model.ConvertMovieFromDTO(movies.Movie)
		converted_publisher := model.ConvertPublisherFromDTO(movies.Publisher)

		movie.Publisher = converted_publisher
		converted_movies = append(converted_movies, &movie)
	}

	return converted_movies, nil
}

func (s *MovieService) DeleteMovieById(id string) error {
	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return errors.ErrBadRequest.SetMessage("Invalid movie ID")
	}

	if err := s.repository.DeleteMovieById(uuid); err != nil {
		return err
	}

	return nil
}

func (s *MovieService) GetMoviesByPublisherId(id string) ([]*model.Movie, error) {
	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return nil, errors.ErrBadRequest.SetMessage("Invalid publisher ID")
	}

	var converted_movies []*model.Movie
	movies, _ := s.repository.GetMoviesByPublisherId(uuid)

	if len(*movies) == 0 {
		return nil, errors.ErrNotFound.SetMessage("resource not found")
	}

	for _, movies := range *movies {
		movie := model.ConvertMovieFromDTO(movies.Movie)
		converted_publisher := model.ConvertPublisherFromDTO(movies.Publisher)

		movie.Publisher = converted_publisher
		converted_movies = append(converted_movies, &movie)
	}

	return converted_movies, nil
}

func (s *MovieService) InsertNewMovie(movie_data model.Movie) (string, error) {
	var score pgtype.Numeric
	if err := score.Scan(strconv.FormatFloat(movie_data.Score, 'f', -1, 64)); err != nil {
		return "", errors.ErrBadRequest.SetMessage("invalid movie score")
	}

	var date pgtype.Date
	if err := date.Scan(movie_data.ReleaseDate); err != nil {
		return "", errors.ErrBadRequest.SetMessage("invalid date format for release date")
	}

	var publisher_uuid pgtype.UUID
	if err := publisher_uuid.Scan(movie_data.Publisher.Id); err != nil {
		return "", errors.ErrBadRequest.SetMessage("invalid publisher ID")
	}

	random_uuid := uuid.NewString()
	var movie_uuid pgtype.UUID
	if err := movie_uuid.Scan(random_uuid); err != nil {
		return "", errors.ErrInternalError.SetMessage("failed to generate new id")
	}

	movie_param := db.InsertNewMovieParams{
		ID:          movie_uuid,
		Title:       movie_data.Title,
		Score:       score,
		Picture:     movie_data.Picture,
		ReleaseDate: date,
		Synopsis:    movie_data.Synopsis,
		PublisherID: publisher_uuid,
	}

	var genre_param []db.InsertNewMovieGenreParams
	for _, genre := range movie_data.Genre {

		var genre_uuid pgtype.UUID
		if err := genre_uuid.Scan(genre.Id); err != nil {
			return "", errors.ErrBadRequest.SetMessage("invalid genre id")
		}

		converted_genre := db.InsertNewMovieGenreParams{
			MovieID: movie_uuid,
			GenreID: genre_uuid,
		}

		genre_param = append(genre_param, converted_genre)
	}

	if err := s.repository.InsertNewMovie(movie_param, genre_param); err != nil {
		return "", err
	}

	return movie_uuid.String(), nil
}

func NewMovieService(repository repository.MovieRepository) *MovieService {
	return &MovieService{
		repository: repository,
	}
}
