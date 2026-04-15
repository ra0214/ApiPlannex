package infraestructure

import (
	"Plannex/src/eventos/application"
	"Plannex/src/eventos/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEventosController struct {
	useCase *application.CreateEventos
	repo    domain.IEventos
}

func NewCreateEventosController(useCase *application.CreateEventos, repo domain.IEventos) *CreateEventosController {
	return &CreateEventosController{useCase: useCase, repo: repo}
}

type RequestBody struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
	QRCodeData  string   `json:"qr_code_data,omitempty"`
}

func (cp_c *CreateEventosController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	// Si el usuario está autenticado, obtener su ID del contexto
	var createdBy *int32
	if userID, exists := c.Get("userID"); exists {
		if userIDInt32, ok := userID.(int32); ok {
			createdBy = &userIDInt32
		}
	}

	id, err := cp_c.useCase.Execute(body.Title, body.Description, body.Date, body.Latitude, body.Longitude, body.QRCodeData, createdBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el evento", "detalles": err.Error()})
		return
	}

	evento, err := cp_c.repo.GetEventoById(id)
	if err == nil {
		GetHub().BroadcastEvent("create", id, evento)
		// Notificar por push notification (opcional: notificar a usuarios específicos)
		// NotifyEventCreated(id, evento.Title, evento.Description, []int32{})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Evento creado correctamente", "id": id, "evento": evento})
}
