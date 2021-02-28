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
func AuthorizeJWT(jwtService interfaces.JWTAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := jwtService.ExtractToken(c.Request)
		err := jwtService.TokenValid(c.Request)
		if err != nil || jwtService.FetchToken(token) == false {
			c.JSON(http.StatusUnauthorized, "Voce nao possui autorizacao para acessar essa rota")
			c.Abort()
			return
		}
		c.Next()
	}
}
