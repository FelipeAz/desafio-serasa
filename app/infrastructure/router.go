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

	rg.GET("/", negativacaoController.GetAll)
	rg.GET("/:id", negativacaoController.Get)
	rg.POST("/", negativacaoController.Save)
	rg.PUT("/:id", negativacaoController.Change)
	rg.DELETE("/:id", negativacaoController.Destroy)
}

// Listen inicia o servidor
func (r Router) Listen() {
	r.router.Run(":8080")
}
