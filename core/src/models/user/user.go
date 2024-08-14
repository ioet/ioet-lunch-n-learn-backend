package user

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID  `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	PhotoUrl *string    `json:"photoUrl,omitempty"`
	HouseId  *uuid.UUID `json:"houseId"`
}

func New(name string, email string, photoUrl *string, houseId *uuid.UUID) User {
	return User{
		Id:       uuid.New(),
		Name:     name,
		Email:    email,
		PhotoUrl: photoUrl,
		HouseId:  houseId,
	}
}
