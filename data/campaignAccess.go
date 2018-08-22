package data

import (
	"github.com/ecclesia-dev/campaign-service/models"
)

type CampaignAccess interface {
 		CreateCampaign(campaign models.Campaign) (string, error)
 		GetCampaignByID(id string) (models.Campaign)
 		GetCampaignByTitle(title string) (models.Campaign)
 		GetAllCampaigns() ([]models.Campaign)
 		RemoveCampaign(id string) error
 		UpdateCampaign(id string, updates map[string]interface{}) error
 		CreateAction(action models.Action) (string, error)
 		RemoveAction(id string) error
		GetAllActions(id string) error
 		EnrollUserInCampaign(userID string, campaignID string) (string, error)
 		RemoveUserFromCampaign(userID string, campaignID string) error
}