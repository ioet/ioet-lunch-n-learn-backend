package repositoryfactories

import (
	"context"

	geminiLLMRepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/gemini/lunchNLearn"
)

func GeminiLLMRepository() *geminiLLMRepository.LnLGeminiRepository {
	geminiRepository, err := geminiLLMRepository.NewLnLRepository(context.Background())
	if err != nil {
		panic("Error initializing the Gemini LLM repository: " + err.Error())
	}
	return geminiRepository
}
