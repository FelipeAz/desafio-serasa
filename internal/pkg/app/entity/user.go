package entity

// User representa o usuario.
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
