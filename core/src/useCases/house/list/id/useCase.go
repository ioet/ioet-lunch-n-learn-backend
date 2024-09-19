package houselistingbyidusecase

import (
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/house"
	houseorchestrator "github.com/ioet/ioet-lunch-n-learn-backend/core/src/orchestrators/House"
)

type HouseListingByIdUseCase struct {
	Orchestrator houseorchestrator.HouseOrquestrator
}

func NewHouseListingByIdUseCase(orchestrator houseorchestrator.HouseOrquestrator) *HouseListingByIdUseCase {
	return &HouseListingByIdUseCase{Orchestrator: orchestrator}
}

func (u *HouseListingByIdUseCase) Execute(id string) (house.House, error) {
	house, err := u.Orchestrator.GetHouseByID(id)

	if err != nil {
		return house, err
	}

	return house, nil
}
