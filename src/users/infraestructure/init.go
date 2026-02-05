package infraestructure

import (
	"Plannex/src/config"
	"log"
)

func InitUser() error {
	log.Println("Inicializando usuarios...")

	db, err := config.GetDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("Conexión a la base de datos para usuarios establecida correctamente")
	return nil
}