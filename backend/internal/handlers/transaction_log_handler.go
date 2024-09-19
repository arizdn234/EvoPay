package handlers

import (
	"strconv"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/arizdn234/EvoPay/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type TransactionLogHandler struct {
	repo repositories.TransactionLogRepository
}

func NewTransactionLogHandler(repo repositories.TransactionLogRepository) *TransactionLogHandler {
	return &TransactionLogHandler{repo}
}

// CreateTransactionLog creates a new transaction log
func (h *TransactionLogHandler) CreateTransactionLog(c *fiber.Ctx) error {
	log := new(models.TransactionLog)

	// Parse the request body into the log model
	if err := c.BodyParser(log); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Save the log to the database
	if err := h.repo.Create(log); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create transaction log"})
	}

	return c.Status(fiber.StatusCreated).JSON(log)
}

// GetTransactionLogsByTransactionID retrieves logs by transaction ID
func (h *TransactionLogHandler) GetTransactionLogsByTransactionID(c *fiber.Ctx) error {
	// Parse transaction ID from URL parameter
	transactionID, err := strconv.Atoi(c.Params("transactionID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid transaction ID"})
	}

	// Retrieve logs from the repository
	logs, err := h.repo.FindByTransactionID(uint(transactionID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve transaction logs"})
	}

	return c.JSON(logs)
}

// GetTransactionLogsByUserID retrieves logs by user ID
func (h *TransactionLogHandler) GetTransactionLogsByUserID(c *fiber.Ctx) error {
	// Parse user ID from URL parameter
	userID, err := strconv.Atoi(c.Params("userID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Retrieve logs from the repository
	logs, err := h.repo.FindByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve transaction logs"})
	}

	return c.JSON(logs)
}
