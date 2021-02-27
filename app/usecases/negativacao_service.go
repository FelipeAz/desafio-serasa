package usecases

import "github.com/FelipeAz/desafio-serasa/app/entity"

// NegativacaoService pertence a camada Usecases.
type NegativacaoService struct {
	NegativacaoRepository NegativacaoRepository
}

// GetAll busca todas as negativacoes
func (ns *NegativacaoService) GetAll() (n entity.Negativacoes, err error) {
	n, err = ns.NegativacaoRepository.Find()

	return
}

// Get busca uma negativacao com o ID especificado.
func (ns *NegativacaoService) Get(ID int) (n entity.Negativacao, err error) {
	n, err = ns.NegativacaoRepository.FindById(ID)

	return
}

// Save insere uma negativacao no banco de dados.
func (ns *NegativacaoService) Save(n entity.Negativacao) (id int64, err error) {
	id, err = ns.NegativacaoRepository.Create(n)

	return
}

// Change atualiza uma negativacao no banco de dados.
func (ns *NegativacaoService) Change(ID int, neg entity.Negativacao) (n entity.Negativacao, err error) {
	n, err = ns.NegativacaoRepository.Update(ID, neg)

	return
}

// Destroy deleta uma negativacao do banco de dados.
func (ns *NegativacaoService) Destroy(ID int) (err error) {
	err = ns.NegativacaoRepository.Delete(ID)

	return
}
