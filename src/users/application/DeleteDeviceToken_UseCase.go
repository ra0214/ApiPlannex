package application

import (
	"Plannex/src/users/domain"
	"fmt"
)

type DeleteDeviceToken struct {
	db domain.IDeviceToken
}

func NewDeleteDeviceToken(db domain.IDeviceToken) *DeleteDeviceToken {
	return &DeleteDeviceToken{db: db}
}

func (ddt *DeleteDeviceToken) Execute(userID int32, fcmToken string) error {
	if userID <= 0 {
		return fmt.Errorf("user ID debe ser válido")
	}

	if fcmToken == "" {
		return fmt.Errorf("FCM token no puede estar vacío")
	}

	err := ddt.db.DeleteDeviceToken(userID, fcmToken)
	if err != nil {
		return err
	}

	return nil
}
