package interfaces

import (
	"net/http"

	"github.com/FelipeAz/desafio-serasa/app/usecases"
	"github.com/gin-gonic/gin"
)

// MainframeController representa uma instancia do controller.
type MainframeController struct {
	MainframeService usecases.MainframeService
}

// NewMainframeController retorna uma instancia do controller do Mainframe.
func NewMainframeController(sqlHandler SQLHandler) *MainframeController {
	return &MainframeController{
		MainframeService: usecases.MainframeService{
			NegativacaoRepository: &NegativacaoRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// Integrate retorna todas as Negativacoes do servidor JSON
func (mc *MainframeController) Integrate(c *gin.Context) {
	err := mc.MainframeService.Integrate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, "Integration Finished")
}
