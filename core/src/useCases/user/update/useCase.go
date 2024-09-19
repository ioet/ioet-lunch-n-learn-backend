package usermodificationusecase

import (
	userfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/user"
	"github.com/ioet/ioet-lunch-n-learn-backend/api/dtos"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/user"
)

type UserHouseChangeUseCase struct {
	Repository userfirebaserepository.UserRepository
}

func NewUserHouseChangeUseCase(repository userfirebaserepository.UserRepository) *UserHouseChangeUseCase {
	return &UserHouseChangeUseCase{Repository: repository}
}

func (u *UserHouseChangeUseCase) Execute(data dtos.UserChangeHouseIn) (user.User, error) {
	usr, err := u.Repository.GetByID(data.ID)

	if err != nil {
		return user.User{}, err
	}

	usr.ChangeHouse(data.HouseID)
	updatedUser, err := u.Repository.Update(data.ID, usr)

	if err != nil {
		return user.User{}, err
	}

	return updatedUser, nil
}

type UserAddPointsUseCase struct {
	Repository userfirebaserepository.UserRepository
}

func NewUserAddPointsUseCase(repository userfirebaserepository.UserRepository) *UserAddPointsUseCase {
	return &UserAddPointsUseCase{Repository: repository}
}

func (u *UserAddPointsUseCase) Execute(data dtos.UserAddPointsIn) (user.User, error) {
	usr, err := u.Repository.GetByID(data.ID)

	if err != nil {
		return user.User{}, err
	}

	usr.AddPoints(data.NewPoints)
	updatedUser, err := u.Repository.Update(data.ID, usr)

	if err != nil {
		return user.User{}, err
	}

	return updatedUser, nil
}
