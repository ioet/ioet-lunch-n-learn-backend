package userlistingbyidusecase

import (
	userfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/user"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/user"
)

type UserListingByIDUseCase struct {
	Repository userfirebaserepository.UserRepository
}

func NewUserListingByIDUseCase(repository userfirebaserepository.UserRepository) *UserListingByIDUseCase {
	return &UserListingByIDUseCase{Repository: repository}
}

func (u *UserListingByIDUseCase) Execute(id string) (user.User, error) {
	user, err := u.Repository.GetByID(id)

	if err != nil {
		return user, err
	}

	return user, nil
}
