package usecases

import "github.com/FelipeAz/desafio-serasa/app/entity"

// NegativacaoRepository pertence a camada de usecases.
type NegativacaoRepository interface {
	Get() []entity.Negativacao
	GetByID(int) (entity.Negativacao, error)
	Create(entity.Negativacao) (uint, error)
	Update(int, entity.Negativacao) (entity.Negativacao, error)
	Delete(int) error
}
