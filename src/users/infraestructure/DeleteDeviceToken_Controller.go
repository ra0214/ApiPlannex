package infraestructure

import (
	"Plannex/src/users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteDeviceTokenController struct {
	useCase *application.DeleteDeviceToken
}

func NewDeleteDeviceTokenController(useCase *application.DeleteDeviceToken) *DeleteDeviceTokenController {
	return &DeleteDeviceTokenController{useCase: useCase}
}

type DeleteDeviceTokenRequest struct {
	FCMToken string `json:"fcm_token" binding:"required"`
}

func (ddt *DeleteDeviceTokenController) Execute(c *gin.Context) {
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	var body DeleteDeviceTokenRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	err = ddt.useCase.Execute(int32(userID), body.FCMToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar token de dispositivo", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token de dispositivo eliminado correctamente"})
}
