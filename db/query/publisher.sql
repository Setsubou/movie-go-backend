-- name: GetAllPublishersName :many
SELECT p.id, p.publisher_name
FROM publisher p;