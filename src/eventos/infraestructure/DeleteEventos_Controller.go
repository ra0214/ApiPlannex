package infraestructure

import (
	"Plannex/src/eventos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteEventosController struct {
	useCase *application.DeleteEventos
}

func NewDeleteEventosController(useCase *application.DeleteEventos) *DeleteEventosController {
	return &DeleteEventosController{useCase: useCase}
}

func (dp_c *DeleteEventosController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de evento inválido"})
		return
	}

	eventoID := int32(id)
	err = dp_c.useCase.Execute(eventoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el evento", "detalles": err.Error()})
		return
	}

	GetHub().BroadcastEvent("delete", eventoID, gin.H{"id": eventoID})

	c.JSON(http.StatusOK, gin.H{"message": "Evento eliminado correctamente"})
}
