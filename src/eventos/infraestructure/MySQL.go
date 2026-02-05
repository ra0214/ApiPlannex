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

func (mysql *MySQL) CreateEvento(nombre string, fecha string) error {
	query := "INSERT INTO evento (nombre, fecha) VALUES (?, ?)"
	_, err := mysql.conn.ExecutePreparedQuery(query, nombre, fecha)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de inserción: %v", err)
	}
	return nil
}

func (mysql *MySQL) GetEventoById(id int32) (*domain.Eventos, error) {
	query := "SELECT id, nombre, fecha FROM evento WHERE id = ?"
	row, err := mysql.conn.FetchRow(query, id)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar consulta: %v", err)
	}

	var evento domain.Eventos
	if err := row.Scan(&evento.ID, &evento.Nombre, &evento.Fecha); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No se encontró evento con ID: %d", id)
		}
		return nil, fmt.Errorf("Error al escanear fila: %v", err)
	}
	return &evento, nil
}

func (mysql *MySQL) GetAllEventos() ([]domain.Eventos, error) {
	query := "SELECT id, nombre, fecha FROM evento"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar consulta: %v", err)
	}
	defer rows.Close()

	var eventos []domain.Eventos
	for rows.Next() {
		var evento domain.Eventos
		if err := rows.Scan(&evento.ID, &evento.Nombre, &evento.Fecha); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		eventos = append(eventos, evento)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return eventos, nil
}

func (mysql *MySQL) UpdateEvento(id int32, nombre string, fecha string) error {
	query := "UPDATE evento SET nombre = ?, fecha = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, nombre, fecha, id)
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
	query := "UPDATE invitacion SET estado = ? WHERE evento_id = ? AND user_id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, status, eventoId, userId)
	if err != nil {
		return fmt.Errorf("Error al confirmar asistencia: %v", err)
	}
	return nil
}
