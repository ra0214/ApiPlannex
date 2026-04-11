package application

import (
	"Plannex/src/users/domain"
	"fmt"
)

type SaveDeviceToken struct {
	db domain.IDeviceToken
}

func NewSaveDeviceToken(db domain.IDeviceToken) *SaveDeviceToken {
	return &SaveDeviceToken{db: db}
}

func (sdt *SaveDeviceToken) Execute(userID int32, fcmToken string, deviceName string) error {
	if userID <= 0 {
		return fmt.Errorf("user ID debe ser válido")
	}

	if fcmToken == "" {
		return fmt.Errorf("FCM token no puede estar vacío")
	}

	err := sdt.db.SaveDeviceToken(userID, fcmToken, deviceName)
	if err != nil {
		return err
	}

	return nil
}
