package repositories

import (
	models "github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/LLM"
)

type LLMRepository interface {
	UploadDocument(document byte) error
	SendMessage(message string) (models.LLMResponseMessage, error)
}
