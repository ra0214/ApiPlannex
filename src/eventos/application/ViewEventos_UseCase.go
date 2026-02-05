package application

import "Plannex/src/eventos/domain"

type ViewEventos struct {
	db domain.IEventos
}

func NewViewEventos(db domain.IEventos) *ViewEventos {
	return &ViewEventos{db: db}
}

func (vp *ViewEventos) Execute() ([]domain.Eventos, error) {
	return vp.db.GetAllEventos()
}
