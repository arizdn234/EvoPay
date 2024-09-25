package handlers

import (
	"strconv"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/arizdn234/EvoPay/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type WithdrawalHandler struct {
	repo repositories.WithdrawalRepository
}

func NewWithdrawalHandler(repo repositories.WithdrawalRepository) *WithdrawalHandler {
	return &WithdrawalHandler{repo}
}

// CreateWithdrawal processes a withdrawal request for a user
func (h *WithdrawalHandler) CreateWithdrawal(c *fiber.Ctx) error {
	withdrawal := new(models.Withdrawal)
	if err := c.BodyParser(withdrawal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate the withdrawal amount
	if withdrawal.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Withdrawal amount must be greater than zero"})
	}

	// Call the repository to handle withdrawal logic
	if err := h.repo.Withdraw(withdrawal.UserID, withdrawal.Amount, withdrawal.Currency); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(withdrawal)
}

// GetWithdrawalByID retrieves a withdrawal by its ID
func (h *WithdrawalHandler) GetWithdrawalByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid withdrawal ID"})
	}

	withdrawal, err := h.repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Withdrawal not found"})
	}

	return c.JSON(withdrawal)
}

// GetWithdrawalsByUserID retrieves all withdrawals for a specific user by userID
func (h *WithdrawalHandler) GetWithdrawalsByUserID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("userID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	withdrawals, err := h.repo.FindByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve withdrawals"})
	}

	return c.JSON(withdrawals)
}
