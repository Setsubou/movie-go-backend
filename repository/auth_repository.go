package repository

type AuthService interface {
	GetOneUserByUsername(username string) string
}
