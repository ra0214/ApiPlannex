package domain

type IEventos interface {
	CreateEvento(title, description, date string, latitude, longitude *float64, qrCodeData string, createdBy *int32) (int32, error)
	InviteUser(eventoId int32, userId int32) error
	ConfirmAttendance(eventoId int32, userId int32, status string) error
	GetAllEventos() ([]Eventos, error)
	GetEventoById(id int32) (*Eventos, error)
	GetEventosByCreator(creatorId int32) ([]Eventos, error)
	UpdateEvento(id int32, title, description, date string, latitude, longitude *float64, qrCodeData string) error
	DeleteEvento(id int32) error
}

type Eventos struct {
	ID          int32    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
	QRCodeData  string   `json:"qr_code_data,omitempty"`
	LastUpdated string   `json:"last_updated,omitempty"`
	CreatedBy   *int32   `json:"created_by,omitempty"`
}

type Invitacion struct {
	EventoID int32  `json:"evento_id"`
	UserID   int32  `json:"user_id"`
	Estado   string `json:"estado"`
}

func NewEventos(title, description, date string, latitude, longitude *float64, qrCodeData string) *Eventos {
	return &Eventos{
		Title:       title,
		Description: description,
		Date:        date,
		Latitude:    latitude,
		Longitude:   longitude,
		QRCodeData:  qrCodeData,
	}
}

func NewInvitacion(eventoId int32, userId int32, estado string) *Invitacion {
	return &Invitacion{EventoID: eventoId, UserID: userId, Estado: estado}
}
