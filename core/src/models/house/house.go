package house

import (
	"time"

	"github.com/google/uuid"
)

type House struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Points    int        `json:"points"`
	CaptainID *uuid.UUID `json:"captainId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func New(name string, captainID *uuid.UUID) House {
	return House{
		ID:        uuid.New(),
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
