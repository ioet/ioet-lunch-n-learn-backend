package house

import (
	"time"

	"github.com/google/uuid"
)

type House struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Points    int       `json:"points"`
	CaptainID string    `json:"captainId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func New(name string, captainID string) House {
	return House{
		ID:        uuid.New().String(),
		Name:      name,
		Points:    0,
		CaptainID: captainID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (house *House) AddPoints(points int) {
	house.Points += points
	house.UpdatedAt = time.Now()
}
