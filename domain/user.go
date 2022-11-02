package domain

import "github.com/google/uuid"

type User struct {
	Name string
	Id   string
}

func NewUser(name string) *User {
	return &User{
		Id:   uuid.New().String(),
		Name: name,
	}
}
