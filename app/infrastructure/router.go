package infrastructure

import (
	"github.com/FelipeAz/desafio-serasa/app/interfaces"
	"github.com/gin-gonic/gin"
)

// Router .
type Router struct {
	router *gin.Engine
}

// NewRouter retorna uma instancia do Router
func NewRouter() interfaces.Router {
	return Router{router: gin.Default()}
}

// Dispatch ativa as rota
func (r Router) Dispatch(sqlHandler interfaces.SQLHandler) {
	rg := r.router.Group("/negativacao")

	negativacaoController := interfaces.NewNegativacaoController(sqlHandler)

	rg.GET("/", negativacaoController.Find)
	rg.GET("/:id", negativacaoController.FindByID)
	rg.POST("/", negativacaoController.Persist)
	rg.PUT("/:id", negativacaoController.Update)
	rg.DELETE("/:id", negativacaoController.Destroy)

	r.listen()
}

func (r Router) listen() {
	r.router.Run(":8080")
}
