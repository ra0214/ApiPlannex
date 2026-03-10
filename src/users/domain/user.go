package domain

type IUser interface {
	SaveUser(userName string, email string, password string) error
	DeleteUser(id int32) error
	UpdateUser(id int32, userName string, email string, password string) error
	UpdateUserAuthToken(id int32, authToken string) error
	GetAll() ([]User, error)
	GetUserByCredentials(userName string) (*User, error)
	GetUserByID(id int32) (*User, error)
}

type User struct {
	ID                int32  `json:"id"`
	UserName          string `json:"userName"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	AuthToken         string `json:"auth_token,omitempty"`
	Role              string `json:"role"` // "admin" | "guest"
	ProfileImagePath  string `json:"profile_image_path,omitempty"`
}

func NewUser(userName string, email string, password string) *User {
	return &User{UserName: userName, Email: email, Password: password, Role: "guest"}
}

func (u *User) SetUserName(userName string) {
	u.UserName = userName
}
