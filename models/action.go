package models


type Action struct {
	ID           string        `json:"id" db:"id"`
	Description  string        `json:"description" db:"description"`
	Milestones   []interface{} `json:"milestones" db:"milestones"`
	Location        string        `json:"Location" db:"Location"`
}

