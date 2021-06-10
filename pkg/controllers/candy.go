package controllers

import (
	"database/sql"
	"github.com/wisdomfusion/candy-code-box/pkg/models"
	"strconv"
	"time"
)

type CandyModel struct {
	DB *sql.DB
}

func (m *CandyModel) Show(id int) (*models.Candy, error) {
	return nil, nil
}

func (m *CandyModel) Store(title, candy, expires string) (int, error) {
	datetime := time.Now()
	dt := datetime.Format(time.RFC3339)
	createdAt, updatedAt := dt, dt
	expireDays, err := strconv.Atoi(expires)
	if err != nil {
		return 0, err
	}
	expiredAt := "NULL"
	if expireDays > 0 {
		expiredAt = datetime.AddDate(0, 0, expireDays).Format(time.RFC3339)
	}

	stmt := `
INSERT INTO candies
    (title, candy, created_at, updated_at, expired_at)
VALUES
       (?, ?, ?, ?, ?)
`
	result, err := m.DB.Exec(stmt, title, candy, createdAt, updatedAt, expiredAt)
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
	return nil, nil
}
