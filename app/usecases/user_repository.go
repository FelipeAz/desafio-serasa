package usecases

import "github.com/FelipeAz/desafio-serasa/app/entity"

// UserRepository pertence a camada de usecases.
type UserRepository interface {
	Login(string, string) (entity.User, error)
	AuthUser(uint, string) entity.Access
	SignUp(*entity.User) *entity.User
	Logout(string, string) bool
}
