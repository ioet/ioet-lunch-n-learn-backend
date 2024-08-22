package housecreationusecase

import (
	housefirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/house"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/house"
)

type HouseCreationUseCase struct {
	Repository housefirebaserepository.HouseRepository
}

func NewHouseCreationUseCase(repository housefirebaserepository.HouseRepository) *HouseCreationUseCase {
	return &HouseCreationUseCase{Repository: repository}
}

func (u *HouseCreationUseCase) Execute(data house.House) (house.House, error) {
	createdHouse, err := u.Repository.Create(data)

	if err != nil {
		return house.House{}, err
	}

	return createdHouse, nil
}
