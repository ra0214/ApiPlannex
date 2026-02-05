package infraestructure

import (
	"Plannex/src/users/application"
	"Plannex/src/users/domain"

	"github.com/gin-gonic/gin"
)

// InitRouter registra las rutas de users en un router existente
func InitRouter(r *gin.Engine, repo domain.IUser) {
	setupUserRoutes(r, repo)
}

// setupUserRoutes registra todas las rutas de usuarios
func setupUserRoutes(r *gin.Engine, repo domain.IUser) {
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
	loginUserController := NewLoginUserController(loginUser)

	r.POST("/user", createUserController.Execute)
	r.GET("/user", viewUserController.Execute)
	r.PUT("/user/:id", editUserController.Execute)
	r.DELETE("/user/:id", deleteUserController.Execute)
	r.POST("/login", loginUserController.Execute)
}
