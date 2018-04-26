package activity

import (
	"time"
)

type Activity struct {
	ID              int64          `json:"id"`
	ActivityName    string         `json:"activity_name"`
	Slug            string         `json:"slug"`
	HostName        string         `json:"host_name"`
	HostProfile     string         `json:"host_profile"`
	Duration        int64          `json:"duration"`
	Description     string         `json:"description"`
	MaxParticipants int64          `json:"max_participants"`
	Price           int64          `json:"price"`
	Provide         string         `json:"provide"`
	Location        string         `json:"location"`
	Itinerary       string         `json:"itinerary"`
	ActivityDates   []ActivityDate `json:"activity_dates"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}
