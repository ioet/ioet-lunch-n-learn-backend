package GenerateQuizUseCase


import lnlfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/lunchNLearn"

type LnLModificationUseCase struct {
	FireBaseRepository lnlfirebaserepository.LnLRepository
	LLMRepository lnlfirebaserepository.LnLRepository
}

func NewLnLModificationUseCase(repository lnlfirebaserepository.LnLRepository) *LnLModificationUseCase {
	return &LnLModificationUseCase{Repository: repository}
}

func (u *LnLModificationUseCase) Execute(data dtos.LnLUpdateIn) (lunchnlearn.LunchNLearn, error) {
	lnl, err := u.Repository.GetByID(data.ID)

	if err != nil {
		return lunchnlearn.LunchNLearn{}, err
	}

	date, err := time.Parse("2006-01-02 15:04:05", data.PresentationDate)

	if err != nil {
		return lunchnlearn.LunchNLearn{}, err
	}

	lnl.Name = data.Name
	lnl.PresentationDate = date

	updatedLnL, err := u.Repository.Update(data.ID, lnl)

	if err != nil {
		return lunchNLearn.LunchNLearn{}, err
	}

	return updatedLnL, nil
}