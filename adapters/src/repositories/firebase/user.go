package firebase

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/google/uuid"
	"github.com/ioet/ioet-lunch-n-learn-backend/api/config"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/user"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/repositories"
	"google.golang.org/api/option"
)

type UserRepository struct {
	client  *db.Client
	eventID uuid.UUID
}

var _ repositories.Base[user.User] = (*UserRepository)(nil)

func NewUserRepository(ctx context.Context, eventID uuid.UUID) (*UserRepository, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	credsJSON, err := config.FormatCredsJSON(cfg)

	if err != nil {
		log.Fatalf("Error formatting credentials: %v", err)
		return nil, err
	}

	opt := option.WithCredentialsJSON(credsJSON)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	client, err := app.DatabaseWithURL(ctx, cfg.FirebaseAppURL)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase Realtime Database client: %v", err)
	}

	return &UserRepository{client: client, eventID: eventID}, nil
}

func (u *UserRepository) SetEventPath() *db.Ref {
	return u.client.NewRef(fmt.Sprintf("attendance/users_%s", u.eventID.String()))
}

func (u *UserRepository) GetAll() ([]user.User, error) {
	var users []user.User
	ref := u.SetEventPath()
	if err := ref.Get(context.Background(), &users); err != nil {
		log.Printf("Error getting users: %v", err)
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) GetByID(ID uuid.UUID) (user.User, error) {
	var userData user.User
	ref := u.SetEventPath().Child(ID.String())
	if err := ref.Get(context.Background(), &userData); err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return user.User{}, err
	}
	return userData, nil
}

func (u *UserRepository) Create(data user.User) (user.User, error) {
	ref, err := u.SetEventPath().Push(context.Background(), nil)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return user.User{}, err
	}

	data.ID = uuid.New()
	if err := ref.Set(context.Background(), data); err != nil {
		log.Printf("Error creating user: %v", err)
		return user.User{}, err
	}
	return data, nil
}

func (u *UserRepository) Update(ID uuid.UUID, data user.User) (user.User, error) {
	ref := u.SetEventPath().Child(ID.String())
	dataMap := map[string]interface{}{
		"id":        data.ID.String(),
		"name":      data.Name,
		"email":     data.Email,
		"photoUrl":  data.PhotoURL,
		"houseId":   data.HouseID.String(),
		"createdAt": data.CreatedAt.Format(time.RFC3339),
		"updatedAt": data.UpdatedAt.Format(time.RFC3339),
	}
	if err := ref.Update(context.Background(), dataMap); err != nil {
		log.Printf("Error updating user: %v", err)
		return user.User{}, err
	}
	return data, nil
}

func (u *UserRepository) Delete(ID uuid.UUID) error {
	ref := u.SetEventPath().Child(ID.String())
	if err := ref.Delete(context.Background()); err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}

func (u *UserRepository) CheckExist(ID uuid.UUID) (bool, error) {
	ref := u.SetEventPath().Child(ID.String())
	var userData user.User
	if err := ref.Get(context.Background(), &userData); err != nil {
		log.Printf("Error checking if user exists: %v", err)
		return false, err
	}
	return userData.ID != uuid.Nil, nil
}
