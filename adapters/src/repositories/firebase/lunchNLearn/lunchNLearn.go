package lnlfirebaserepository

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/ioet/ioet-lunch-n-learn-backend/api/config"
	lunchnlearn "github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/lunchNLearn"
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/repositories"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type LnLRepository struct {
	client     *firestore.Client
	collection string
}

var _ repositories.Base[lunchnlearn.LunchNLearn] = (*LnLRepository)(nil)

func NewLnLRepository(ctx context.Context) (*LnLRepository, error) {
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

	return &LnLRepository{client: client, collection: "lunch-n-learn"}, nil
}

func (u *LnLRepository) GetAll() ([]lunchnlearn.LunchNLearn, error) {
	iter := u.client.Collection(u.collection).Documents(context.Background())

	var lunchnlearns []lunchnlearn.LunchNLearn

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return []lunchnlearn.LunchNLearn{}, fmt.Errorf("error getting lunch-n-learns: %v", err)
		}

		var h lunchnlearn.LunchNLearn
		docErr := doc.DataTo(&h)

		if docErr != nil {
			return []lunchnlearn.LunchNLearn{}, fmt.Errorf("error getting lunch-n-learn: %v", docErr)
		}

		lunchnlearns = append(lunchnlearns, h)
	}

	return lunchnlearns, nil
}

func (u *LnLRepository) GetByID(id string) (lunchnlearn.LunchNLearn, error) {
	iter := u.client.Collection(u.collection).Where("ID", "==", id).Documents(context.Background())

	var lnl lunchnlearn.LunchNLearn
	doc, err := iter.Next()

	if err != nil {
		return lunchnlearn.LunchNLearn{}, fmt.Errorf("error getting house: %v", err)
	}

	docErr := doc.DataTo(&lnl)

	if docErr != nil {
		return lunchnlearn.LunchNLearn{}, fmt.Errorf("error getting house data: %v", docErr)
	}

	return lnl, nil
}

func (u *LnLRepository) Create(data lunchnlearn.LunchNLearn) (lunchnlearn.LunchNLearn, error) {
	_, _, err := u.client.Collection(u.collection).Add(context.Background(), data)

	if err != nil {
		return lunchnlearn.LunchNLearn{}, fmt.Errorf("error creating house: %v", err)
	}

	return data, nil
}

func (u *LnLRepository) Update(ID string, data lunchnlearn.LunchNLearn) (lunchnlearn.LunchNLearn, error) {
	iter := u.client.Collection(u.collection).Where("ID", "==", ID).Documents(context.Background())

	doc, err := iter.Next()
	if err == iterator.Done {
		return lunchnlearn.LunchNLearn{}, fmt.Errorf("no lunch-n-learn found with ID %s", ID)
	}
	if err != nil {
		return lunchnlearn.LunchNLearn{}, fmt.Errorf("error finding lunch-n-learn with ID %s: %v", ID, err)
	}

	_, err = doc.Ref.Set(context.Background(), data, firestore.MergeAll)
	if err != nil {
		return lunchnlearn.LunchNLearn{}, fmt.Errorf("error updating lunch-n-learn with ID %s: %v", ID, err)
	}

	return data, nil
}

// TODO: To implement if needed

func (u *LnLRepository) Delete(_ string) error {
	return fmt.Errorf("not implemented")
}
