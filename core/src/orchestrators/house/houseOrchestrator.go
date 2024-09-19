package houseorchestrator

import (
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/house"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/repositories"
)

type HouseOrquestrator struct {
	userRepository  repositories.BaseUserRepository
	houseRepository repositories.Base[house.House]
}

func NewHouseOrquestrator(userRepository repositories.BaseUserRepository, houseRepository repositories.Base[house.House]) *HouseOrquestrator {
	return &HouseOrquestrator{
		userRepository:  userRepository,
		houseRepository: houseRepository,
	}
}

func (u *HouseOrquestrator) GetAllHouses() ([]house.House, error) {
	houses, err := u.houseRepository.GetAll()

	if err != nil {
		return []house.House{}, err
	}

	for i, _house := range houses {
		houseUsers, err := u.userRepository.GetAllByHouseID(_house.ID)

		if err != nil {
			return []house.House{}, err
		}

		pointsCounter := 0

		for _, _user := range houseUsers {
			pointsCounter += _user.Points
		}

		houses[i].ChangePoints(pointsCounter)
	}

	return houses, nil
}

func (u *HouseOrquestrator) GetHouseByID(id string) (house.House, error) {
	_house, err := u.houseRepository.GetByID(id)

	if err != nil {
		return house.House{}, err
	}

	houseUsers, err := u.userRepository.GetAllByHouseID(_house.ID)

	if err != nil {
		return house.House{}, err
	}

	pointsCounter := 0

	for _, _user := range houseUsers {
		pointsCounter += _user.Points
	}

	_house.ChangePoints(pointsCounter)

	return _house, nil
}
