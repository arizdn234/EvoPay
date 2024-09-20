package handlers

import (
	"strconv"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/arizdn234/EvoPay/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	repo repositories.TransactionRepository
}

func NewTransactionHandler(repo repositories.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{repo}
}

// CreateTransaction creates a new transaction
func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	transaction := new(models.Transaction)

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.repo.Create(transaction); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create transaction"})
	}

	return c.Status(fiber.StatusCreated).JSON(transaction)
}

// GetAllTransactions retrieves all transactions
func (h *TransactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	transactions, err := h.repo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve transactions"})
	}

	return c.JSON(transactions)
}

// GetTransactionByID retrieves a transaction by ID
func (h *TransactionHandler) GetTransactionByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid transaction ID"})
	}

	transaction, err := h.repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Transaction not found"})
	}

	return c.JSON(transaction)
}

// UpdateTransaction updates a transaction
func (h *TransactionHandler) UpdateTransaction(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid transaction ID"})
	}

	transaction := new(models.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	transaction.ID = uint(id)
	if err := h.repo.Update(transaction); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update transaction"})
	}

	return c.JSON(transaction)
}

// // DeleteTransaction deletes a transaction
// func (h *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid transaction ID"})
// 	}

// 	if err := h.repo.Delete(uint(id)); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete transaction"})
// 	}

// 	return c.SendStatus(fiber.StatusNoContent)
// }
