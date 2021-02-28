package interfaces

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/FelipeAz/desafio-serasa/app/entity"
	"github.com/dgrijalva/jwt-go"
)

// JWTService contem variaveis do servico.
type JWTService struct {
	issure string
}

// NewJWTAuthService instancia do servico.
func NewJWTAuthService() *JWTService {
	return &JWTService{
		issure: "Felipe",
	}
}

// CreateToken cria um token JWT.
func (service *JWTService) CreateToken(auth entity.Access) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["token"] = auth.AccessToken
	claims["user_id"] = auth.UserID
	claims["Issuer"] = service.issure
	claims["ExpiresAt"] = time.Now().Add(time.Hour * 2).Unix()
	claims["IssuedAt"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// TokenValid retorna se o token eh valido.
func (service *JWTService) TokenValid(r *http.Request) error {
	token, err := service.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// VerifyToken extrai o token e verifica.
func (service *JWTService) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := service.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken extrai o token da requisicao.
func (service *JWTService) ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
