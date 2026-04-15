package infraestructure

import (
	"Plannex/src/eventos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewMyEventosController struct {
	useCase *application.ViewMyEventos
}

func NewViewMyEventosController(useCase *application.ViewMyEventos) *ViewMyEventosController {
	return &ViewMyEventosController{useCase: useCase}
}

func (vmp_c *ViewMyEventosController) Execute(c *gin.Context) {
	// Obtener el ID del usuario del contexto (se establece en el middleware de autenticación)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	userIDInt32, ok := userID.(int32)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el ID del usuario"})
		return
	}

	eventos, err := vmp_c.useCase.Execute(userIDInt32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"eventos": eventos})
}
