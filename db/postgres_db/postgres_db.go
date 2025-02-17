package db

import (
	"backend/db/sqlc"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres_db struct{
	pgx_pool *pgxpool.Pool
}

func NewPostgresDb(pgx_pool *pgxpool.Pool) *Postgres_db {
	return &Postgres_db{
		pgx_pool: pgx_pool,
	}
}

func (p *Postgres_db) GetMovieById(id string) (*db.GetMovieByIdRow, error) {
	ctx := context.Background()

	query := db.New(p.pgx_pool)

	var UUID pgtype.UUID
	err := UUID.Scan(id)

	if err != nil {
		fmt.Println("Invalid UUID")
		return nil, err
	}

	movie, err := query.GetMovieById(ctx, UUID)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (p *Postgres_db) GetGenreByMovieId(id string) (*[]db.GetGenresByMovieIdRow, error) {
	ctx := context.Background()

	query := db.New(p.pgx_pool)

	var UUID pgtype.UUID
	err := UUID.Scan(id)

	if err != nil {
		fmt.Println("Invalid UUID")
		return nil, err
	}

	movie_genre, err := query.GetGenresByMovieId(ctx, UUID)

	if err != nil {
		fmt.Println("Error when fetching Genre by Movie ID")
		return nil, err
	}

	return &movie_genre, nil
}