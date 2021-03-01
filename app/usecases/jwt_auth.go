package usecases

import (
	"net/http"

	"github.com/FelipeAz/desafio-serasa/app/entity"
	"github.com/dgrijalva/jwt-go"
)

// JWTAuth pertence a camada usecases.
type JWTAuth interface {
	CreateToken(entity.Access) (*entity.TokenDetails, error)
	TokenValid(*http.Request) error
	VerifyToken(*http.Request) (*jwt.Token, error)
	ExtractToken(*http.Request) string
	FetchToken(string, *http.Request) bool
}
