package lnlcreationusecase

import (
	lnlfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/lunchNLearn"
	lunchnlearn "github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/lunchNLearn"
)

type LnLCreationUseCase struct {
	Repository lnlfirebaserepository.LnLRepository
}

func NewLnLCreationUseCase(repository lnlfirebaserepository.LnLRepository) *LnLCreationUseCase {
	return &LnLCreationUseCase{Repository: repository}
}

func (u *LnLCreationUseCase) Execute(data lunchnlearn.LunchNLearn) (lunchnlearn.LunchNLearn, error) {
	createdHouse, err := u.Repository.Create(data)

	if err != nil {
		return lunchnlearn.LunchNLearn{}, err
	}

	return createdHouse, nil
}
