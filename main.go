package main

import (
	"log"
	"os"

	"Plannex/src/config"
	eventosInfra "Plannex/src/eventos/infraestructure"
	usersInfra "Plannex/src/users/infraestructure"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	godotenv.Load(".env")

	// Inicializar Firebase
	credentialsPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if credentialsPath == "" {
		credentialsPath = "./serviceAccountKey.json"
	}

	if err := config.InitFirebase(credentialsPath); err != nil {
		log.Printf("Advertencia: Firebase no se inicializó: %v (notificaciones no funcionarán)", err)
	} else {
		log.Println("Firebase inicializado exitosamente")
	}

	// Crear un único router
	r := gin.Default()

	// Inicializar users y eventos pasando el mismo router
	usersInfra.InitRouter(r, usersInfra.NewMySQL())
	eventosInfra.Init(r)

	// Obtener puerto del .env o usar 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Ejecutar servidor
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
