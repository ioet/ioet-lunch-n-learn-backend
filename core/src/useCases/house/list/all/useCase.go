package houselistingusecase

import (
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/house"
	houseorchestrator "github.com/ioet/ioet-lunch-n-learn-backend/core/src/orchestrators/House"
)

type HouseListingUseCase struct {
	Orchestrator houseorchestrator.HouseOrquestrator
}

func NewHouseListingUseCase(orchestrator houseorchestrator.HouseOrquestrator) *HouseListingUseCase {
	return &HouseListingUseCase{Orchestrator: orchestrator}
}

func (u *HouseListingUseCase) Execute() ([]house.House, error) {
	houses, err := u.Orchestrator.GetAllHouses()

	if err != nil {
		return nil, err
	}

	return houses, nil
}
