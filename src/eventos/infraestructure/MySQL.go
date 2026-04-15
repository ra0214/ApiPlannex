package infraestructure

import (
	"database/sql"
	"fmt"
	"log"

	"Plannex/src/config"
	"Plannex/src/eventos/domain"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IEventos = (*MySQL)(nil)

func NewMySQL() domain.IEventos {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) CreateEvento(title, description, date string, latitude, longitude *float64, qrCodeData string, createdBy *int32) (int32, error) {
	query := `INSERT INTO evento (title, description, date, latitude, longitude, qr_code_data, created_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := mysql.conn.ExecutePreparedQuery(query, title, description, date, latitude, longitude, qrCodeData, createdBy)
	if err != nil {
		return 0, fmt.Errorf("Error al ejecutar la consulta de inserción: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error al obtener ID del evento: %v", err)
	}
	return int32(id), nil
}

func (mysql *MySQL) GetEventoById(id int32) (*domain.Eventos, error) {
	query := `SELECT id, title, description, date, latitude, longitude, qr_code_data, last_updated, created_by FROM evento WHERE id = ?`
	row, err := mysql.conn.FetchRow(query, id)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar consulta: %v", err)
	}

	var evento domain.Eventos
	var lat, lng sql.NullFloat64
	var qrCodeData, lastUpdated sql.NullString
	var createdBy sql.NullInt32
	if err := row.Scan(&evento.ID, &evento.Title, &evento.Description, &evento.Date, &lat, &lng, &qrCodeData, &lastUpdated, &createdBy); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No se encontró evento con ID: %d", id)
		}
		return nil, fmt.Errorf("Error al escanear fila: %v", err)
	}
	if lat.Valid {
		evento.Latitude = &lat.Float64
	}
	if lng.Valid {
		evento.Longitude = &lng.Float64
	}
	if qrCodeData.Valid {
		evento.QRCodeData = qrCodeData.String
	}
	if lastUpdated.Valid {
		evento.LastUpdated = lastUpdated.String
	}
	if createdBy.Valid {
		evento.CreatedBy = &createdBy.Int32
	}
	return &evento, nil
}

func (mysql *MySQL) GetAllEventos() ([]domain.Eventos, error) {
	query := `SELECT id, title, description, date, latitude, longitude, qr_code_data, last_updated, created_by FROM evento`
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar consulta: %v", err)
	}
	defer rows.Close()

	var eventos []domain.Eventos
	for rows.Next() {
		var evento domain.Eventos
		var lat, lng sql.NullFloat64
		var qrCodeData, lastUpdated sql.NullString
		var createdBy sql.NullInt32
		if err := rows.Scan(&evento.ID, &evento.Title, &evento.Description, &evento.Date, &lat, &lng, &qrCodeData, &lastUpdated, &createdBy); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		if lat.Valid {
			evento.Latitude = &lat.Float64
		}
		if lng.Valid {
			evento.Longitude = &lng.Float64
		}
		if qrCodeData.Valid {
			evento.QRCodeData = qrCodeData.String
		}
		if lastUpdated.Valid {
			evento.LastUpdated = lastUpdated.String
		}
		if createdBy.Valid {
			evento.CreatedBy = &createdBy.Int32
		}
		eventos = append(eventos, evento)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return eventos, nil
}

func (mysql *MySQL) GetEventosByCreator(creatorId int32) ([]domain.Eventos, error) {
	query := `SELECT id, title, description, date, latitude, longitude, qr_code_data, last_updated, created_by FROM evento WHERE created_by = ? ORDER BY date DESC`
	rows, err := mysql.conn.FetchRows(query, creatorId)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar consulta: %v", err)
	}
	defer rows.Close()

	var eventos []domain.Eventos
	for rows.Next() {
		var evento domain.Eventos
		var lat, lng sql.NullFloat64
		var qrCodeData, lastUpdated sql.NullString
		var createdBy sql.NullInt32
		if err := rows.Scan(&evento.ID, &evento.Title, &evento.Description, &evento.Date, &lat, &lng, &qrCodeData, &lastUpdated, &createdBy); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		if lat.Valid {
			evento.Latitude = &lat.Float64
		}
		if lng.Valid {
			evento.Longitude = &lng.Float64
		}
		if qrCodeData.Valid {
			evento.QRCodeData = qrCodeData.String
		}
		if lastUpdated.Valid {
			evento.LastUpdated = lastUpdated.String
		}
		if createdBy.Valid {
			evento.CreatedBy = &createdBy.Int32
		}
		eventos = append(eventos, evento)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return eventos, nil
}

func (mysql *MySQL) UpdateEvento(id int32, title, description, date string, latitude, longitude *float64, qrCodeData string) error {
	query := `UPDATE evento SET title = ?, description = ?, date = ?, latitude = ?, longitude = ?, qr_code_data = ? WHERE id = ?`
	_, err := mysql.conn.ExecutePreparedQuery(query, title, description, date, latitude, longitude, qrCodeData, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de actualización: %v", err)
	}
	return nil
}

func (mysql *MySQL) DeleteEvento(id int32) error {
	query := "DELETE FROM evento WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de eliminación: %v", err)
	}
	return nil
}

func (mysql *MySQL) InviteUser(eventoId int32, userId int32) error {
	query := "INSERT INTO invitacion (evento_id, user_id, estado) VALUES (?, ?, ?)"
	_, err := mysql.conn.ExecutePreparedQuery(query, eventoId, userId, "invitado")
	if err != nil {
		return fmt.Errorf("Error al invitar usuario: %v", err)
	}
	return nil
}

func (mysql *MySQL) ConfirmAttendance(eventoId int32, userId int32, status string) error {
	query := "UPDATE invitacion SET estado = ?, responded_at = CURRENT_TIMESTAMP WHERE evento_id = ? AND user_id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, status, eventoId, userId)
	if err != nil {
		return fmt.Errorf("Error al confirmar asistencia: %v", err)
	}
	return nil
}
