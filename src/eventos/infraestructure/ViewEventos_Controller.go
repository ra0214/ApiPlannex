package infraestructure

import (
	"Plannex/src/eventos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewEventosController struct {
	useCase *application.ViewEventos
}

func NewViewEventosController(useCase *application.ViewEventos) *ViewEventosController {
	return &ViewEventosController{useCase: useCase}
}

func (vp_c *ViewEventosController) Execute(c *gin.Context) {
	eventos, err := vp_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"eventos": eventos})
}
