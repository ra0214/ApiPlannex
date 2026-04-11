package infraestructure

import (
	"context"
	"fmt"
	"log"

	"Plannex/src/config"
	"Plannex/src/users/infraestructure"
)

// NotifyUserInvited envía notificación cuando se invita a un usuario a un evento
func NotifyUserInvited(userID int32, eventoTitulo string, eventoID int32) {
	go func() {
		ctx := context.Background()
		notificationService := config.NewNotificationService()

		deviceTokenRepo := infraestructure.NewMySQLDeviceToken()
		tokens, err := deviceTokenRepo.GetDeviceTokensByUserID(userID)
		if err != nil {
			log.Printf("[FCM] Error obteniendo tokens para usuario %d: %v", userID, err)
			return
		}

		var fcmTokens []string
		for _, token := range tokens {
			fcmTokens = append(fcmTokens, token.FCMToken)
		}

		if len(fcmTokens) == 0 {
			log.Printf("[FCM] Usuario %d no tiene dispositivos registrados", userID)
			return
		}

		_, err = notificationService.SendNotificationToMultiple(
			ctx,
			fcmTokens,
			"¡Te han invitado!",
			"Has sido invitado a: "+eventoTitulo,
			map[string]string{
				"evento_id": fmt.Sprintf("%d", eventoID),
				"tipo":      "invitacion",
			},
		)

		if err != nil {
			log.Printf("[FCM] Error enviando invitación: %v", err)
		}
	}()
}

// NotifyEventCreated envía notificación cuando se crea un evento a los usuarios especificados
func NotifyEventCreated(eventoID int32, eventoTitulo string, eventoDescripcion string, usuariosANotificar []int32) {
	go func() {
		ctx := context.Background()
		notificationService := config.NewNotificationService()

		if len(usuariosANotificar) == 0 {
			log.Println("[FCM] No hay usuarios para notificar")
			return
		}

		deviceTokenRepo := infraestructure.NewMySQLDeviceToken()

		var fcmTokens []string
		for _, userID := range usuariosANotificar {
			tokens, err := deviceTokenRepo.GetDeviceTokensByUserID(userID)
			if err != nil {
				log.Printf("[FCM] Error obteniendo tokens para usuario %d: %v", userID, err)
				continue
			}

			for _, token := range tokens {
				fcmTokens = append(fcmTokens, token.FCMToken)
			}
		}

		if len(fcmTokens) == 0 {
			log.Println("[FCM] No hay dispositivos registrados para notificar")
			return
		}

		_, err := notificationService.SendNotificationToMultiple(
			ctx,
			fcmTokens,
			"Nuevo Evento",
			eventoTitulo,
			map[string]string{
				"evento_id":   fmt.Sprintf("%d", eventoID),
				"titulo":      eventoTitulo,
				"descripcion": eventoDescripcion,
				"tipo":        "evento_creado",
			},
		)

		if err != nil {
			log.Printf("[FCM] Error enviando notificación de evento: %v", err)
		}
	}()
}

// NotifyAttendanceConfirmed envía notificación cuando alguien confirma asistencia
func NotifyAttendanceConfirmed(eventoID int32, usuarioID int32, estado string, usuariosANotificar []int32) {
	go func() {
		ctx := context.Background()
		notificationService := config.NewNotificationService()

		deviceTokenRepo := infraestructure.NewMySQLDeviceToken()

		var fcmTokens []string
		for _, userID := range usuariosANotificar {
			if userID == usuarioID {
				continue // No notificar al mismo usuario
			}

			tokens, err := deviceTokenRepo.GetDeviceTokensByUserID(userID)
			if err != nil {
				continue
			}

			for _, token := range tokens {
				fcmTokens = append(fcmTokens, token.FCMToken)
			}
		}

		if len(fcmTokens) == 0 {
			return
		}

		var mensaje string
		switch estado {
		case "asistira":
			mensaje = "Usuario confirmó que asistirá"
		case "quiza":
			mensaje = "Usuario confirmó que quizá asista"
		case "no_asistira":
			mensaje = "Usuario confirmó que no asistirá"
		default:
			mensaje = "Cambio de estado de asistencia"
		}

		_, err := notificationService.SendNotificationToMultiple(
			ctx,
			fcmTokens,
			"Actualización de asistencia",
			mensaje,
			map[string]string{
				"evento_id": fmt.Sprintf("%d", eventoID),
				"tipo":      "asistencia_confirmada",
				"estado":    estado,
			},
		)

		if err != nil {
			log.Printf("[FCM] Error enviando notificación de asistencia: %v", err)
		}
	}()
}
