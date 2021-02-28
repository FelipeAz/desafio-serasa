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
func AuthorizeJWT(jwtService interfaces.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwtService.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Voce nao possui autorizacao para acessar essa rota")
			c.Abort()
			return
		}
		c.Next()
	}
}
