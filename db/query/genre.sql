-- name: GetAllGenres :many
SELECT sqlc.embed(g)
FROM genres g;