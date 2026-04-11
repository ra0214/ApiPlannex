package infraestructure

import (
	"Plannex/src/users/application"
	"Plannex/src/users/domain"

	"github.com/gin-gonic/gin"
)

// InitRouter registra las rutas de users en un router existente
func InitRouter(r *gin.Engine, repo domain.IUser) {
	deviceTokenRepo := NewMySQLDeviceToken()
	setupUserRoutes(r, repo, deviceTokenRepo)
}

// setupUserRoutes registra todas las rutas de usuarios
func setupUserRoutes(r *gin.Engine, repo domain.IUser, deviceTokenRepo domain.IDeviceToken) {
	// Mismo contenido que SetupRouter pero usando el router existente
	createUser := application.NewCreateUser(repo)
	createUserController := NewCreateUserController(createUser)

	viewUser := application.NewViewUser(repo)
	viewUserController := NewViewUserController(viewUser)

	editUserUseCase := application.NewEditUser(repo)
	editUserController := NewEditUserController(editUserUseCase)

	deleteUserUseCase := application.NewDeleteUser(repo)
	deleteUserController := NewDeleteUserController(deleteUserUseCase)

	loginUser := application.NewLoginUser(repo)
	loginUserController := NewLoginUserController(loginUser, repo)

	saveDeviceTokenUseCase := application.NewSaveDeviceToken(deviceTokenRepo)
	saveDeviceTokenController := NewSaveDeviceTokenController(saveDeviceTokenUseCase)

	deleteDeviceTokenUseCase := application.NewDeleteDeviceToken(deviceTokenRepo)
	deleteDeviceTokenController := NewDeleteDeviceTokenController(deleteDeviceTokenUseCase)

	r.POST("/user", createUserController.Execute)
	r.GET("/user", viewUserController.Execute)
	r.PUT("/user/:id", editUserController.Execute)
	r.DELETE("/user/:id", deleteUserController.Execute)
	r.POST("/login", loginUserController.Execute)

	// Device Token Routes
	r.POST("/user/:userId/fcm-token", saveDeviceTokenController.Execute)
	r.DELETE("/user/:userId/fcm-token", deleteDeviceTokenController.Execute)
}
