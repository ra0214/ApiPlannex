package infraestructure

import (
	"net/http"
	"strconv"

	"Plannex/src/eventos/domain"

	"github.com/gin-gonic/gin"
)

type InviteUserController struct {
	repo domain.IEventos
}

func NewInviteUserController(repo domain.IEventos) *InviteUserController {
	return &InviteUserController{repo: repo}
}

type InviteUserRequest struct {
	UserID int32 `json:"user_id"`
}

func (controller *InviteUserController) Execute(c *gin.Context) {
	eventoIdStr := c.Param("id")
	eventoId, err := strconv.ParseInt(eventoIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de evento inválido"})
		return
	}

	var body InviteUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	if body.UserID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id debe ser mayor a 0"})
		return
	}

	err = controller.repo.InviteUser(int32(eventoId), body.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al invitar usuario", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario invitado correctamente"})
}
