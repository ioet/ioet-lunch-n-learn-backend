package lnllistingbyidusecase

import (
	lnlfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/lunchNLearn"
	lunchnlearn "github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/lunchNLearn"
)

type LnLListingByIDUseCase struct {
	Repository lnlfirebaserepository.LnLRepository
}

func NewLnLListingByIDUseCase(repository lnlfirebaserepository.LnLRepository) *LnLListingByIDUseCase {
	return &LnLListingByIDUseCase{Repository: repository}
}

func (u *LnLListingByIDUseCase) Execute(id string) (lunchnlearn.LunchNLearn, error) {
	lnl, err := u.Repository.GetByID(id)

	if err != nil {
		return lnl, err
	}

	return lnl, nil
}
