package model

import (
	"time"
)

type Base struct {
	ID        uint      `json:"-" gorm:"primary_key" form:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
