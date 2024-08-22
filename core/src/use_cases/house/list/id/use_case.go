package houselistingbyidusecase

import (
	housefirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/house"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/house"
)

type HouseListingByIdUseCase struct {
	Repository housefirebaserepository.HouseRepository
}

func NewHouseListingByIdUseCase(repository housefirebaserepository.HouseRepository) *HouseListingByIdUseCase {
	return &HouseListingByIdUseCase{Repository: repository}
}

func (u *HouseListingByIdUseCase) Execute(id string) (house.House, error) {
	house, err := u.Repository.GetByID(id)

	if err != nil {
		return house, err
	}

	return house, nil
}
