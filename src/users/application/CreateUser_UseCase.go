package application

import (
	"Plannex/src/users/domain"

	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	db domain.IUser
}

func NewCreateUser(db domain.IUser) *CreateUser {
	return &CreateUser{db: db}
}

func (cu *CreateUser) Execute(userName string, email string, password string) error {
	// Generar hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Guardar usuario
	err = cu.db.SaveUser(userName, email, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}
