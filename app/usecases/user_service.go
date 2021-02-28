package usecases

import "github.com/FelipeAz/desafio-serasa/app/entity"

// UserService pertence a camada usecases.
type UserService struct {
	UserRepository UserRepository
}

// Login do usuario no sistema.
func (us *UserService) Login(email, password string) (usr entity.User, err error) {
	usr, err = us.UserRepository.Login(email, password)
	return
}

// AuthUser insere o token de autenticacao no banco.
func (us *UserService) AuthUser(id uint, token string) (access entity.Access) {
	access = us.UserRepository.AuthUser(id, token)
	return
}

// SignUp registra um novo usuario ao sistema.
func (us *UserService) SignUp(usr *entity.User) (user *entity.User) {
	user = us.UserRepository.SignUp(usr)
	return
}

// Logout desloga o usuario do sistema
func (us *UserService) Logout(email, password string) (logout bool) {
	logout = us.UserRepository.Logout(email, password)
	return
}
