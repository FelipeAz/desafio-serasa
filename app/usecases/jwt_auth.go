package usecases

import (
	"net/http"

	"github.com/FelipeAz/desafio-serasa/app/entity"
	"github.com/dgrijalva/jwt-go"
)

// JWTService pertence a camada usecases.
type JWTService interface {
	CreateToken(entity.Access) (string, error)
	TokenValid(*http.Request) error
	VerifyToken(*http.Request) (*jwt.Token, error)
	ExtractToken(*http.Request) string
}
