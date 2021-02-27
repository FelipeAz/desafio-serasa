package interfaces

import (
	"github.com/FelipeAz/desafio-serasa/app/usecases"
	"github.com/gin-gonic/gin"
)

// NegativacaoController pertence a camada de interface.
type NegativacaoController struct {
	NegativacaoService usecases.NegativacaoService
}

// GetAll retorna todas as negativacoes.
func (nc *NegativacaoController) GetAll(c *gin.Context) {

}

// Get busca uma negativacao com o ID especificado.
func (nc *NegativacaoController) Get(c *gin.Context) {

}

// Save insere uma negativacao no banco de dados.
func (nc *NegativacaoController) Save(c *gin.Context) {

}

// Change atualiza uma negativacao no banco de dados.
func (nc *NegativacaoController) Change(c *gin.Context) {

}

// Destroy deleta uma negativacao do banco de dados.
func (nc *NegativacaoController) Destroy(c *gin.Context) {

}
