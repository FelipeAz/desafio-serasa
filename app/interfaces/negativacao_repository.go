package interfaces

import (
	"encoding/json"
	"strconv"

	"github.com/FelipeAz/desafio-serasa/app/entity"
)

// NegativacaoRepository eh responsavel por toda operacao que envolve banco.
type NegativacaoRepository struct {
	SQLHandler SQLHandler
	Redis      Redis
}

// Get retorna todas as negativacoes.
func (nr *NegativacaoRepository) Get() []entity.Negativacao {
	var negativacoes []entity.Negativacao

	redisName := "all"
	data, err := nr.Redis.Get(redisName)
	if err != nil {
		db := nr.SQLHandler.GetGorm()
		err = db.Debug().Model(&entity.Negativacao{}).Find(&negativacoes).Error
		if err != nil {
			return []entity.Negativacao{}
		}

		newData, _ := json.Marshal(negativacoes)
		nr.Redis.Set(redisName, newData)
	}

	json.Unmarshal(data, &negativacoes)

	return negativacoes
}

// GetByID retorna uma unica negativacao.
func (nr *NegativacaoRepository) GetByID(ID int) (entity.Negativacao, error) {
	var negativacao entity.Negativacao

	data, err := nr.Redis.Get(strconv.Itoa(ID))
	if err != nil {
		db := nr.SQLHandler.GetGorm()
		if err := db.Where("id = ?", ID).First(&negativacao).Error; err != nil {
			return negativacao, err
		}

		newData, _ := json.Marshal(negativacao)
		nr.Redis.Set(strconv.Itoa(ID), newData)
	}

	json.Unmarshal(data, &negativacao)

	return negativacao, nil
}

// Create cria uma negativacao.
func (nr *NegativacaoRepository) Create(neg entity.Negativacao) (uint, error) {
	db := nr.SQLHandler.GetGorm()
	result := db.Create(&neg)

	newData, _ := json.Marshal(neg)
	nr.Redis.Set(strconv.Itoa(int(neg.ID)), newData)
	nr.Redis.Flush("all")

	return neg.ID, result.Error
}

// Update atualiza uma negativacao.
func (nr *NegativacaoRepository) Update(ID int, input entity.Negativacao) (entity.Negativacao, error) {
	db := nr.SQLHandler.GetGorm()
	neg, err := nr.GetByID(ID)

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

	newData, _ := json.Marshal(neg)
	nr.Redis.Set(strconv.Itoa(int(neg.ID)), newData)
	nr.Redis.Flush("all")

	return neg, nil
}

// Delete deleta uma negativacao.
func (nr *NegativacaoRepository) Delete(ID int) error {
	db := nr.SQLHandler.GetGorm()
	n, err := nr.GetByID(ID)

	if err != nil {
		return err
	}

	db.Delete(&n)

	nr.Redis.Flush(strconv.Itoa(ID))
	nr.Redis.Flush("all")

	return nil
}
