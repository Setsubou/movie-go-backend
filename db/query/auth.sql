-- name: CheckUserCredentials :one
SELECT u.password_hash
FROM users u
WHERE u.user_name = $1;