package models

import "time"

type Transaction struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Type        string `json:"type"`
	CategoryID  int    `json:"category_id" gorm:"index"`
	Category    Category
	Month       int       `json:"month"`
	Nominal     int       `json:"nominal"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResponseTransactions struct {
	Status  int           `json:"status"`
	Data    []Transaction `json:"data"`
	Message string        `json:"message"`
}
