package service

type TokenService interface {
	GenerateToken(userId string) (string, error)
	VerifyToken(token string) (string, error)
}