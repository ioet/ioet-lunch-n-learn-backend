package useCasesfactories

import (
	"context"
)

func LLMUseCase() *housefirebaserepository.LnLGeminiRepository {
	firebaseRepository, err := housefirebaserepository.NewHouseRepository(context.Background())

	if err != nil {
		panic("Error initializing the House repository: " + err.Error())
	}

	return firebaseRepository
}
