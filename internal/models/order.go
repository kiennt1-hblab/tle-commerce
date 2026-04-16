package models

import "time"

type Order struct {
	ID              int         `json:"id" db:"id"`
	CustomerName    string      `json:"customerName" db:"customer_name"`
	CustomerEmail   string      `json:"customerEmail" db:"customer_email"`
	CustomerPhone   string      `json:"customerPhone" db:"customer_phone"`
	CustomerAddress string      `json:"customerAddress" db:"customer_address"`
	CustomerNote    string      `json:"customerNote" db:"customer_note"`
	Total           float64     `json:"total" db:"total"`
	CreatedAt       time.Time   `json:"createdAt" db:"created_at"`
	Items           []OrderItem `json:"items" db:"-"`
}

type OrderItem struct {
	ID       int     `json:"id" db:"id"`
	OrderID  int     `json:"orderId" db:"order_id"`
	BookID   int     `json:"bookId" db:"book_id"`
	Quantity int     `json:"quantity" db:"quantity"`
	Price    float64 `json:"price" db:"price"`
}

type CreateOrderRequest struct {
	Customer struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
		Phone   string `json:"phone" validate:"required"`
		Address string `json:"address" validate:"required"`
		Note    string `json:"note"`
	} `json:"customer" validate:"required"`
	Items []struct {
		ID       int     `json:"id" validate:"required"`
		Quantity int     `json:"quantity" validate:"required,min=1"`
		Price    float64 `json:"price" validate:"required,min=0"`
	} `json:"items" validate:"required,min=1"`
	Total float64 `json:"total" validate:"required,min=0"`
}
