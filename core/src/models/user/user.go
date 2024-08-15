package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	PhotoURL  *string    `json:"photoUrl,omitempty"`
	HouseID   *uuid.UUID `json:"houseId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func New(name string, email string, photoURL *string, houseID *uuid.UUID) User {
	return User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		PhotoURL:  photoURL,
		HouseID:   houseID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
