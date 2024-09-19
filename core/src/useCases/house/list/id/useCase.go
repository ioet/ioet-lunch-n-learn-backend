package houselistingbyidusecase

import (
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/house"
	houseorchestrator "github.com/ioet/ioet-lunch-n-learn-backend/core/src/orchestrators/house"
)

type HouseListingByIDUseCase struct {
	Orchestrator houseorchestrator.HouseOrquestrator
}

func NewHouseListingByIDUseCase(orchestrator houseorchestrator.HouseOrquestrator) *HouseListingByIDUseCase {
	return &HouseListingByIDUseCase{Orchestrator: orchestrator}
}

func (u *HouseListingByIDUseCase) Execute(id string) (house.House, error) {
	house, err := u.Orchestrator.GetHouseByID(id)

	if err != nil {
		return house, err
	}

	return house, nil
}
