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

// Get retorna todas as negativacoes.
func (nc *NegativacaoController) Get(c *gin.Context) {
	n := nc.NegativacaoService.Get()

	c.JSON(http.StatusOK, gin.H{"data": n})
}

// GetByID busca uma negativacao com o ID especificado.
func (nc *NegativacaoController) GetByID(c *gin.Context) {
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

// Create insere uma negativacao no banco de dados.
func (nc *NegativacaoController) Create(c *gin.Context) {
	var input entity.Negativacao
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := nc.NegativacaoService.Persist(input)

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

	neg, err := nc.NegativacaoService.Update(id, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": neg})
}

// Delete deleta uma negativacao do banco de dados.
func (nc *NegativacaoController) Delete(c *gin.Context) {
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
