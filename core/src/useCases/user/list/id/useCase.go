package userlistingbyidusecase

import (
	userfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/user"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/user"
)

type UserListingByIdUseCase struct {
	Repository userfirebaserepository.UserRepository
}

func NewUserListingByIdUseCase(repository userfirebaserepository.UserRepository) *UserListingByIdUseCase {
	return &UserListingByIdUseCase{Repository: repository}
}

func (u *UserListingByIdUseCase) Execute(id string) (user.User, error) {
	user, err := u.Repository.GetByID(id)

	if err != nil {
		return user, err
	}

	return user, nil
}
