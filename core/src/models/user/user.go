package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PhotoURL  string    `json:"photoUrl"`
	Points    int       `json:"points"`
	HouseID   string    `json:"houseId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func New(name string, email string, photoURL string, houseID string) User {
	return User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		PhotoURL:  photoURL,
		Points:    0,
		HouseID:   houseID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (user *User) ChangeHouse(newHouseID string) {
	user.HouseID = newHouseID
	user.UpdatedAt = time.Now()
}

func (user *User) AddPoints(points int) {
	user.Points += points
	user.UpdatedAt = time.Now()
}
