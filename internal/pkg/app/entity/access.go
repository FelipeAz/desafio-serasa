package entity

// Access representa o acesso da autenticacao do usuario
type Access struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"userId" binding:"required"`
	AccessToken  string `json:"token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// TableName Sobrescreve o nome da tabela no banco de dados.
func (a Access) TableName() string {
	return "access"
}
