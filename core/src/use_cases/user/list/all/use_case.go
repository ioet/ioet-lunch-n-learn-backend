package userlistingusecase

import (
	userfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/user"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/user"
)

type UserListingUseCase struct {
	Repository userfirebaserepository.UserRepository
}

func NewUserListingUseCase(repository userfirebaserepository.UserRepository) *UserListingUseCase {
	return &UserListingUseCase{Repository: repository}
}

func (u *UserListingUseCase) Execute() ([]user.User, error) {
	users, err := u.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
