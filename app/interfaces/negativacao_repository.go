package interfaces

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/FelipeAz/desafio-serasa/app/entity"
)

// Clean Architecture: A camada interface eh responsavel por transformar data em entidade, portanto, como o repository
// realiza essa funcao, nao faz sentido implementar aqui uma interface, mas sim a execucao das operacoes
// no banco de dados.

// NegativacaoRepository eh responsavel por toda operacao que envolve banco.
type NegativacaoRepository struct {
	SQLHandler SQLHandler
	Redis      Redis
}

// Get retorna todas as negativacoes.
func (nr *NegativacaoRepository) Get() ([]entity.Negativacao, error) {
	var negativacoes []entity.Negativacao

	redisName := "all"
	data, err := nr.Redis.Get(redisName)
	if err != nil {
		db := nr.SQLHandler.GetGorm()
		err = db.Debug().Model(&entity.Negativacao{}).Find(&negativacoes).Error
		if err != nil {
			return []entity.Negativacao{}, err
		}

		newData, _ := json.Marshal(negativacoes)
		err := nr.Redis.Set(redisName, newData)
		if err != nil {
			log.Println(err)
			return []entity.Negativacao{}, err
		}

		return negativacoes, nil
	}

	err = json.Unmarshal(data, &negativacoes)
	if err != nil {
		log.Println(err)
		return []entity.Negativacao{}, err
	}

	return negativacoes, nil
}

// GetByID retorna uma unica negativacao.
func (nr *NegativacaoRepository) GetByID(ID int) (entity.Negativacao, error) {
	var negativacao entity.Negativacao

	data, err := nr.Redis.Get(strconv.Itoa(ID))
	if err != nil {
		db := nr.SQLHandler.GetGorm()
		if err := db.Where("id = ?", ID).First(&negativacao).Error; err != nil {
			return entity.Negativacao{}, err
		}

		newData, _ := json.Marshal(negativacao)
		err := nr.Redis.Set(strconv.Itoa(int(negativacao.ID)), newData)
		if err != nil {
			log.Println(err)
			return entity.Negativacao{}, err
		}

		return negativacao, nil
	}

	err = json.Unmarshal(data, &negativacao)
	if err != nil {
		log.Println(err)
		return entity.Negativacao{}, err
	}

	return negativacao, nil
}

// GetByCPF retorna uma unica negativacao.
func (nr *NegativacaoRepository) GetByCPF(cryptedCPF string) ([]entity.Negativacao, error) {
	var negativacoes []entity.Negativacao

	db := nr.SQLHandler.GetGorm()

	err := db.Debug().Model(&entity.Negativacao{}).Where("customer_document=?", cryptedCPF).Find(&negativacoes).Error
	if err != nil {
		return []entity.Negativacao{}, err
	}

	return negativacoes, nil
}

// Create cria uma negativacao.
func (nr *NegativacaoRepository) Create(neg entity.Negativacao) (uint, error) {
	db := nr.SQLHandler.GetGorm()
	result := db.Create(&neg)

	newData, _ := json.Marshal(neg)
	err := nr.Redis.Set(strconv.Itoa(int(neg.ID)), newData)
	if err != nil {
		log.Println(err)
	}
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
		"Contract":         input.Contract,
		"DebtDate":         input.DebtDate,
		"InclusionDate":    input.InclusionDate,
	}).Error; err != nil {
		return entity.Negativacao{}, err
	}

	newData, _ := json.Marshal(neg)
	err = nr.Redis.Set(strconv.Itoa(int(neg.ID)), newData)
	if err != nil {
		log.Println(err)
	}
	nr.Redis.Flush("all")

	return neg, nil
}

// Delete deleta uma negativacao.
func (nr *NegativacaoRepository) Delete(ID int) error {
	db := nr.SQLHandler.GetGorm()
	n, err := nr.GetByID(ID)

	if err != nil {
		log.Println(err)
		return err
	}

	db.Delete(&n)

	nr.Redis.Flush(strconv.Itoa(ID))
	nr.Redis.Flush("all")

	return nil
}
