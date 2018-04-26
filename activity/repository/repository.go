package repository

import models "github.com/skiple/skiple-api/activity"

// ActivityRepository is the main repository interface for Activity model
type ActivityRepository interface {
	Fetch(cursor string, num int64) ([]*models.Activity, error)
	GetByID(id int64) (*models.Activity, error)
	GetBySlug(slug string) (*models.Activity, error)
	Update(activity *models.Activity) (*models.Activity, error)
	Store(a *models.Activity) (int64, error)
	Delete(id int64) (bool, error)
}
