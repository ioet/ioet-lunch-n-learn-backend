package LNLGeminiRepository

import (
	"context"
	"log"

	"github.com/ioet/ioet-lunch-n-learn-backend/api/config"
)

type LnLRepository struct {
	APIKey string
}

func NewLnLRepository(ctx context.Context) (*LnLRepository, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
		return nil, err
	}
	return &LnLRepository{
		APIKey: cfg.GeminiAPIKey,
		nil,
	}
}
