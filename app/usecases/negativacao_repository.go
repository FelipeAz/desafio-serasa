package usecases

import "github.com/FelipeAz/desafio-serasa/app/entity"

// NegativacaoRepository pertence a camada de usecases.
type NegativacaoRepository interface {
	Find() (entity.Negativacoes, error)
	FindById(int) (entity.Negativacao, error)
	Create(entity.Negativacao) (int64, error)
	Update(int, entity.Negativacao) (entity.Negativacao, error)
	Delete(int) error
}
