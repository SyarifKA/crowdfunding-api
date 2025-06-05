package models

import "time"

type User struct {
	ID        string
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
