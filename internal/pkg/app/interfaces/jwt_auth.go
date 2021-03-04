package interfaces

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/FelipeAz/desafio-serasa/internal/pkg/app/entity"
	"github.com/dgrijalva/jwt-go"
)

// JWTAuth contem variaveis do servico.
type JWTAuth struct {
	SQLHandler SQLHandler
}

// NewJWTAuth instancia do servico.
func NewJWTAuth(sqlHandler SQLHandler) *JWTAuth {
	return &JWTAuth{
		SQLHandler: sqlHandler,
	}
}

// CreateToken cria um token JWT.
func (jwtAuth *JWTAuth) CreateToken(auth entity.Access) (td *entity.TokenDetails, err error) {
	td = &entity.TokenDetails{
		AtExpires: time.Now().Add(time.Minute * 15).Unix(),
		RtExpires: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	aToken, err := generateAccessTokenString(auth, td.AtExpires)
	if err != nil {
		return nil, err

	}
	td.AccessToken = aToken

	rToken, err := generateRefreshTokenString(auth, td.RtExpires)
	if err != nil {
		return nil, err

	}
	td.RefreshToken = rToken

	return
}

func generateAccessTokenString(auth entity.Access, expires int64) (aToken string, err error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["token"] = auth.AccessToken
	claims["user_id"] = auth.UserID
	claims["issuer"] = "Felipe"
	claims["exp"] = expires
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	aToken, err = token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return
}

func generateRefreshTokenString(auth entity.Access, expires int64) (rToken string, err error) {
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_token"] = auth.RefreshToken
	rtClaims["user_id"] = auth.UserID
	rtClaims["sub"] = 1
	rtClaims["exp"] = expires
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rToken, err = token.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))

	return
}

// TokenValid retorna se o token eh valido, considerando seu tempo de expiracao
func (jwtAuth *JWTAuth) TokenValid(r *http.Request) error {
	tokenStr := jwtAuth.ExtractToken(r)
	token, err := verifyToken(tokenStr)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

// Verifica se o metodo de signing do token eh valido
func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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
func (jwtAuth *JWTAuth) ExtractToken(r *http.Request) string {
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

// FetchToken verifica se existe um usuario logado no BD atraves do token da requisicao.
// Se o usuario estiver logado mas o token for invalido, revalidamos sua sessao SE o refresh token estiver valido.
// Caso o usuario deslogue, sera necessario logar novamente e utilizar um novo token, pois o mesmo eh apagado quando
// damos logout.
func (jwtAuth *JWTAuth) FetchToken(token string, r *http.Request) bool {
	var access entity.Access
	var isValid = false
	db := jwtAuth.SQLHandler.GetGorm()
	if err := db.Where("access_token=?", token).First(&access).Error; err != nil {
		return isValid
	}

	isValid = true
	err := jwtAuth.TokenValid(r)
	if err != nil {
		isValid = jwtAuth.refreshToken(access.RefreshToken)
	}

	return isValid
}

// refreshToken verifica se o refresh token eh valido, se sim, o usuario mesmo com o access token
// expirado pode ter acesso ao sistema.
func (jwtAuth *JWTAuth) refreshToken(refreshToken string) bool {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int(claims["sub"].(float64)) == 1
	}

	if err != nil {
		return false
	}

	return false
}
