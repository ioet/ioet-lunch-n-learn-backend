package usercreationusecase

import (
	userfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/user"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/user"
)

type UserCreationUseCase struct {
	Repository userfirebaserepository.UserRepository
}

func NewUserCreationUseCase(repository userfirebaserepository.UserRepository) *UserCreationUseCase {
	return &UserCreationUseCase{Repository: repository}
}

func (u *UserCreationUseCase) Execute(data user.User) (user.User, error) {
	_, err := u.Repository.Create(data)

	if err != nil {
		return user.User{}, err
	}

	return data, nil
}
