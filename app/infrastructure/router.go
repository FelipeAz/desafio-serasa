package infrastructure

import (
	"github.com/FelipeAz/desafio-serasa/app/interfaces"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Dispatch ativa as rota
func Dispatch(sqlHandler interfaces.SQLHandler) {
	r := router.Group("/negativacao")

	negativacaoController := interfaces.NewNegativacaoController(sqlHandler)

	r.GET("/", negativacaoController.GetAll)
	r.GET("/:id", negativacaoController.Get)
	r.POST("/", negativacaoController.Save)
	r.PUT("/:id", negativacaoController.Change)
	r.DELETE("/:id", negativacaoController.Destroy)
}
