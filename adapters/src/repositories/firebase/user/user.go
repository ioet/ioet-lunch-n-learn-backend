package userfirebaserepository

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/ioet/ioet-lunch-n-learn-backend/api/config"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/user"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/repositories"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type UserRepository struct {
	client     *firestore.Client
	collection string
}

var _ repositories.BaseUserRepository = (*UserRepository)(nil)

func NewUserRepository(ctx context.Context) (*UserRepository, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
		return nil, err
	}

	credsJSON, err := config.FormatCredsJSON(cfg)
	if err != nil {
		log.Fatalf("Error formatting credentials: %v", err)
		return nil, err
	}

	opt := option.WithCredentialsJSON(credsJSON)
	client, err := firestore.NewClient(ctx, cfg.FirebaseProjectID, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firestore client: %v", err)
	}

	return &UserRepository{client: client, collection: "users"}, nil
}

func (u *UserRepository) GetAll() ([]user.User, error) {
	iter := u.client.Collection(u.collection).Documents(context.Background())

	var users []user.User

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return []user.User{}, fmt.Errorf("error getting users: %v", err)
		}

		var h user.User
		docErr := doc.DataTo(&h)

		if docErr != nil {
			return []user.User{}, fmt.Errorf("error getting house data: %v", docErr)
		}
		users = append(users, h)
	}

	return users, nil
}

func (u *UserRepository) GetByID(id string) (user.User, error) {
	iter := u.client.Collection(u.collection).Where("ID", "==", id).Documents(context.Background())

	var us user.User
	doc, err := iter.Next()

	if err != nil {
		return user.User{}, fmt.Errorf("error getting user: %v", err)
	}

	docErr := doc.DataTo(&us)

	if docErr != nil {
		return user.User{}, fmt.Errorf("error getting user data: %v", docErr)
	}

	return us, nil
}

func (u *UserRepository) Create(data user.User) (user.User, error) {
	_, _, err := u.client.Collection(u.collection).Add(context.Background(), data)

	if err != nil {
		return user.User{}, fmt.Errorf("error creating house: %v", err)
	}

	return data, nil
}

func (u *UserRepository) Update(ID string, data user.User) (user.User, error) {
	iter := u.client.Collection(u.collection).Where("ID", "==", ID).Documents(context.Background())

	doc, err := iter.Next()
	if err == iterator.Done {
		return user.User{}, fmt.Errorf("no user found with ID %s", ID)
	}
	if err != nil {
		return user.User{}, fmt.Errorf("error finding user with ID %s: %v", ID, err)
	}

	_, err = doc.Ref.Set(context.Background(), data, firestore.MergeAll)
	if err != nil {
		return user.User{}, fmt.Errorf("error updating user with ID %s: %v", ID, err)
	}

	return data, nil
}

// TODO: To implement if needed

func (u *UserRepository) Delete(_ string) error {
	return fmt.Errorf("not implemented")
}

func (u *UserRepository) GetAllByHouseID(houseID string) ([]user.User, error) {
	iter := u.client.Collection(u.collection).Where("HouseID", "==", houseID).Documents(context.Background())

	var users []user.User

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return []user.User{}, fmt.Errorf("error getting users: %v", err)
		}

		var h user.User
		docErr := doc.DataTo(&h)

		if docErr != nil {
			return []user.User{}, fmt.Errorf("error getting house data: %v", docErr)
		}

		users = append(users, h)
	}

	return users, nil
}
