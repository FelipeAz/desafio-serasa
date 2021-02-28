package interfaces

import (
	"github.com/FelipeAz/desafio-serasa/app/entity"
)

// UserRepository representa o repositorio dos usuarios e realiza as operacoes de BD.
type UserRepository struct {
	SQLHandler SQLHandler
}

// Login retorna um usuario do banco de dados.
func (ur *UserRepository) Login(email string, password string) (entity.User, error) {
	var usr entity.User
	db := ur.SQLHandler.GetGorm()
	if err := db.Where("email=? AND password=?", email, password).First(&usr).Error; err != nil {
		return usr, err
	}

	return usr, nil
}

// AuthUser insere o token de autenticacao no banco.
func (ur *UserRepository) AuthUser(id uint, token string) entity.Access {
	db := ur.SQLHandler.GetGorm()

	// Atualiza o token de acesso caso ja exista o ID do usuario na tabela access
	var refreshAuth entity.Access
	if err := db.Where("user_id=?", id).First(&refreshAuth).Error; err == nil {
		db.Model(&refreshAuth).Update("access_token", token)
		return refreshAuth
	}

	auth := entity.Access{
		UserID:      id,
		AccessToken: token,
	}

	db.Create(&auth)

	return auth
}

// SignUp registra um novo usuario ao sistema.
func (ur *UserRepository) SignUp(usr *entity.User) *entity.User {
	db := ur.SQLHandler.GetGorm()
	db.Create(&usr)

	return usr
}

// Logout desloga o usuario do sistema
func (ur *UserRepository) Logout(email, password string) bool {
	db := ur.SQLHandler.GetGorm()

	var usr entity.User
	if err := db.Where("email=? AND password=?", email, password).First(&usr).Error; err != nil {
		return false
	}

	var access entity.Access
	if err := db.Where("user_id = ?", usr.ID).First(&access).Error; err != nil {
		return false
	}

	db.Delete(&access)
	return true
}
