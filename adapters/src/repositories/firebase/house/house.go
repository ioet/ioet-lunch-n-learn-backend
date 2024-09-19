package housefirebaserepository

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/ioet/ioet-lunch-n-learn-backend/api/config"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/house"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/repositories"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type HouseRepository struct {
	client     *firestore.Client
	collection string
}

var _ repositories.Base[house.House] = (*HouseRepository)(nil)

func NewHouseRepository(ctx context.Context) (*HouseRepository, error) {
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

	return &HouseRepository{client: client, collection: "houses"}, nil
}

func (u *HouseRepository) GetAll() ([]house.House, error) {
	iter := u.client.Collection(u.collection).Documents(context.Background())

	var houses []house.House

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return []house.House{}, fmt.Errorf("error getting houses: %v", err)
		}

		var h house.House
		docErr := doc.DataTo(&h)

		if docErr != nil {
			return []house.House{}, fmt.Errorf("error getting house data: %v", docErr)
		}

		houses = append(houses, h)
	}

	return houses, nil
}

func (u *HouseRepository) GetByID(id string) (house.House, error) {
	iter := u.client.Collection(u.collection).Where("ID", "==", id).Documents(context.Background())

	var h house.House
	doc, err := iter.Next()

	if err != nil {
		return house.House{}, fmt.Errorf("error getting house: %v", err)
	}

	docErr := doc.DataTo(&h)

	if docErr != nil {
		return house.House{}, fmt.Errorf("error getting house data: %v", docErr)
	}
	return h, nil
}

func (u *HouseRepository) Create(data house.House) (house.House, error) {
	_, _, err := u.client.Collection(u.collection).Add(context.Background(), data)

	if err != nil {
		return house.House{}, fmt.Errorf("error creating house: %v", err)
	}

	return data, nil
}

// TODO: To implement if needed

func (u *HouseRepository) Update(_ string, _ house.House) (house.House, error) {
	return house.House{}, fmt.Errorf("not implemented")
}

func (u *HouseRepository) Delete(_ string) error {
	return fmt.Errorf("not implemented")
}
