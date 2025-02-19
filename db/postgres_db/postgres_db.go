package db

import (
	"backend/db/sqlc"
	"backend/errors"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres_db struct {
	pgx_pool *pgxpool.Pool
}

func NewPostgresDb(connection_string string) *Postgres_db {
	pgx_pool, err := pgxpool.New(context.Background(), connection_string)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return &Postgres_db{
		pgx_pool: pgx_pool,
	}
}

// Implements Auth Repository
func (p *Postgres_db) GetOneUserByUsername(username string) string {
	ctx := context.Background()
	query := db.New(p.pgx_pool)

	count, _ := query.CheckUserCredentials(ctx, username)

	return count
}

// Implements Movie Repository
func (p *Postgres_db) InsertNewMovie(movie_param db.InsertNewMovieParams, genre_param []db.InsertNewMovieGenreParams) error {
	ctx := context.Background()
	tx, err := p.pgx_pool.Begin(ctx)

	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := db.New(tx)
	if err := qtx.InsertNewMovie(ctx, movie_param); err != nil {
		return err
	}

	for _, genre := range genre_param {
		if err := qtx.InsertNewMovieGenre(ctx, genre); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (p *Postgres_db) GetMovieById(uuid pgtype.UUID) (*db.GetMovieByIdRow, error) {
	ctx := context.Background()
	query := db.New(p.pgx_pool)

	movie, err := query.GetMovieById(ctx, uuid)

	if err != nil {
		return nil, errors.ErrNotFound
	}

	return &movie, nil
}

func (p *Postgres_db) GetGenreByMovieId(uuid pgtype.UUID) (*[]db.GetGenresByMovieIdRow, error) {
	ctx := context.Background()
	query := db.New(p.pgx_pool)

	movie_genre, err := query.GetGenresByMovieId(ctx, uuid)

	if err != nil {
		return nil, errors.ErrNotFound
	}

	return &movie_genre, nil
}

func (p *Postgres_db) GetAllMovies() (*[]db.GetAllMoviesRow, error) {
	ctx := context.Background()
	query := db.New(p.pgx_pool)

	movies, err := query.GetAllMovies(ctx)

	if err != nil {
		return nil, errors.ErrNotFound
	}

	return &movies, nil
}

func (p *Postgres_db) GetMoviesByPublisherId(uuid pgtype.UUID) (*[]db.GetMoviesByPublisherIdRow, error) {
	ctx := context.Background()
	query := db.New(p.pgx_pool)

	movie, err := query.GetMoviesByPublisherId(ctx, uuid)

	if err != nil {
		return nil, errors.ErrNotFound
	}

	return &movie, nil
}

func (p *Postgres_db) DeleteMovieById(id pgtype.UUID) error {
	ctx := context.Background()
	query := db.New(p.pgx_pool)

	err := query.DeleteMovieById(ctx, id)

	if err != nil {
		return errors.ErrNotFound
	}

	return nil
}

// Implements publisher repository
func (p *Postgres_db) GetListAllPublishersName() (*[]db.GetAllPublishersNameRow, error) {
	ctx := context.Background()
	query := db.New(p.pgx_pool)

	publishers, err := query.GetAllPublishersName(ctx)

	if err != nil {
		return nil, errors.ErrNotFound
	}

	return &publishers, nil
}