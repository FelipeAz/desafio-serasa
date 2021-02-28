package usecases

import (
	"github.com/FelipeAz/desafio-serasa/app/entity"
)

// UserService pertence a camada usecases.
type UserService struct {
	UserRepository UserRepository
}

// Login do usuario no sistema.
func (us *UserService) Login(email, password string, jwt JWTService) (access entity.Access, err error) {
	usr, err := us.UserRepository.Login(email, password)
	if err != nil {
		return entity.Access{}, err
	}

	tokenString, err := jwt.CreateToken(entity.Access{})
	access, err = us.AuthUser(usr.ID, tokenString)

	return
}

// AuthUser insere o token de autenticacao no banco.
func (us *UserService) AuthUser(id uint, token string) (access entity.Access, err error) {
	access, err = us.UserRepository.AuthUser(id, token)
	return
}

// SignUp registra um novo usuario ao sistema.
func (us *UserService) SignUp(usr *entity.User) (user *entity.User, err error) {
	user, err = us.UserRepository.SignUp(usr)
	return
}

// Logout desloga o usuario do sistema
func (us *UserService) Logout(email, password string) (logout bool) {
	logout = us.UserRepository.Logout(email, password)
	return
}
