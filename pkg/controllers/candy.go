package controllers

import (
	"database/sql"
	"errors"
	"github.com/wisdomfusion/candy-code-box/pkg/models"
	"strconv"
	"time"
)

type CandyModel struct {
	DB *sql.DB
}

func (m *CandyModel) Show(id int) (*models.Candy, error) {
	query := `
SELECT id, title, candy, created_at, updated_at, expired_at
FROM candies
WHERE id=? AND deleted_at IS NULL
`
	c := &models.Candy{}
	err := m.DB.QueryRow(query, id).Scan(
		&c.Id,
		&c.Title,
		&c.Candy,
		&c.CreatedAt,
		&c.UpdatedAt,
		&c.ExpiredAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return c, nil
}

func (m *CandyModel) Store(title, candy, expireDays string) (int, error) {
	datetime := time.Now()
	dt := datetime.Format(time.RFC3339)
	createdAt, updatedAt := dt, dt

	days, err := strconv.Atoi(expireDays)
	if err != nil {
		return 0, err
	}

	expiredAt := "NULL"
	if days > 0 {
		expiredAt = datetime.AddDate(0, 0, days).Format(time.RFC3339)
	}

	query := `
INSERT INTO candies (title, candy, created_at, updated_at, expired_at)
VALUES (?, ?, ?, ?, ?)
`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(title, candy, createdAt, updatedAt, expiredAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *CandyModel) Latest() ([]*models.Candy, error) {
	query := `
SELECT id, title, candy, created_at, updated_at, expired_at
FROM candies
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT 10
`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	candies := []*models.Candy{}

	for rows.Next() {
		c := &models.Candy{}

		err = rows.Scan(
			&c.Id,
			&c.Title,
			&c.Candy,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.ExpiredAt,
		)
		if err != nil {
			return nil, err
		}

		candies = append(candies, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return candies, nil
}
