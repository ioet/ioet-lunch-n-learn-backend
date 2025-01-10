package LNLGeminiRepository

import (
	"context"
	"log"

	"github.com/ioet/ioet-lunch-n-learn-backend/api/config"
	models "github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/LLM"
)

// type LnLGeminiRepository struct {
// 	APIKey string
// }

type LnLGeminiRepository struct {
	geminiLLMConnection *GeminiAPIClient
}

func NewLnLRepository(ctx context.Context) (*LnLGeminiRepository, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
		return nil, err
	}
	return &LnLGeminiRepository{APIKey: cfg.GeminiAPIKey}, nil
}

func (r *LnLGeminiRepository) UploadDocument(document byte) error {
	return nil
}

func (r *LnLGeminiRepository) SendMessage(message string) (models.LLMResponseMessage, error) {
	return models.LLMResponseMessage{}, nil
}
