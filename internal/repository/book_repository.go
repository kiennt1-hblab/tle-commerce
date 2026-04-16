package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kiennt1/bookstore-backend/internal/models"
)

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	query := "SELECT * FROM books ORDER BY created_at DESC"
	
	err := r.db.Select(&books, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get books: %w", err)
	}
	
	return books, nil
}

func (r *BookRepository) GetByID(id int) (*models.Book, error) {
	var book models.Book
	query := "SELECT * FROM books WHERE id = ?"
	
	err := r.db.Get(&book, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get book: %w", err)
	}
	
	return &book, nil
}

func (r *BookRepository) Search(q string) ([]models.Book, error) {
	var books []models.Book
	searchTerm := "%" + strings.ToLower(q) + "%"
	
	query := `SELECT * FROM books 
			  WHERE LOWER(title) LIKE ? OR LOWER(author) LIKE ? 
			  ORDER BY created_at DESC`
	
	err := r.db.Select(&books, query, searchTerm, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to search books: %w", err)
	}
	
	return books, nil
}
