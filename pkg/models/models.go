package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Candy struct {
	Id int
	Title string
	Candy string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt time.Time
	DeletedAt time.Time
}
