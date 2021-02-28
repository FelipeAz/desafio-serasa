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

	usr, err := uc.UserService.Login(credential.Email, credential.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := uc.JWTService.CreateToken(entity.Access{})
	access := uc.UserService.AuthUser(usr.ID, tokenString)

	c.JSON(http.StatusOK, gin.H{"token": access})
}

// SignUp cria uma conta
func (uc *UserController) SignUp(c *gin.Context) {
	var input entity.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	usr := entity.User{
		Email:    input.Email,
		Password: input.Password,
	}

	uc.UserService.SignUp(&usr)

	c.JSON(http.StatusOK, gin.H{"data": usr})
}

// Logout remove a sessao do usuario
func (uc *UserController) Logout(c *gin.Context) {
	var credential entity.User
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logout := uc.UserService.Logout(credential.Email, credential.Password)
	c.JSON(http.StatusOK, gin.H{"Logout": logout})
}
