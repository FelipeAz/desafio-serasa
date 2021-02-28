package usecases

import (
	"encoding/json"
	"io/ioutil"

	"github.com/FelipeAz/desafio-serasa/app/entity"
)

// MainframeService representa uma instancia do servico relacionado ao mainframe.
type MainframeService struct {
	NegativacaoRepository NegativacaoRepository
}

func (ms *MainframeService) connectMainframe() error {
	return nil
}

// Integrate realiza a integracao com o mainframe persistindo as negativacoes no BD.
func (ms *MainframeService) Integrate() error {
	var data []entity.Negativacao
	if err := ms.connectMainframe(); err != nil {
		return err
	}

	err := ms.readJSONFile(&data, "app/negativacoes.json")
	if err != nil {
		return err
	}

	// Persist Negativacoes

	return nil
}

func (ms *MainframeService) readJSONFile(data *[]entity.Negativacao, fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MainframeService) persistNegativacao(negativacao entity.Negativacao) (id uint, err error) {
	id, err = ms.NegativacaoRepository.Create(negativacao)
	return
}
