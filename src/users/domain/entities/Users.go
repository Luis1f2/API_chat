package entities

type User struct {
	ID       int
	Username string
	Password string
}

func NewUser(username, password string) *User {

	user := User{ID: 1, Username: username, Password: password}
	return &user
}
