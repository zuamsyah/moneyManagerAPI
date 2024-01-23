package models

import "time"

type Category struct {
	ID          int           `json:"id" gorm:"primaryKey"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Type        string        `json:"type"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Transaction []Transaction `gorm:"foreignKey:CategoryID"`
}
