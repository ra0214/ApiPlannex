package infraestructure

import (
	"Plannex/src/eventos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditEventosController struct {
	useCase *application.EditEventos
}

func NewEditEventosController(useCase *application.EditEventos) *EditEventosController {
	return &EditEventosController{useCase: useCase}
}

func (ep_c *EditEventosController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de evento inválido"})
		return
	}

	var p struct {
		Nombre string `json:"nombre"`
		Fecha  string `json:"fecha"`
	}

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer los datos"})
		return
	}

	err = ep_c.useCase.Execute(int32(id), p.Nombre, p.Fecha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el evento", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Evento actualizado correctamente"})
}
