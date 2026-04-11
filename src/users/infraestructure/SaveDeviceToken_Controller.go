package infraestructure

import (
	"Plannex/src/users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SaveDeviceTokenController struct {
	useCase *application.SaveDeviceToken
}

func NewSaveDeviceTokenController(useCase *application.SaveDeviceToken) *SaveDeviceTokenController {
	return &SaveDeviceTokenController{useCase: useCase}
}

type SaveDeviceTokenRequest struct {
	FCMToken   string `json:"fcm_token" binding:"required"`
	DeviceName string `json:"device_name"`
}

func (sdt *SaveDeviceTokenController) Execute(c *gin.Context) {
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	var body SaveDeviceTokenRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	err = sdt.useCase.Execute(int32(userID), body.FCMToken, body.DeviceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar token de dispositivo", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Token de dispositivo registrado correctamente"})
}
