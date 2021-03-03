package usecases

import "github.com/FelipeAz/desafio-serasa/app/entity"

// NegativacaoService pertence a camada Usecases.
type NegativacaoService struct {
	NegativacaoRepository NegativacaoRepository
	CryptoHandler         CryptoHandler
}

// Get busca todas as negativacoes
func (ns *NegativacaoService) Get() (n []entity.Negativacao, err error) {
	n, err = ns.NegativacaoRepository.Get()
	if err != nil {
		return
	}
	if len(n) > 0 {
		for i := 0; i < len(n); i++ {
			n[i].CustomerDocument, err = ns.CryptoHandler.DecryptString(n[i].CustomerDocument)
			if err != nil {
				return nil, err
			}
		}
	}
	return
}

// GetByCPF busca uma negativacao com o CPF especificado.
func (ns *NegativacaoService) GetByCPF(cpf string) (n []entity.Negativacao, err error) {
	cryptedCPF, err := ns.CryptoHandler.EncryptString(cpf)
	if err != nil {
		return
	}
	n, err = ns.NegativacaoRepository.GetByCPF(cryptedCPF)
	if err != nil {
		return
	}
	if len(n) > 0 {
		for i := 0; i < len(n); i++ {
			n[i].CustomerDocument, err = ns.CryptoHandler.DecryptString(n[i].CustomerDocument)
			if err != nil {
				return
			}
		}
	}
	return
}

// Create insere uma negativacao no banco de dados.
func (ns *NegativacaoService) Create(n entity.Negativacao) (id uint, err error) {
	n.CustomerDocument, err = ns.CryptoHandler.EncryptString(n.CustomerDocument)
	if err != nil {
		return
	}
	id, err = ns.NegativacaoRepository.Create(n)
	return
}

// Update atualiza uma negativacao no banco de dados.
func (ns *NegativacaoService) Update(ID int, neg entity.Negativacao) (n entity.Negativacao, err error) {
	neg.CustomerDocument, err = ns.CryptoHandler.EncryptString(neg.CustomerDocument)
	if err != nil {
		return
	}
	n, err = ns.NegativacaoRepository.Update(ID, neg)
	if err != nil {
		return
	}
	n.CustomerDocument, err = ns.CryptoHandler.DecryptString(n.CustomerDocument)
	return
}

// Delete deleta uma negativacao do banco de dados.
func (ns *NegativacaoService) Delete(ID int) (err error) {
	err = ns.NegativacaoRepository.Delete(ID)
	return
}
