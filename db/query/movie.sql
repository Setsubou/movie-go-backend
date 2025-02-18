-- name: GetMovieById :one
SELECT sqlc.embed(m), sqlc.embed(p), sqlc.embed(c)
FROM movies m
LEFT JOIN publisher p ON m.publisher_id = p.id
LEFT JOIN country c ON p.country_id = c.id
WHERE m.id = $1;

-- name: GetGenresByMovieId :many
SELECT sqlc.embed(g)
FROM movies m
LEFT JOIN movie_genres mg ON m.id = mg.movie_id
LEFT JOIN genres g ON mg.genre_id = g.id
WHERE m.id = $1;

-- name: GetAllMovies :many
SELECT sqlc.embed(m), sqlc.embed(p)
FROM movies m
LEFT JOIN publisher p ON m.publisher_id = p.id;