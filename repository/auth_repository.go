package repository

type AuthRepository interface {
	GetOneUserByUsername(username string) string
}
