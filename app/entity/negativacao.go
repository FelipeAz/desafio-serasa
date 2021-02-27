package entity

import (
	"time"
)

// Negativacao representa a estrutura de uma negativacao.
type Negativacao struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	CompanyDocument  string    `json:"companyDocument" binding:"required"`
	CompanyName      string    `json:"companyName" binding:"required"`
	CustomerDocument string    `json:"customerDocument" binding:"required"`
	Value            float64   `json:"value" binding:"required"`
	DebtDate         time.Time `json:"debtDate" binding:"required"`
	InclusionDate    time.Time `json:"inclusionDate" binding:"required"`
}

// TableName Sobrescreve o nome da tabela no banco de dados.
func (Negativacao) TableName() string {
	return "negativacoes"
}
