package entity

import (
	"time"
)

// Negativacao representa a estrutura de uma negativacao.
type Negativacao struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	CompanyDocument  string    `json:"companyDocument"`
	CompanyName      string    `json:"companyName"`
	CustomerDocument string    `json:"customerDocument"`
	Value            float64   `json:"value"`
	DebtDate         time.Time `json:"debtDate"`
	InclusionDate    time.Time `json:"inclusionDate"`
}

// TableName Sobrescreve o nome da tabela no banco de dados.
func (Negativacao) TableName() string {
	return "negativacoes"
}
