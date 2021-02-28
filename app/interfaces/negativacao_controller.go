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

// Find retorna todas as negativacoes.
func (nc *NegativacaoController) Find(c *gin.Context) {
	n := nc.NegativacaoService.Get()

	c.JSON(http.StatusOK, gin.H{"data": n})
}

// FindByID busca uma negativacao com o ID especificado.
func (nc *NegativacaoController) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID."})
		return
	}

	n, err := nc.NegativacaoService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"data": n})
}

// Persist insere uma negativacao no banco de dados.
func (nc *NegativacaoController) Persist(c *gin.Context) {
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

	id := nc.NegativacaoService.Persist(negativacao)

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// Update atualiza uma negativacao no banco de dados.
func (nc *NegativacaoController) Update(c *gin.Context) {
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
		InclusionDate:    input.InclusionDate,
	}

	neg, err := nc.NegativacaoService.Update(id, &negativacao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
