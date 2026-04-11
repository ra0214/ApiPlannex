package config

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var messagingClient *messaging.Client

// InitFirebase inicializa el cliente de Firebase Messaging
// credentialsPath: ruta al archivo serviceAccountKey.json
func InitFirebase(credentialsPath string) error {
	ctx := context.Background()

	// Verificar que el archivo existe
	if _, err := os.Stat(credentialsPath); err != nil {
		return fmt.Errorf("archivo de credenciales no encontrado en %s: %v", credentialsPath, err)
	}

	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error inicializando Firebase: %v", err)
		return fmt.Errorf("error inicializando Firebase: %v", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("Error obteniendo cliente de Messaging: %v", err)
		return fmt.Errorf("error obteniendo cliente de Messaging: %v", err)
	}

	messagingClient = client
	log.Println("[Firebase] ✓ Cliente de Messaging inicializado correctamente")
	return nil
}

// GetMessagingClient retorna el cliente de Firebase Messaging
func GetMessagingClient() *messaging.Client {
	if messagingClient == nil {
		log.Fatal("Firebase Messaging client no fue inicializado. Llamar a InitFirebase() primero")
	}
	return messagingClient
}

// Close cierra la conexión con Firebase
func CloseFirebase() error {
	if messagingClient != nil {
		return nil // Firebase Admin SDK no necesita cierre explícito
	}
	return nil
}
