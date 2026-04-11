package config

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/v4/messaging"
)

type NotificationService struct {
	client *messaging.Client
}

// NewNotificationService crea una nueva instancia del servicio de notificaciones
func NewNotificationService() *NotificationService {
	return &NotificationService{
		client: GetMessagingClient(),
	}
}

// SendNotification envía una notificación a un dispositivo específico
func (ns *NotificationService) SendNotification(
	ctx context.Context,
	fcmToken string,
	titulo string,
	descripcion string,
	data map[string]string,
) (string, error) {

	if fcmToken == "" {
		return "", fmt.Errorf("FCM token no puede estar vacío")
	}

	message := &messaging.Message{
		Token: fcmToken,
		Notification: &messaging.Notification{
			Title: titulo,
			Body:  descripcion,
		},
		Data: data,
	}

	response, err := ns.client.Send(ctx, message)
	if err != nil {
		log.Printf("[Firebase Error] Error enviando a %s: %v", fcmToken, err)
		return "", fmt.Errorf("error enviando notificación: %v", err)
	}

	log.Printf("[Firebase] ✓ Notificación enviada exitosamente a %s. Response: %s", fcmToken, response)
	return response, nil
}

// SendNotificationToMultiple envía notificaciones a varios dispositivos
func (ns *NotificationService) SendNotificationToMultiple(
	ctx context.Context,
	fcmTokens []string,
	titulo string,
	descripcion string,
	data map[string]string,
) (map[string]string, error) {

	if len(fcmTokens) == 0 {
		return nil, fmt.Errorf("lista de tokens vacía")
	}

	messages := make([]*messaging.Message, 0)
	for _, token := range fcmTokens {
		message := &messaging.Message{
			Token: token,
			Notification: &messaging.Notification{
				Title: titulo,
				Body:  descripcion,
			},
			Data: data,
		}
		messages = append(messages, message)
	}

	br, err := ns.client.SendAll(ctx, messages)
	if err != nil {
		log.Printf("[Firebase Error] Error enviando notificaciones: %v", err)
		return nil, fmt.Errorf("error enviando notificaciones: %v", err)
	}

	resultado := make(map[string]string)
	for i, resp := range br.Responses {
		if resp.Success {
			resultado[fcmTokens[i]] = "enviado exitosamente"
		} else {
			resultado[fcmTokens[i]] = fmt.Sprintf("error: %v", resp.Error)
		}
	}

	log.Printf("[Firebase] ✓ Enviadas %d notificaciones: %d éxitos, %d errores",
		len(fcmTokens), br.SuccessCount, len(fcmTokens)-br.SuccessCount)

	return resultado, nil
}

// SendNotificationToTopic envía a todos los dispositivos suscritos a un tema
func (ns *NotificationService) SendNotificationToTopic(
	ctx context.Context,
	topic string,
	titulo string,
	descripcion string,
	data map[string]string,
) (string, error) {

	if topic == "" {
		return "", fmt.Errorf("topic no puede estar vacío")
	}

	message := &messaging.Message{
		Topic: topic,
		Notification: &messaging.Notification{
			Title: titulo,
			Body:  descripcion,
		},
		Data: data,
	}

	response, err := ns.client.Send(ctx, message)
	if err != nil {
		log.Printf("[Firebase Error] Error enviando a topic %s: %v", topic, err)
		return "", fmt.Errorf("error enviando notificación a topic: %v", err)
	}

	log.Printf("[Firebase] ✓ Notificación enviada a topic '%s'. Response: %s", topic, response)
	return response, nil
}

// SubscribeToTopic suscribe un dispositivo a un tema
func (ns *NotificationService) SubscribeToTopic(
	ctx context.Context,
	topic string,
	fcmTokens []string,
) error {

	if topic == "" {
		return fmt.Errorf("topic no puede estar vacío")
	}

	if len(fcmTokens) == 0 {
		return fmt.Errorf("lista de tokens vacía")
	}

	resp, err := ns.client.SubscribeToTopic(ctx, fcmTokens, topic)
	if err != nil {
		log.Printf("[Firebase Error] Error suscribiendo a topic %s: %v", topic, err)
		return fmt.Errorf("error suscribiendo a topic: %v", err)
	}

	log.Printf("[Firebase] ✓ %d dispositivos suscritos al topic '%s'", resp.SuccessCount, topic)
	return nil
}

// UnsubscribeFromTopic desuscribe un dispositivo de un tema
func (ns *NotificationService) UnsubscribeFromTopic(
	ctx context.Context,
	topic string,
	fcmTokens []string,
) error {

	if topic == "" {
		return fmt.Errorf("topic no puede estar vacío")
	}

	if len(fcmTokens) == 0 {
		return fmt.Errorf("lista de tokens vacía")
	}

	resp, err := ns.client.UnsubscribeFromTopic(ctx, fcmTokens, topic)
	if err != nil {
		log.Printf("[Firebase Error] Error desuscribiendo de topic %s: %v", topic, err)
		return fmt.Errorf("error desuscribiendo de topic: %v", err)
	}

	log.Printf("[Firebase] ✓ %d dispositivos desuscritos del topic '%s'", resp.SuccessCount, topic)
	return nil
}
