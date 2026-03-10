package infraestructure

import (
	"net/http"
	"strconv"

	"Plannex/src/eventos/domain"

	"github.com/gin-gonic/gin"
)

type ConfirmAttendanceController struct {
	repo domain.IEventos
}

func NewConfirmAttendanceController(repo domain.IEventos) *ConfirmAttendanceController {
	return &ConfirmAttendanceController{repo: repo}
}

type ConfirmAttendanceRequest struct {
	UserID int32  `json:"user_id"`
	Estado string `json:"estado"`
}

func (controller *ConfirmAttendanceController) Execute(c *gin.Context) {
	eventoIdStr := c.Param("id")
	eventoId, err := strconv.ParseInt(eventoIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de evento inválido"})
		return
	}

	var body ConfirmAttendanceRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	if body.UserID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id debe ser mayor a 0"})
		return
	}

	validStates := map[string]bool{
		"asistira":    true,
		"quiza":       true,
		"no_asistira": true,
		"invitado":    true,
	}

	if !validStates[body.Estado] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estado inválido. Valores permitidos: asistira, quiza, no_asistira, invitado"})
		return
	}

	err = controller.repo.ConfirmAttendance(int32(eventoId), body.UserID, body.Estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al confirmar asistencia", "detalles": err.Error()})
		return
	}

	// Notificar por WebSocket: alguien confirmó asistencia al evento
	GetHub().BroadcastEvent("attendance", int32(eventoId), gin.H{
		"evento_id": eventoId,
		"user_id":   body.UserID,
		"estado":    body.Estado,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Asistencia confirmada correctamente"})
}
