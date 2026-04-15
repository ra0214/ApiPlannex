package domain

type IUser interface {
	SaveUser(userName string, email string, password string) error
	DeleteUser(id int32) error
	UpdateUser(id int32, userName string, email string, password string) error
	UpdateUserAuthToken(id int32, authToken string) error
	GetAll() ([]User, error)
	GetUserByCredentials(userName string) (*User, error)
	GetUserByID(id int32) (*User, error)
	GetUserByToken(token string) (*User, error)
}

type IDeviceToken interface {
	SaveDeviceToken(userID int32, fcmToken string, deviceName string) error
	DeleteDeviceToken(userID int32, fcmToken string) error
	GetDeviceTokensByUserID(userID int32) ([]DeviceToken, error)
	GetDeviceTokenByFCMToken(fcmToken string) (*DeviceToken, error)
}

type User struct {
	ID               int32  `json:"id"`
	UserName         string `json:"userName"`
	Email            string `json:"email"`
	Password         string `json:"password,omitempty"`
	AuthToken        string `json:"auth_token,omitempty"`
	Role             string `json:"role"` // "admin" | "guest"
	ProfileImagePath string `json:"profile_image_path,omitempty"`
}

type DeviceToken struct {
	ID         int32  `json:"id"`
	UserID     int32  `json:"user_id"`
	FCMToken   string `json:"fcm_token"`
	DeviceName string `json:"device_name,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}

func NewUser(userName string, email string, password string) *User {
	return &User{UserName: userName, Email: email, Password: password, Role: "guest"}
}

func (u *User) SetUserName(userName string) {
	u.UserName = userName
}

func NewDeviceToken(userID int32, fcmToken string, deviceName string) *DeviceToken {
	return &DeviceToken{UserID: userID, FCMToken: fcmToken, DeviceName: deviceName}
}
