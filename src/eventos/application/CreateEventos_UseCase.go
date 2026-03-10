package application

import (
	"Plannex/src/eventos/domain"
)

type CreateEventos struct {
	db domain.IEventos
}

func NewCreateEventos(db domain.IEventos) *CreateEventos {
	return &CreateEventos{db: db}
}

func (cp *CreateEventos) Execute(title, description, date string, latitude, longitude *float64, qrCodeData string, createdBy *int32) (int32, error) {
	return cp.db.CreateEvento(title, description, date, latitude, longitude, qrCodeData, createdBy)
}
