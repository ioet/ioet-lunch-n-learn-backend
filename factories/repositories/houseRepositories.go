package repositoryfactories

import (
	"context"

	housefirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/house"
)

func HouseFirebaseRepository() *housefirebaserepository.HouseRepository {
	firebaseRepository, err := housefirebaserepository.NewHouseRepository(context.Background())

	if err != nil {
		panic("Error initializing the House repository: " + err.Error())
	}

	return firebaseRepository
}
