package middleware

import (
	"net/http"

	"github.com/FelipeAz/desafio-serasa/app/interfaces"
	"github.com/gin-gonic/gin"
)

const (
	bearerSchema = "Bearer"
)

// AuthorizeJWT valida o token da requisicao.
func AuthorizeJWT(jwt *interfaces.JWTAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := jwt.ExtractToken(c.Request)
		fetched := jwt.FetchToken(tokenString) // fetched indica se achou no banco
		err := jwt.TokenValid(c.Request)
		if err != nil || !fetched {
			c.JSON(http.StatusUnauthorized, "Voce nao possui autorizacao para acessar essa rota")
			c.Abort()
			return
		}
		c.Next()
	}
}
