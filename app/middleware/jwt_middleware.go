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
func AuthorizeJWT(jwt interfaces.JWTAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt.TokenValid(c.Request)
		tokenString := jwt.ExtractToken(c.Request)
		if err != nil && jwt.FetchToken(tokenString) == false {
			c.JSON(http.StatusUnauthorized, "Voce nao possui autorizacao para acessar essa rota")
			c.Abort()
			return
		}
		c.Next()
	}
}
