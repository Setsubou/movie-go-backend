// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: movie.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getGenresByMovieId = `-- name: GetGenresByMovieId :many
SELECT g.id, g.genre, g.created_at, g.updated_at
FROM movies m
LEFT JOIN movie_genres mg ON m.id = mg.movie_id
LEFT JOIN genres g ON mg.genre_id = g.id
WHERE m.id = $1
`

type GetGenresByMovieIdRow struct {
	Genre Genre
}

func (q *Queries) GetGenresByMovieId(ctx context.Context, id pgtype.UUID) ([]GetGenresByMovieIdRow, error) {
	rows, err := q.db.Query(ctx, getGenresByMovieId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGenresByMovieIdRow
	for rows.Next() {
		var i GetGenresByMovieIdRow
		if err := rows.Scan(
			&i.Genre.ID,
			&i.Genre.Genre,
			&i.Genre.CreatedAt,
			&i.Genre.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMovieById = `-- name: GetMovieById :one
SELECT m.id, m.title, m.score, m.picture, m.release_date, m.synopsis, m.publisher_id, m.created_at, m.updated_at, p.id, p.publisher_name, p.year_founded, p.country_id, p.created_at, p.updated_at, c.id, c.country_name
FROM movies m
LEFT JOIN publisher p ON m.publisher_id = p.id
LEFT JOIN country c ON p.country_id = c.id
WHERE m.id = $1
`

type GetMovieByIdRow struct {
	Movie     Movie
	Publisher Publisher
	Country   Country
}

func (q *Queries) GetMovieById(ctx context.Context, id pgtype.UUID) (GetMovieByIdRow, error) {
	row := q.db.QueryRow(ctx, getMovieById, id)
	var i GetMovieByIdRow
	err := row.Scan(
		&i.Movie.ID,
		&i.Movie.Title,
		&i.Movie.Score,
		&i.Movie.Picture,
		&i.Movie.ReleaseDate,
		&i.Movie.Synopsis,
		&i.Movie.PublisherID,
		&i.Movie.CreatedAt,
		&i.Movie.UpdatedAt,
		&i.Publisher.ID,
		&i.Publisher.PublisherName,
		&i.Publisher.YearFounded,
		&i.Publisher.CountryID,
		&i.Publisher.CreatedAt,
		&i.Publisher.UpdatedAt,
		&i.Country.ID,
		&i.Country.CountryName,
	)
	return i, err
}
