package models

import "time"

type Book struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Author      string    `json:"author" db:"author"`
	Price       float64   `json:"price" db:"price"`
	Description string    `json:"description" db:"description"`
	Cover       string    `json:"cover" db:"cover"`
	Category    string    `json:"category" db:"category"`
	Rating      float64   `json:"rating" db:"rating"`
	Stock       int       `json:"stock" db:"stock"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}
