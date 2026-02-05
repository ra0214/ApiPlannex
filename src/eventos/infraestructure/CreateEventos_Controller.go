package infraestructure

import (
	"Plannex/src/eventos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEventosController struct {
	useCase *application.CreateEventos
}

func NewCreateEventosController(useCase *application.CreateEventos) *CreateEventosController {
	return &CreateEventosController{useCase: useCase}
}

type RequestBody struct {
	Nombre string `json:"nombre"`
	Fecha  string `json:"fecha"`
}

func (cp_c *CreateEventosController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	err := cp_c.useCase.Execute(body.Nombre, body.Fecha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el evento", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Evento creado correctamente"})
}
