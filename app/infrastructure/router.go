package infrastructure

import (
	"github.com/FelipeAz/desafio-serasa/app/interfaces"
	"github.com/FelipeAz/desafio-serasa/app/middleware"
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
	userController := interfaces.NewUserController(sqlHandler, *interfaces.NewJWTAuthService())
	r.router.POST("/signup", userController.SignUp)
	r.router.POST("/logout", userController.Logout)
	r.router.POST("/login", userController.Login)

	jwt := interfaces.NewJWTAuthService()
	negativacaoController := interfaces.NewNegativacaoController(sqlHandler)
	rg := r.router.Group("/negativacao")

	rg.GET("/", middleware.AuthorizeJWT(*jwt), negativacaoController.Find)
	rg.GET("/:id", middleware.AuthorizeJWT(*jwt), negativacaoController.FindByID)
	rg.POST("/", middleware.AuthorizeJWT(*jwt), negativacaoController.Persist)
	rg.PUT("/:id", middleware.AuthorizeJWT(*jwt), negativacaoController.Update)
	rg.DELETE("/:id", middleware.AuthorizeJWT(*jwt), negativacaoController.Destroy)

	r.listen()
}

func (r Router) listen() {
	r.router.Run(":8080")
}
