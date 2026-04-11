package infraestructure

import (
	"Plannex/src/config"
	"Plannex/src/users/domain"
	"database/sql"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IUser = (*MySQL)(nil)

type MySQLDeviceToken struct {
	conn *config.Conn_MySQL
}

var _ domain.IDeviceToken = (*MySQLDeviceToken)(nil)

func NewMySQL() domain.IUser {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SaveUser(userName string, email string, password string) error {
	query := "INSERT INTO users (user_name, email, password, role) VALUES (?, ?, ?, 'guest')"
	result, err := mysql.conn.ExecutePreparedQuery(query, userName, email, password)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Usuario creado correctamente: Username:%s Email:%s", userName, email)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.User, error) {
	query := "SELECT id, user_name, email, password, auth_token, role, profile_image_path FROM users"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		var authToken, profilePath sql.NullString
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &authToken, &user.Role, &profilePath); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err)
		}
		if authToken.Valid {
			user.AuthToken = authToken.String
		}
		if profilePath.Valid {
			user.ProfileImagePath = profilePath.String
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterando sobre las filas: %v", err)
	}
	return users, nil
}

func (mysql *MySQL) UpdateUser(id int32, userName string, email string, password string) error {
	query := "UPDATE users SET user_name = ?, email = ?, password = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, userName, email, password, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Usuario actualizado correctamente: ID: %d Username:%s Email: %s", id, userName, email)
	} else {
		log.Println("[MySQL] - No se actualizó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) UpdateUserAuthToken(id int32, authToken string) error {
	query := "UPDATE users SET auth_token = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, authToken, id)
	return err
}

func (mysql *MySQL) DeleteUser(id int32) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Usuario eliminado correctamente: ID: %d", id)
	} else {
		log.Println("[MySQL] - No se eliminó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetUserByCredentials(userName string) (*domain.User, error) {
	query := "SELECT id, user_name, email, password, auth_token, role, profile_image_path FROM users WHERE user_name = ?"
	row, err := mysql.conn.FetchRow(query, userName)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	var user domain.User
	var authToken, profilePath sql.NullString
	err = row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &authToken, &user.Role, &profilePath)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado")
	}
	if authToken.Valid {
		user.AuthToken = authToken.String
	}
	if profilePath.Valid {
		user.ProfileImagePath = profilePath.String
	}

	return &user, nil
}

func (mysql *MySQL) GetUserByID(id int32) (*domain.User, error) {
	query := "SELECT id, user_name, email, password, auth_token, role, profile_image_path FROM users WHERE id = ?"
	row, err := mysql.conn.FetchRow(query, id)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	var user domain.User
	var authToken, profilePath sql.NullString
	err = row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &authToken, &user.Role, &profilePath)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado")
	}
	if authToken.Valid {
		user.AuthToken = authToken.String
	}
	if profilePath.Valid {
		user.ProfileImagePath = profilePath.String
	}

	return &user, nil
}

// Device Token Methods
func NewMySQLDeviceToken() domain.IDeviceToken {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQLDeviceToken{conn: conn}
}

func (mysql *MySQLDeviceToken) SaveDeviceToken(userID int32, fcmToken string, deviceName string) error {
	query := `INSERT INTO device_tokens (user_id, fcm_token, device_name) 
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE user_id = ?, device_name = ?, updated_at = CURRENT_TIMESTAMP`
	result, err := mysql.conn.ExecutePreparedQuery(query, userID, fcmToken, deviceName, userID, deviceName)
	if err != nil {
		return fmt.Errorf("error al guardar token de dispositivo: %v", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		log.Printf("[MySQL] - Token de dispositivo guardado: UserID:%d Token:%s", userID, fcmToken)
	}
	return nil
}

func (mysql *MySQLDeviceToken) DeleteDeviceToken(userID int32, fcmToken string) error {
	query := "DELETE FROM device_tokens WHERE user_id = ? AND fcm_token = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, userID, fcmToken)
	if err != nil {
		return fmt.Errorf("error al eliminar token de dispositivo: %v", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Token de dispositivo eliminado: UserID:%d", userID)
	}
	return nil
}

func (mysql *MySQLDeviceToken) GetDeviceTokensByUserID(userID int32) ([]domain.DeviceToken, error) {
	query := "SELECT id, user_id, fcm_token, device_name, created_at, updated_at FROM device_tokens WHERE user_id = ?"
	rows, err := mysql.conn.FetchRows(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var tokens []domain.DeviceToken
	for rows.Next() {
		var token domain.DeviceToken
		var deviceName, updatedAt sql.NullString
		if err := rows.Scan(&token.ID, &token.UserID, &token.FCMToken, &deviceName, &token.CreatedAt, &updatedAt); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		if deviceName.Valid {
			token.DeviceName = deviceName.String
		}
		if updatedAt.Valid {
			token.UpdatedAt = updatedAt.String
		}
		tokens = append(tokens, token)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}
	return tokens, nil
}

func (mysql *MySQLDeviceToken) GetDeviceTokenByFCMToken(fcmToken string) (*domain.DeviceToken, error) {
	query := "SELECT id, user_id, fcm_token, device_name, created_at, updated_at FROM device_tokens WHERE fcm_token = ?"
	row, err := mysql.conn.FetchRow(query, fcmToken)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	var token domain.DeviceToken
	var deviceName, updatedAt sql.NullString
	err = row.Scan(&token.ID, &token.UserID, &token.FCMToken, &deviceName, &token.CreatedAt, &updatedAt)
	if err != nil {
		return nil, fmt.Errorf("token de dispositivo no encontrado")
	}
	if deviceName.Valid {
		token.DeviceName = deviceName.String
	}
	if updatedAt.Valid {
		token.UpdatedAt = updatedAt.String
	}

	return &token, nil
}
