package entity

import (
	"time"
)

// Negativacao representa a estrutura de uma negativacao.
type Negativacao struct {
	ID               uint      `json:"id,omitempty" gorm:"primaryKey;autoIncrement;not null"`
	CompanyDocument  string    `json:"companyDocument" binding:"required"`
	CompanyName      string    `json:"companyName" gorm:"primaryKey" binding:"required"`
	CustomerDocument string    `json:"customerDocument" binding:"required"`
	Value            float64   `json:"value" binding:"required"`
	Contract         string    `json:"contract" gorm:"unique" binding:"required"`
	DebtDate         time.Time `json:"debtDate" binding:"required"`
	InclusionDate    time.Time `json:"inclusionDate" binding:"required"`
}

// TableName Sobrescreve o nome da tabela no banco de dados.
func (Negativacao) TableName() string {
	return "negativacoes"
}
