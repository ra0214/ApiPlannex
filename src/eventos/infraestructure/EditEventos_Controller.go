package infraestructure

import (
	"Plannex/src/eventos/application"
	"Plannex/src/eventos/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditEventosController struct {
	useCase *application.EditEventos
	repo    domain.IEventos
}

func NewEditEventosController(useCase *application.EditEventos, repo domain.IEventos) *EditEventosController {
	return &EditEventosController{useCase: useCase, repo: repo}
}

func (ep_c *EditEventosController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de evento inválido"})
		return
	}

	var body struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Date        string   `json:"date"`
		Latitude    *float64 `json:"latitude,omitempty"`
		Longitude   *float64 `json:"longitude,omitempty"`
		QRCodeData  string   `json:"qr_code_data,omitempty"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer los datos"})
		return
	}

	err = ep_c.useCase.Execute(int32(id), body.Title, body.Description, body.Date, body.Latitude, body.Longitude, body.QRCodeData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el evento", "detalles": err.Error()})
		return
	}

	evento, err := ep_c.repo.GetEventoById(int32(id))
	if err == nil {
		GetHub().BroadcastEvent("update", int32(id), evento)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Evento actualizado correctamente", "evento": evento})
}
