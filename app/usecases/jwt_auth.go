package usecases

import (
	"net/http"

	"github.com/FelipeAz/desafio-serasa/app/entity"
	"github.com/dgrijalva/jwt-go"
)

// TokenDetails detalhe dos tokens
type TokenDetails struct {
	AccessToken  string
	AtExpires    int64
	RefreshToken string
	RtExpires    int64
}

// JWTAuth pertence a camada usecases.
type JWTAuth interface {
	CreateToken(entity.Access) (*TokenDetails, error)
	TokenValid(*http.Request) error
	VerifyToken(*http.Request) (*jwt.Token, error)
	ExtractToken(*http.Request) string
	FetchToken(string) bool
}
