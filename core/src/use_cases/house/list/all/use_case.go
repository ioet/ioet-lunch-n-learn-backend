package houselistingusecase

import (
	housefirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/house"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/house"
)

type HouseListingUseCase struct {
	Repository housefirebaserepository.HouseRepository
}

func NewHouseListingUseCase(repository housefirebaserepository.HouseRepository) *HouseListingUseCase {
	return &HouseListingUseCase{Repository: repository}
}

func (u *HouseListingUseCase) Execute() ([]house.House, error) {
	houses, err := u.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return houses, nil
}
