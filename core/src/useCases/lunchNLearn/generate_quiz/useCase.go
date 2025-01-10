package GenerateQuizUseCase

import (
	lnlfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/lunchNLearn"
	"github.com/ioet/ioet-lunch-n-learn-backend/api/dtos"
	repositories "github.com/ioet/ioet-lunch-n-learn-backend/core/src/repositories"
)

type LnLGenerateQuizUseCase struct {
	FireBaseRepository lnlfirebaserepository.LnLRepository
	LLMRepository      repositories.LLMRepository
}

func NewLnLGenerateQuizUseCase(firebaseRepository lnlfirebaserepository.LnLRepository, LLMRepository repositories.LLMRepository) *LnLGenerateQuizUseCase {
	return &LnLGenerateQuizUseCase{
		FireBaseRepository: firebaseRepository,
		LLMRepository:      LLMRepository,
	}
}

func (u *LnLGenerateQuizUseCase) Execute(data dtos.LnLGenerateQuizIn) (string, error) {
	err := u.LLMRepository.UploadDocument(data.PDF) // Responds with URL
	if err != nil {
		return "There was an error while generating the quiz", err // Change the error for a custom error.
	}

	quiz = u.LLMRepository.SendMessage(data) // URL is provided here

	quiz = u.FireBaseRepository().SaveQuiz(quiz)

	return "Quiz generated successfully", nil
}
