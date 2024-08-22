package lnlmodificationusecase

import (
	"time"

	lnlfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/lunch-n-learn"
	dtos "github.com/ioet/ioet-lunch-n-learn-backend/api/dtos"
	lunchnlearn "github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/lunch_n_learn"
)

type LnLModificationUseCase struct {
	Repository lnlfirebaserepository.LnLRepository
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
		return lunchnlearn.LunchNLearn{}, err
	}

	return updatedLnL, nil
}

type LnLAddAssistantUseCase struct {
	Repository lnlfirebaserepository.LnLRepository
}

func NewLnLAddAssistantUseCase(repository lnlfirebaserepository.LnLRepository) *LnLAddAssistantUseCase {
	return &LnLAddAssistantUseCase{Repository: repository}
}

func (u *LnLAddAssistantUseCase) Execute(data dtos.LnLAddAssistantIn) (lunchnlearn.LunchNLearn, error) {
	lnl, err := u.Repository.GetByID(data.ID)

	if err != nil {
		return lunchnlearn.LunchNLearn{}, err
	}

	lnl.AddAssistant(data.AssistantID)
	updatedHouse, err := u.Repository.Update(data.ID, lnl)

	if err != nil {
		return lunchnlearn.LunchNLearn{}, err
	}

	return updatedHouse, nil
}

type LnLAddPresenterUseCase struct {
	Repository lnlfirebaserepository.LnLRepository
}

func NewLnLAddPresenterUseCase(repository lnlfirebaserepository.LnLRepository) *LnLAddPresenterUseCase {
	return &LnLAddPresenterUseCase{Repository: repository}
}

func (u *LnLAddPresenterUseCase) Execute(data dtos.LnLAddPresenterIn) (lunchnlearn.LunchNLearn, error) {
	lnl, err := u.Repository.GetByID(data.ID)

	if err != nil {
		return lunchnlearn.LunchNLearn{}, err
	}

	lnl.AddPresenter(data.PresenterID)
	updatedHouse, err := u.Repository.Update(data.ID, lnl)

	if err != nil {
		return lunchnlearn.LunchNLearn{}, err
	}

	return updatedHouse, nil
}
