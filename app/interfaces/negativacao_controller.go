package interfaces

import (
	"net/http"
	"strconv"

	"github.com/FelipeAz/desafio-serasa/app/entity"
	"github.com/FelipeAz/desafio-serasa/app/usecases"
	"github.com/gin-gonic/gin"
)

// NegativacaoController redireciona a requisicao HTTP ao servico.
type NegativacaoController struct {
	NegativacaoService usecases.NegativacaoService
}

// NewNegativacaoController retorna uma instancia do controller.
func NewNegativacaoController(sqlHandler SQLHandler) *NegativacaoController {
	return &NegativacaoController{
		NegativacaoService: usecases.NegativacaoService{
			NegativacaoRepository: &NegativacaoRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// GetAll retorna todas as negativacoes.
func (nc *NegativacaoController) GetAll(c *gin.Context) {
	n := nc.NegativacaoService.GetAll()

	c.JSON(http.StatusOK, gin.H{"data": n})
}

// Get busca uma negativacao com o ID especificado.
func (nc *NegativacaoController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID."})
		return
	}

	n, err := nc.NegativacaoService.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"data": n})
}

// Save insere uma negativacao no banco de dados.
func (nc *NegativacaoController) Save(c *gin.Context) {
	var input entity.Negativacao
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	negativacao := entity.Negativacao{
		CompanyDocument:  input.CompanyDocument,
		CompanyName:      input.CompanyName,
		CustomerDocument: input.CustomerDocument,
		Value:            input.Value,
		DebtDate:         input.DebtDate,
		InclusionDate:    input.DebtDate,
	}

	id := nc.NegativacaoService.Save(negativacao)

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// Change atualiza uma negativacao no banco de dados.
func (nc *NegativacaoController) Change(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID."})
		return
	}

	var input entity.Negativacao
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	negativacao := entity.Negativacao{
		CompanyDocument:  input.CompanyDocument,
		CompanyName:      input.CompanyName,
		CustomerDocument: input.CustomerDocument,
		Value:            input.Value,
		DebtDate:         input.DebtDate,
		InclusionDate:    input.DebtDate,
	}

	neg, err := nc.NegativacaoService.Change(id, negativacao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": neg})
}

// Destroy deleta uma negativacao do banco de dados.
func (nc *NegativacaoController) Destroy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID."})
		return
	}

	err = nc.NegativacaoService.Destroy(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
