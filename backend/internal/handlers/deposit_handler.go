package handlers

import (
	"strconv"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/arizdn234/EvoPay/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type DepositHandler struct {
	repo repositories.DepositRepository
}

func NewDepositHandler(repo repositories.DepositRepository) *DepositHandler {
	return &DepositHandler{repo}
}

// CreateDeposit creates a new deposit for a user
func (h *DepositHandler) CreateDeposit(c *fiber.Ctx) error {
	deposit := new(models.Deposit)
	if err := c.BodyParser(deposit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Check if the amount is valid
	if deposit.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Deposit amount must be greater than zero"})
	}

	// Call the repository to handle deposit logic
	if err := h.repo.Deposit(deposit.UserID, deposit.Amount, deposit.Currency); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(deposit)
}

// GetDepositByID retrieves a deposit by its ID
func (h *DepositHandler) GetDepositByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid deposit ID"})
	}

	deposit, err := h.repo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Deposit not found"})
	}

	return c.JSON(deposit)
}

// GetDepositsByUserID retrieves all deposits for a specific user by userID
func (h *DepositHandler) GetDepositsByUserID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("userID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	deposits, err := h.repo.FindByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve deposits"})
	}

	return c.JSON(deposits)
}
