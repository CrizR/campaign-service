package data

import (
	"context"

	log "github.com/sirupsen/logrus"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/ecclesia-dev/campaign-service/models"
	"google.golang.org/api/option"
)

type Firebase struct {
	app    *firebase.App
	client *firestore.Client
	auth   *auth.Client
}

func NewFirebase() CampaignAccess {
	opt := option.WithCredentialsFile("../keys/ecclesia-firebase-key.json")
	// TODO: set FIREBASE_CONFIG as an envornment variable so config can be passed in as nil.
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize firebase app: %v", err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize Auth: %v", err)
	}

	log.Info("Firebase connection initialized.")
	return Firebase{app: app, client: client, auth: auth}
}

func (fb Firebase) CreateCampaign(account models.Campaign) (string, error) {
	return "", nil
}

func (fb Firebase) GetAllCampaigns() ([]models.Campaign, error) {
	return []models.Campaign, nil
}