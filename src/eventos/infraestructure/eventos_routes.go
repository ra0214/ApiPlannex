package infraestructure

import (
	"Plannex/src/eventos/application"
	"Plannex/src/eventos/domain"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IEventos, r *gin.Engine) {
	createEventos := application.NewCreateEventos(repo)
	createEventosController := NewCreateEventosController(createEventos)

	viewEventos := application.NewViewEventos(repo)
	viewEventosController := NewViewEventosController(viewEventos)

	editEventosUseCase := application.NewEditEventos(repo)
	editEventosController := NewEditEventosController(editEventosUseCase)

	deleteEventosUseCase := application.NewDeleteEventos(repo)
	deleteEventosController := NewDeleteEventosController(deleteEventosUseCase)

	getEventosByIdUseCase := application.NewGetEventosById(repo)
	getEventosByIdController := NewGetEventoByIdController(getEventosByIdUseCase)

	inviteUserController := NewInviteUserController(repo)
	confirmAttendanceController := NewConfirmAttendanceController(repo)

	r.POST("/eventos", createEventosController.Execute)
	r.GET("/eventos", viewEventosController.Execute)
	r.PUT("/eventos/:id", editEventosController.Execute)
	r.DELETE("/eventos/:id", deleteEventosController.Execute)
	r.GET("/eventos/:id", getEventosByIdController.Execute)
	r.POST("/eventos/:id/invitar", inviteUserController.Execute)
	r.PUT("/eventos/:id/asistencia", confirmAttendanceController.Execute)
}
