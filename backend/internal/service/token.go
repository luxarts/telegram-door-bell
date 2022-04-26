package service

import (
	"os"
	"strconv"
	"telegram-door-bell/internal/defines"
	"telegram-door-bell/internal/utils/jwt"
	"time"
)

type TokenService interface {
	Generate(ID int64, issuedAt time.Time) string
}

type tokenService struct {
}

func NewTokenService() TokenService {
	return &tokenService{}
}
func (s *tokenService) Generate(ID int64, issuedAt time.Time) string {
	return jwt.GenerateToken(jwt.Payload{
		Subject:  strconv.FormatInt(ID, 10),
		IssuedAt: issuedAt.UTC().Unix(),
	}, os.Getenv(defines.EnvTokenSecret))
}
