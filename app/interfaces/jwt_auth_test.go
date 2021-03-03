package interfaces

import (
	"net/http"
	"testing"
	"time"

	"github.com/FelipeAz/desafio-serasa/app/entity"
)

// Testa a funcao de geracao de access token
func TestGenerateAccessTokenString(t *testing.T) {
	// Init
	atExpires := time.Now().Add(time.Minute * 15).Unix()
	access := entity.Access{
		UserID:       1,
		AccessToken:  "",
		RefreshToken: "",
	}

	// Execution
	aToken, _ := generateAccessTokenString(access, atExpires)

	// Validation
	if aToken == "" {
		t.Error("token de acesso nao foi gerado")
	}
}

// Testa a funcao de geracao de refresh token
func TestGenerateRefreshToken(t *testing.T) {
	// Init
	rtExpires := time.Now().Add(time.Hour * 24 * 7).Unix()
	access := entity.Access{
		UserID:       1,
		AccessToken:  "",
		RefreshToken: "",
	}

	// Execution
	rToken, _ := generateAccessTokenString(access, rtExpires)

	// Validation
	if rToken == "" {
		t.Error("token de refresh nao foi gerado")
	}
}

// Testa a extracao do token da requisicao
func TestExtractToken(t *testing.T) {
	// Init
	atExpires := time.Now().Add(time.Minute * 15).Unix()
	jwt := NewJWTAuth(nil)
	access := entity.Access{
		UserID: 1,
	}
	aToken, _ := generateAccessTokenString(access, atExpires)

	req, _ := http.NewRequest("GET", "localhost:8080/negativacao", nil)
	req.Header.Add("Authorization", "Bearer "+aToken)

	// Execution
	tokenStr := jwt.ExtractToken(req)

	// Validation
	if tokenStr != aToken {
		t.Errorf("token esperado `%s` token recebido `%s`", aToken, tokenStr)
	}
}

// Testa o metodo signing do token
func TestVerifyToken(t *testing.T) {
	// Init
	atExpires := time.Now().Add(time.Minute * 15).Unix()
	access := entity.Access{
		UserID: 1,
	}

	aToken, _ := generateAccessTokenString(access, atExpires)

	req, _ := http.NewRequest("GET", "localhost:8080/negativacao", nil)
	req.Header.Add("Authorization", "Bearer "+aToken)

	// Execution
	_, isValid := verifyToken(aToken)

	// Validation
	if isValid != nil {
		t.Error("token signing invalido")
	}
}

// Testa se o token de refresh gerado esta sendo valido
func TestRefreshToken(t *testing.T) {
	// Init
	jwt := NewJWTAuth(nil)
	rtExpires := time.Now().Add(time.Minute * 15).Unix()
	access := entity.Access{
		UserID: 1,
	}

	result := true
	rToken, _ := generateRefreshTokenString(access, rtExpires)

	// Execution
	isValid := jwt.refreshToken(rToken)

	if isValid != result {
		t.Error("refresh token invalido")
	}
}
