package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kiennt1/bookstore-backend/internal/database"
	"github.com/kiennt1/bookstore-backend/internal/handlers"
	"github.com/kiennt1/bookstore-backend/internal/repository"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
	port := getEnv("PORT", "3000")
	dbPath := getEnv("DB_PATH", "./data/bookstore.db")
	allowedOrigins := getEnv("ALLOWED_ORIGINS", "http://localhost:5173")

	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	if err := database.RunMigrations(db, "./migrations/001_init.sql"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	if err := database.SeedBooks(db); err != nil {
		log.Fatalf("Failed to seed books: %v", err)
	}

	bookRepo := repository.NewBookRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	bookHandler := handlers.NewBookHandler(bookRepo)
	orderHandler := handlers.NewOrderHandler(orderRepo)
	healthHandler := handlers.NewHealthHandler(db)

	app := fiber.New(fiber.Config{
		AppName: "BookStore API",
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	app.Get("/health", healthHandler.HealthCheck)

	api := app.Group("/api")

	api.Get("/books", bookHandler.GetBooks)
	api.Get("/books/search", bookHandler.SearchBooks)
	api.Get("/books/:id", bookHandler.GetBookByID)
	api.Post("/orders", orderHandler.CreateOrder)

	go func() {
		addr := fmt.Sprintf(":%s", port)
		log.Printf("Server starting on http://localhost%s", addr)
		if err := app.Listen(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
