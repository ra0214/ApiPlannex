package application

import (
	"Plannex/src/eventos/domain"
)

type ViewMyEventos struct {
	db domain.IEventos
}

func NewViewMyEventos(db domain.IEventos) *ViewMyEventos {
	return &ViewMyEventos{db: db}
}

func (vmp *ViewMyEventos) Execute(creatorId int32) ([]domain.Eventos, error) {
	return vmp.db.GetEventosByCreator(creatorId)
}
