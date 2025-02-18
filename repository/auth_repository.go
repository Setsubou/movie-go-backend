package repository

type AuthRepository interface {
	GetOneUserByUsernameAndPasswordHash(username string) string
}
