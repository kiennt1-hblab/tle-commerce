package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kiennt1/bookstore-backend/internal/models"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *models.CreateOrderRequest) (*models.Order, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	orderQuery := `INSERT INTO orders (customer_name, customer_email, customer_phone, customer_address, customer_note, total) 
				   VALUES (?, ?, ?, ?, ?, ?)`
	
	result, err := tx.Exec(orderQuery,
		order.Customer.Name,
		order.Customer.Email,
		order.Customer.Phone,
		order.Customer.Address,
		order.Customer.Note,
		order.Total,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get order ID: %w", err)
	}

	itemQuery := `INSERT INTO order_items (order_id, book_id, quantity, price) VALUES (?, ?, ?, ?)`
	for _, item := range order.Items {
		_, err := tx.Exec(itemQuery, orderID, item.ID, item.Quantity, item.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to create order item: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	var createdOrder models.Order
	err = r.db.Get(&createdOrder, "SELECT * FROM orders WHERE id = ?", orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get created order: %w", err)
	}

	var items []models.OrderItem
	err = r.db.Select(&items, "SELECT * FROM order_items WHERE order_id = ?", orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order items: %w", err)
	}

	createdOrder.Items = items

	return &createdOrder, nil
}
