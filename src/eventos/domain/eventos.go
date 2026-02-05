package domain

type IEventos interface {
	CreateEvento(nombre string, fecha string) error
	InviteUser(eventoId int32, userId int32) error
	ConfirmAttendance(eventoId int32, userId int32, status string) error
	GetAllEventos() ([]Eventos, error)
	GetEventoById(id int32) (*Eventos, error)
	UpdateEvento(id int32, nombre string, fecha string) error
	DeleteEvento(id int32) error
}

type Eventos struct {
	ID     int32  `json:"id"`
	Nombre string `json:"nombre"`
	Fecha  string `json:"fecha"`
}

type Invitacion struct {
	EventoID int32  `json:"evento_id"`
	UserID   int32  `json:"user_id"`
	Estado   string `json:"estado"`
}

func NewEventos(nombre string, fecha string) *Eventos {
	return &Eventos{ID: 1, Nombre: nombre, Fecha: fecha}
}

func NewInvitacion(eventoId int32, userId int32, estado string) *Invitacion {
	return &Invitacion{EventoID: eventoId, UserID: userId, Estado: estado}
}

func (p *Eventos) CreateEvento(nombre string, fecha string) error {
	// Lógica para crear un evento
	return nil
}

// Nueva función para invitar usuarios
func (p *Eventos) InviteUser(eventoId int32, userId int32) error {
	// Lógica para invitar un usuario a un evento
	return nil
}

// Nueva función para confirmar asistencia
func (p *Eventos) ConfirmAttendance(eventoId int32, userId int32, status string) error {
	// Lógica para confirmar asistencia
	return nil
}

// Nueva función para obtener todos los eventos
func (p *Eventos) GetAllEventos() ([]Eventos, error) {
	// Lógica para obtener todos los eventos
	return []Eventos{}, nil
}
