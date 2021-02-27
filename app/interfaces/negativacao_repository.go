package interfaces

import (
	"github.com/FelipeAz/desafio-serasa/app/entity"
)

// NegativacaoRepository eh responsavel por toda operacao que envolve banco.
type NegativacaoRepository struct {
	SQLHandler SQLHandler
}

// Find retorna todas as negativacoes.
func (nr *NegativacaoRepository) Find() []entity.Negativacao {
	var negativacoes []entity.Negativacao
	db := nr.SQLHandler.GetGorm()
	db.Find(&negativacoes)

	return negativacoes
}

// FindByID retorna uma unica negativacao.
func (nr *NegativacaoRepository) FindByID(ID int) (entity.Negativacao, error) {
	var negativacao entity.Negativacao
	db := nr.SQLHandler.GetGorm()
	if err := db.Where("id = ?", ID).First(&negativacao).Error; err != nil {
		return negativacao, err
	}

	return negativacao, nil
}

// Create cria uma negativacao.
func (nr *NegativacaoRepository) Create(neg entity.Negativacao) uint {
	db := nr.SQLHandler.GetGorm()
	db.Create(&neg)

	return neg.ID
}

// Update atualiza uma negativacao.
func (nr *NegativacaoRepository) Update(ID int, input entity.Negativacao) (entity.Negativacao, error) {
	db := nr.SQLHandler.GetGorm()
	neg, err := nr.FindByID(ID)

	if err != nil {
		return neg, err
	}

	if err := db.Model(&neg).Updates(map[string]interface{}{
		"CompanyDocument":  input.CompanyDocument,
		"CompanyName":      input.CompanyName,
		"CustomerDocument": input.CustomerDocument,
		"Value":            input.Value,
		"DebtDate":         input.DebtDate,
		"InclusionDate":    input.InclusionDate,
	}).Error; err != nil {
		return neg, err
	}

	return neg, nil
}

// Delete deleta uma negativacao.
func (nr *NegativacaoRepository) Delete(ID int) error {
	db := nr.SQLHandler.GetGorm()
	n, err := nr.FindByID(ID)

	if err != nil {
		return err
	}

	db.Delete(&n)

	return nil
}
