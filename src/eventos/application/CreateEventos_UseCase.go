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

func (cp *CreateEventos) Execute(nombre string, fecha string) error {
	return cp.db.CreateEvento(nombre, fecha)
}
