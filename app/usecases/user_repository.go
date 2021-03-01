package usecases

import "github.com/FelipeAz/desafio-serasa/app/entity"

// UserRepository pertence a camada de usecases.
type UserRepository interface {
	Login(string, string) (entity.User, error)
	AuthUser(uint, *entity.TokenDetails) (entity.Access, error)
	SignUp(*entity.User) (*entity.User, error)
	Logout(string, string) bool
}
