package application

import (
	"Plannex/src/eventos/domain"
)

type EditEventos struct {
	db domain.IEventos
}

func NewEditEventos(db domain.IEventos) *EditEventos {
	return &EditEventos{db: db}
}

func (ep *EditEventos) Execute(id int32, title, description, date string, latitude, longitude *float64, qrCodeData string) error {
	return ep.db.UpdateEvento(id, title, description, date, latitude, longitude, qrCodeData)
}
