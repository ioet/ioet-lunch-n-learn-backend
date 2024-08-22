package lnllistingbyidusecase

import (
	lnlfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/lunch-n-learn"
	lunchnlearn "github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/lunch_n_learn"
)

type LnLListingByIdUseCase struct {
	Repository lnlfirebaserepository.LnLRepository
}

func NewLnLListingByIdUseCase(repository lnlfirebaserepository.LnLRepository) *LnLListingByIdUseCase {
	return &LnLListingByIdUseCase{Repository: repository}
}

func (u *LnLListingByIdUseCase) Execute(id string) (lunchnlearn.LunchNLearn, error) {
	lnl, err := u.Repository.GetByID(id)

	if err != nil {
		return lnl, err
	}

	return lnl, nil
}
