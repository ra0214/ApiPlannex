package application

import "Plannex/src/eventos/domain"

type GetEventosById struct {
	db domain.IEventos
}

func NewGetEventosById(db domain.IEventos) *GetEventosById {
	return &GetEventosById{db: db}
}

func (gp *GetEventosById) Execute(id int32) (*domain.Eventos, error) {
	return gp.db.GetEventoById(id)
}
