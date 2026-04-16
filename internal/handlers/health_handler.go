package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type HealthHandler struct {
	db *sqlx.DB
}

func NewHealthHandler(db *sqlx.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

func (h *HealthHandler) HealthCheck(c *fiber.Ctx) error {
	if err := h.db.Ping(); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status":  "unhealthy",
			"message": "Database connection failed",
			"time":    time.Now().Format(time.RFC3339),
		})
	}

	var bookCount int
	err := h.db.Get(&bookCount, "SELECT COUNT(*) FROM books")
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status":  "unhealthy",
			"message": "Database query failed",
			"time":    time.Now().Format(time.RFC3339),
		})
	}

	return c.JSON(fiber.Map{
		"status":     "healthy",
		"message":    "Server is running",
		"time":       time.Now().Format(time.RFC3339),
		"database":   "connected",
		"book_count": bookCount,
	})
}
