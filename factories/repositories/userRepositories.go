package repositoryfactories

import (
	"context"

	userfirebaserepository "github.com/ioet/ioet-lunch-n-learn-backend/adapters/src/repositories/firebase/user"
)

func UserFirebaseRepository() *userfirebaserepository.UserRepository {
	firebaseRepository, err := userfirebaserepository.NewUserRepository(context.Background())

	if err != nil {
		panic("Error initializing the User repository: " + err.Error())
	}

	return firebaseRepository
}
