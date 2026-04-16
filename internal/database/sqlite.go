package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbPath string) (*sqlx.DB, error) {
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	db, err := sqlx.Connect("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxOpenConns(1)

	return db, nil
}

func RunMigrations(db *sqlx.DB, migrationPath string) error {
	content, err := os.ReadFile(migrationPath)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	return nil
}

func SeedBooks(db *sqlx.DB) error {
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM books")
	if err != nil {
		return fmt.Errorf("failed to check books count: %w", err)
	}

	if count > 0 {
		return nil
	}

	books := []map[string]interface{}{
		{
			"title":       "To Kill a Mockingbird",
			"author":      "Harper Lee",
			"price":       12.99,
			"description": "A gripping tale of racial injustice and childhood innocence in the American South.",
			"cover":       "https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=400",
			"category":    "Fiction",
			"rating":      4.8,
			"stock":       15,
		},
		{
			"title":       "1984",
			"author":      "George Orwell",
			"price":       14.99,
			"description": "A dystopian social science fiction novel and cautionary tale about totalitarianism.",
			"cover":       "https://images.unsplash.com/photo-1495446815901-a7297e633e8d?w=400",
			"category":    "Science Fiction",
			"rating":      4.7,
			"stock":       20,
		},
		{
			"title":       "Pride and Prejudice",
			"author":      "Jane Austen",
			"price":       11.99,
			"description": "A romantic novel of manners that critiques the British landed gentry at the end of the 18th century.",
			"cover":       "https://images.unsplash.com/photo-1512820790803-83ca734da794?w=400",
			"category":    "Romance",
			"rating":      4.6,
			"stock":       12,
		},
		{
			"title":       "The Hobbit",
			"author":      "J.R.R. Tolkien",
			"price":       15.99,
			"description": "A fantasy novel about the quest of home-loving Bilbo Baggins to win a share of treasure.",
			"cover":       "https://images.unsplash.com/photo-1621351183012-e2f9972dd9bf?w=400",
			"category":    "Fantasy",
			"rating":      4.9,
			"stock":       18,
		},
		{
			"title":       "The Da Vinci Code",
			"author":      "Dan Brown",
			"price":       13.99,
			"description": "A mystery thriller novel that follows symbologist Robert Langdon.",
			"cover":       "https://images.unsplash.com/photo-1543002588-bfa74002ed7e?w=400",
			"category":    "Mystery",
			"rating":      4.3,
			"stock":       10,
		},
		{
			"title":       "The Alchemist",
			"author":      "Paulo Coelho",
			"price":       12.99,
			"description": "A philosophical book about following your dreams and listening to your heart.",
			"cover":       "https://images.unsplash.com/photo-1589998059171-988d887df646?w=400",
			"category":    "Philosophy",
			"rating":      4.5,
			"stock":       14,
		},
		{
			"title":       "Sapiens",
			"author":      "Yuval Noah Harari",
			"price":       16.99,
			"description": "A brief history of humankind from the Stone Age to the modern age.",
			"cover":       "https://images.unsplash.com/photo-1532012197267-da84d127e765?w=400",
			"category":    "Non-Fiction",
			"rating":      4.7,
			"stock":       22,
		},
		{
			"title":       "Atomic Habits",
			"author":      "James Clear",
			"price":       14.99,
			"description": "An easy and proven way to build good habits and break bad ones.",
			"cover":       "https://images.unsplash.com/photo-1506880018603-83d5b814b5a6?w=400",
			"category":    "Self-Help",
			"rating":      4.8,
			"stock":       25,
		},
		{
			"title":       "The Great Gatsby",
			"author":      "F. Scott Fitzgerald",
			"price":       10.99,
			"description": "A classic American novel set in the Jazz Age on Long Island.",
			"cover":       "https://images.unsplash.com/photo-1519682337058-a94d519337bc?w=400",
			"category":    "Fiction",
			"rating":      4.4,
			"stock":       16,
		},
		{
			"title":       "Dune",
			"author":      "Frank Herbert",
			"price":       17.99,
			"description": "A science fiction novel set in the distant future amidst a huge interstellar empire.",
			"cover":       "https://images.unsplash.com/photo-1518770660439-4636190af475?w=400",
			"category":    "Science Fiction",
			"rating":      4.6,
			"stock":       13,
		},
		{
			"title":       "Harry Potter and the Sorcerer's Stone",
			"author":      "J.K. Rowling",
			"price":       13.99,
			"description": "The first novel in the Harry Potter series about a young wizard's journey.",
			"cover":       "https://images.unsplash.com/photo-1551029506-0807df4e2031?w=400",
			"category":    "Fantasy",
			"rating":      4.9,
			"stock":       30,
		},
		{
			"title":       "Thinking, Fast and Slow",
			"author":      "Daniel Kahneman",
			"price":       15.99,
			"description": "A groundbreaking tour of the mind explaining the two systems that drive the way we think.",
			"cover":       "https://images.unsplash.com/photo-1457369804613-52c61a468e7d?w=400",
			"category":    "Non-Fiction",
			"rating":      4.5,
			"stock":       19,
		},
	}

	query := `INSERT INTO books (title, author, price, description, cover, category, rating, stock) 
			  VALUES (:title, :author, :price, :description, :cover, :category, :rating, :stock)`

	for _, book := range books {
		_, err := db.NamedExec(query, book)
		if err != nil {
			return fmt.Errorf("failed to seed book: %w", err)
		}
	}

	return nil
}
