package lunchnlearn

import (
	"time"

	"github.com/google/uuid"
)

type LunchNLearn struct {
	ID               uuid.UUID    `json:"id"`
	Name             string       `json:"name"`
	Assistants       *[]uuid.UUID `json:"assistants"`
	Presenters       []uuid.UUID  `json:"presenters"`
	CreatedAt        time.Time    `json:"createdAt"`
	UpdatedAt        time.Time    `json:"updatedAt"`
	PresentationDate time.Time    `json:"presentationDate"`
}

func New(name string, assistants *[]uuid.UUID, presenters []uuid.UUID, presentationDate string) LunchNLearn {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", presentationDate)

	if err != nil {
		panic(err)
	}

	return LunchNLearn{
		ID:               uuid.New(),
		Name:             name,
		Assistants:       assistants,
		Presenters:       presenters,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		PresentationDate: parsedTime,
	}
}

func (lunchNLearn *LunchNLearn) AddAssistant(assistantID uuid.UUID) {
	*lunchNLearn.Assistants = append(*lunchNLearn.Assistants, assistantID)
	lunchNLearn.UpdatedAt = time.Now()
}
