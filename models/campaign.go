package models

// Account represents details for a particular user account.
type Campaign struct {
	ID           string        `json:"id" db:"id"`
	OwnerID      string        `json:"ownerID" db:"ownerID"`
	Title        string        `json:"title" db:"title"`
	Description  string        `json:"description" db:"description"`
	Category     string        `json:"category" db:"category"`
	Participants []string    `json:"participants" db:"participants"`
	Milestones   []interface{} `json:"milestones" db:"milestones"`
	State        string        `json:"state" db:"state"`
}

func NewCampaign(data map[string]interface{}) Campaign {
	if data == nil {
		return Campaign{}
	}

	return Campaign{
		ID:           data["id"].(string),
		OwnerID:      data["ownerID"].(string),
		Title:        data["title"].(string),
		Description:  data["description"].(string),
		Category:     data["category"].(string),
		Participants: data["participants"].([]string),
		Milestones:   data["milestones"].([]interface{}),
	}
}

func (camp Campaign) Map() map[string]interface{} {
	return map[string]interface{}{
		"ID":           camp.ID,
		"OwnerID":      camp.OwnerID,
		"Title":        camp.Title,
		"Description":  camp.Description,
		"Category":     camp.Category,
		"Participants": camp.Participants,
		"Milestones":   camp.Milestones,
	}
}
