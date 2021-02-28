package usecases

import "github.com/FelipeAz/desafio-serasa/app/entity"

// NegativacaoService pertence a camada Usecases.
type NegativacaoService struct {
	NegativacaoRepository NegativacaoRepository
}

// Get busca todas as negativacoes
func (ns *NegativacaoService) Get() (n []entity.Negativacao) {
	n = ns.NegativacaoRepository.Find()
	return
}

// GetByID busca uma negativacao com o ID especificado.
func (ns *NegativacaoService) GetByID(ID int) (n entity.Negativacao, err error) {
	n, err = ns.NegativacaoRepository.FindByID(ID)
	return
}

// Persist insere uma negativacao no banco de dados.
func (ns *NegativacaoService) Persist(n entity.Negativacao) (id uint) {
	id = ns.NegativacaoRepository.Create(n)
	return
}

// Update atualiza uma negativacao no banco de dados.
func (ns *NegativacaoService) Update(ID int, neg *entity.Negativacao) (n *entity.Negativacao, err error) {
	n, err = ns.NegativacaoRepository.Update(ID, neg)
	return
}

// Destroy deleta uma negativacao do banco de dados.
func (ns *NegativacaoService) Destroy(ID int) (err error) {
	err = ns.NegativacaoRepository.Delete(ID)
	return
}
