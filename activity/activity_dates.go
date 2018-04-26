package activity

import (
	"time"
)

type ActivityDate struct {
	ID        int64     `json:"id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
