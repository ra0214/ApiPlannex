package infraestructure

import (
	"Plannex/src/eventos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetEventoByIdController struct {
	useCase *application.GetEventosById
}

func NewGetEventoByIdController(useCase *application.GetEventosById) *GetEventoByIdController {
	return &GetEventoByIdController{useCase: useCase}
}

func (gp_c *GetEventoByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de evento inválido"})
		return
	}

	evento, err := gp_c.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el evento", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, evento)
}
