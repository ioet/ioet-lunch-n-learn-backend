package lunchnlearn

import (
	"time"

	"github.com/google/uuid"
)

type LunchNLearn struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Assistants       []string  `json:"assistants"`
	Presenters       []string  `json:"presenters"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	PresentationDate time.Time `json:"presentationDate"`
}

func New(name string, presentationDate string) LunchNLearn {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", presentationDate)

	if err != nil {
		panic(err)
	}

	return LunchNLearn{
		ID:               uuid.New().String(),
		Name:             name,
		Assistants:       []string{},
		Presenters:       []string{},
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		PresentationDate: parsedTime,
	}
}

func (lunchNLearn *LunchNLearn) AddAssistant(assistantID string) {
	lunchNLearn.Assistants = append(lunchNLearn.Assistants, assistantID)
	lunchNLearn.UpdatedAt = time.Now()
}

func (lunchNLearn *LunchNLearn) AddPresenter(presenterID string) {
	lunchNLearn.Presenters = append(lunchNLearn.Presenters, presenterID)
	lunchNLearn.UpdatedAt = time.Now()
}
