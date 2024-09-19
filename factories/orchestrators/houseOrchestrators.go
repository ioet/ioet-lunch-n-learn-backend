package orchestratorfactories

import (
	houseorchestrator "github.com/ioet/ioet-lunch-n-learn-backend/core/src/orchestrators/House"
	repositoryfactories "github.com/ioet/ioet-lunch-n-learn-backend/factories/repositories"
)

func HouseFirebaseOrchestrator() *houseorchestrator.HouseOrquestrator {
	userRepository := repositoryfactories.UserFirebaseRepository()
	houseRepository := repositoryfactories.HouseFirebaseRepository()

	orchestrator := houseorchestrator.NewHouseOrquestrator(userRepository, houseRepository)

	return orchestrator
}
