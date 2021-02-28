package interfaces

import (
	"net/http"

	"github.com/FelipeAz/desafio-serasa/app/entity"
	"github.com/FelipeAz/desafio-serasa/app/usecases"
	"github.com/gin-gonic/gin"
)

// UserController contem o servico de usuario.
type UserController struct {
	UserService usecases.UserService
	JWTService  usecases.JWTService
}

// NewUserController retorna o controller do login.
func NewUserController(sqlHandler SQLHandler, jwtService JWTService) *UserController {
	return &UserController{
		UserService: usecases.UserService{
			UserRepository: &UserRepository{
				SQLHandler: sqlHandler,
			},
		},
		JWTService: &jwtService,
	}
}

// Login Valida o login do usuario.
func (uc *UserController) Login(c *gin.Context) {
	var credential entity.User
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access, err := uc.UserService.Login(credential.Email, credential.Password, uc.JWTService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Login Failed": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": access})
}

// SignUp cria uma conta
func (uc *UserController) SignUp(c *gin.Context) {
	var input entity.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, err := uc.UserService.SignUp(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// Logout remove a sessao do usuario
func (uc *UserController) Logout(c *gin.Context) {
	var credential entity.User
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logout := uc.UserService.Logout(credential.Email, credential.Password)
	if logout == false {
		c.JSON(http.StatusOK, "User is not Logged In")
		return
	}

	c.JSON(http.StatusOK, "Successfully Logged Out")
}
