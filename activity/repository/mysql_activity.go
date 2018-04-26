package repository

import (
	"database/sql"
	"fmt"
	"log"

	models "github.com/skiple/skiple-api/activity"
)

type mysqlActivityRepository struct {
	Conn *sql.DB
}

func newMysqlActivityRepository(Conn *sql.DB) ActivityRepository {
	return &mysqlActivityRepository{Conn}
}

func (m *mysqlActivityRepository) fetch(query string, args ...interface{}) ([]*models.Activity, error) {
	rows, err := m.Conn.Query(query, args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Activity, 0)
	for rows.Next() {
		t := new(models.Activity)
		activityDates := make([]models.ActivityDate, 0)
		err = rows.Scan(
			&t.ID,
			&t.ActivityName,
			&t.Slug,
			&t.HostName,
			&t.HostProfile,
			&t.Duration,
			&t.Description,
			&t.MaxParticipants,
			&t.Price,
			&t.Provide,
			&t.Location,
			&t.Itinerary,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		t.ActivityDates = activityDates
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlActivityRepository) Fetch(cursor string, num int64) ([]*models.Activity, error) {
	q := `
            SELECT
                id, activity_name, slug, host_name, host_profile,
                duration, description, max_participants,
                price, provide, location, itinerary,
                updated_at, created_at
            FROM
                tb_activity
            WHERE id > ? LIMIT ?`

	return m.fetch(q, cursor, num)
}

func (m *mysqlActivityRepository) GetByID(id int64) (*models.Activity, error) {
	q := `
            SELECT
                id, activity_name, slug, host_name, host_profile,
                duration, description, max_participants,
                price, provide, location, itinerary,
                updated_at, created_at
            FROM
                tb_activity
            WHERE id = ?`

	list, err := m.fetch(q, id)
	if err != nil {
		return nil, err
	}

	a := &models.Activity{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.ErrorNotFound
	}

	return a, nil
}

func (m *mysqlActivityRepository) GetBySlug(slug string) (*models.Activity, error) {
	q := `
            SELECT
                id, activity_name, slug, host_name, host_profile,
                duration, description, max_participants,
                price, provide, location, itinerary,
                updated_at, created_at
            FROM
                tb_activity
            WHERE slug = ?`

	list, err := m.fetch(q, slug)
	if err != nil {
		return nil, err
	}

	a := &models.Activity{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.ErrorNotFound
	}

	return a, nil
}

func (m *mysqlActivityRepository) Store(a *models.Activity) (int64, error) {
	q := `
            INSERT INTO tb_activity (
                activity_name, slug, host_name, host_profile,
                duration, description, max_participants,
                price, provide, location, itinerary
            )
            VALUES (
                ?, ?, ?, ?,
                ?, ?, ?,
                ?, ?, ?, ?
            )`

	stmt, err := m.Conn.Prepare(q)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(a.ActivityName, a.HostName, a.HostProfile,
		a.Duration, a.Description, a.MaxParticipants,
		a.Price, a.Provide, a.Location, a.Itinerary)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (m *mysqlActivityRepository) Delete(id int64) (bool, error) {
	q := `
            DELETE FROM tb_activity
            WHERE id = ?`

	stmt, err := m.Conn.Prepare(q)
	if err != nil {
		return false, err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAffected != 1 {
		err = fmt.Errorf("weird behavior: Total affected (deleted) rows: %d", rowsAffected)
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

func (m *mysqlActivityRepository) Update(a *models.Activity) (*models.Activity, error) {
	q := `
            UPDATE tb_activity
            SET
                activity_name = ?, slug = ?, host_name = ?, host_profile = ?,
                duration = ?, description = ?, max_participants = ?,
                price = ?, provide = ?, location = ?, itinerary = ?
            WHERE
                id = ?`

	stmt, err := m.Conn.Prepare(q)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(a.ActivityName, a.Slug, a.HostName, a.HostProfile,
		a.Duration, a.Description, a.MaxParticipants,
		a.Price, a.Provide, a.Location, a.Itinerary, a.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected != 1 {
		err = fmt.Errorf("weird behavior: Total affected (updated) rows: %d", rowsAffected)
		log.Fatal(err)
		return nil, err
	}

	return a, nil
}
