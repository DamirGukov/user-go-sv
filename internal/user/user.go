package user

import (
	"encoding/json"
	"github.com/google/uuid"
)

type User struct {
	id           uuid.UUID
	login        string
	passwordHash string
	role         Role
}

func CreateUser(login string, passwordHash string) *User {
	return &User{
		id:           uuid.New(),
		login:        login,
		passwordHash: passwordHash,
		role:         *CreateUserRole(),
	}
}

func (u *User) GetId() uuid.UUID {
	return u.id
}

func (u *User) ChangeRole(r Role) {
	u.role = r
}

func (u *User) GetRole() Role {
	return u.role
}

func (u *User) GetLogin() string {
	return u.login
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id    uuid.UUID `json:"id"`
		Login string    `json:"login"`
		Role  string    `json:"role"`
	}{
		Id:    u.id,
		Login: u.login,
		Role:  u.role.GetRole(),
	})
}
