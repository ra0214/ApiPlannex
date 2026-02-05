package main

import (
	"log"

	eventosInfra "Plannex/src/eventos/infraestructure"
	usersInfra "Plannex/src/users/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crear un único router
	r := gin.Default()

	// Inicializar users y eventos pasando el mismo router
	usersInfra.InitRouter(r, usersInfra.NewMySQL())
	eventosInfra.Init(r)

	// Ejecutar servidor en :8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
