package usecases

import "github.com/FelipeAz/desafio-serasa/app/entity"

// NegativacaoRepository pertence a camada de usecases.
type NegativacaoRepository interface {
	Find() []entity.Negativacao
	FindByID(int) (entity.Negativacao, error)
	Create(entity.Negativacao) uint
	Update(int, *entity.Negativacao) (*entity.Negativacao, error)
	Delete(int) error
}
