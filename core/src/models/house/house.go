package house

import (
	"github.com/google/uuid"
)

type House struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Points    int        `json:"points"`
	CaptainId *uuid.UUID `json:"captainId"`
}

func New(name string, captainId *uuid.UUID) House {
	return House{
		Id:        uuid.New(),
		Name:      name,
		Points:    0,
		CaptainId: captainId,
	}
}
