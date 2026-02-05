package application

import "Plannex/src/eventos/domain"

type DeleteEventos struct {
	db domain.IEventos
}

func NewDeleteEventos(db domain.IEventos) *DeleteEventos {
	return &DeleteEventos{db: db}
}

func (dp *DeleteEventos) Execute(id int32) error {
	return dp.db.DeleteEvento(id)
}
