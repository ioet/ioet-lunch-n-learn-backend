package lnllistingusecase

import (
	lnlfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/lunch-n-learn"
	lunchnlearn "github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/lunch_n_learn"
)

type LnLListingUseCase struct {
	Repository lnlfirebaserepository.LnLRepository
}

func NewLnLListingUseCase(repository lnlfirebaserepository.LnLRepository) *LnLListingUseCase {
	return &LnLListingUseCase{Repository: repository}
}

func (u *LnLListingUseCase) Execute() ([]lunchnlearn.LunchNLearn, error) {
	lnls, err := u.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return lnls, nil
}
