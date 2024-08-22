package repositoryfactories

import (
	"context"

	lnlfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/lunchNLearn"
)

func LnLFirebaseRepository() *lnlfirebaserepository.LnLRepository {
	firebaseRepository, err := lnlfirebaserepository.NewLnLRepository(context.Background())

	if err != nil {
		panic("Error initializing the Lunch-n-learn repository: " + err.Error())
	}

	return firebaseRepository
}
