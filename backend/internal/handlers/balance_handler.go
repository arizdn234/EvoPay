package handlers

import (
	"strconv"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/arizdn234/EvoPay/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type BalanceHandler struct {
	repo repositories.BalanceRepository
}

func NewBalanceHandler(repo repositories.BalanceRepository) *BalanceHandler {
	return &BalanceHandler{repo}
}

func (h *BalanceHandler) CreateBalance(c *fiber.Ctx) error {
	balance := new(models.Balance)

	if err := c.BodyParser(balance); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.repo.Create(balance); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create balance"})
	}

	return c.Status(fiber.StatusCreated).JSON(balance)
}

func (h *BalanceHandler) GetBalanceByUserID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("userID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	balance, err := h.repo.FindByUserID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Balance not found"})
	}

	return c.JSON(balance)
}

func (h *BalanceHandler) UpdateBalance(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("userID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	balance := new(models.Balance)
	if err := c.BodyParser(balance); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	balance.UserID = uint(id)
	if err := h.repo.Update(balance); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update balance"})
	}

	return c.JSON(balance)
}

func (h *BalanceHandler) AddToBalance(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("userID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var payload struct {
		Amount float64 `json:"amount"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.repo.AddToBalance(uint(id), payload.Amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add to balance"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *BalanceHandler) SubtractFromBalance(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("userID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var payload struct {
		Amount float64 `json:"amount"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.repo.SubtractFromBalance(uint(id), payload.Amount); err != nil {
		if err.Error() == "insufficient funds" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Insufficient funds"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to subtract from balance"})
	}

	return c.SendStatus(fiber.StatusOK)
}
