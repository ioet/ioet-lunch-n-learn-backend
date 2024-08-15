package house

import (
	"time"

	"github.com/google/uuid"
)

type House struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Points    int        `json:"points"`
	CaptainId *uuid.UUID `json:"captainId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func New(name string, captainId *uuid.UUID) House {
	return House{
		Id:        uuid.New(),
		Name:      name,
		Points:    0,
		CaptainId: captainId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (house *House) AddPoints(points int) {
	house.Points += points
	house.UpdatedAt = time.Now()
}
