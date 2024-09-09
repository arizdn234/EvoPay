package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/arizdn234/EvoPay/internal/middleware"
	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/arizdn234/EvoPay/internal/redis"
	"github.com/arizdn234/EvoPay/internal/repository"
	"github.com/arizdn234/EvoPay/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserRepository *repository.UserRepository
}

func NewUserHandler(ur *repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepository: ur}
}

func (uh *UserHandler) Welcome(c *fiber.Ctx) error {
	info := `
	━─━────༺EvoPay༻────━─━
	Here is the EvoPay backend service documentation:

	Routes Available:
	- GET    /                  : Welcome & backend routes documentation
	- GET    /users             : Get all users
	- POST   /users             : Create a new user (using create method)
	- GET    /users/{id}        : Get user by ID
	- PUT    /users/{id}        : Update user by ID
	- DELETE /users/{id}        : Delete user by ID
	- POST   /login             : User login
	- POST   /register          : Register new user (using register method)
	`

	return c.SendString(info)
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// validate user data
	if err := uh.validateUser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := utils.HashPassword(user.Password + user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	user.Password = hashedPassword

	if err := uh.UserRepository.Create(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var updatePayload struct {
		Name     *string `json:"name"`
		Email    *string `json:"email"`
		Password *string `json:"password"`
	}

	if err := c.BodyParser(&updatePayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var userUpdate = models.User{
		Name:     *updatePayload.Name,
		Email:    *updatePayload.Email,
		Password: *updatePayload.Password,
	}

	// validate user data
	if err := uh.validateUser(userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := uh.UserRepository.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if updatePayload.Name != nil {
		user.Name = *updatePayload.Name
	}

	if updatePayload.Email != nil {
		user.Email = *updatePayload.Email
	}

	if updatePayload.Password != nil && *updatePayload.Password != "" {
		hashedPassword, err := utils.HashPassword(*updatePayload.Password + user.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		user.Password = hashedPassword
	}

	if err := uh.UserRepository.Update(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Invalidate the cache for this user
	cacheKey := "user:" + strconv.FormatUint(id, 10)
	if err := redis.RedisClient.Del(redis.Ctx, cacheKey).Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to invalidate cache"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (uh *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := uh.UserRepository.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Invalidate the cache for this user
	cacheKey := "user:" + strconv.FormatUint(id, 10)
	if err := redis.RedisClient.Del(redis.Ctx, cacheKey).Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to invalidate cache"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("Delete user with ID=%v success", id)})
}

func (uh *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	if err := uh.UserRepository.FindAll(&users); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (uh *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	cacheKey := "user:" + strconv.FormatUint(id, 10)
	cachedUser, err := redis.Get(cacheKey)
	if err == nil && cachedUser != "" {
		var user models.User
		if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
			return c.JSON(user)
		}
	}

	user, err := uh.UserRepository.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (uh *UserHandler) UserRegister(c *fiber.Ctx) error {
	var registerPayload models.UserRegister
	if err := c.BodyParser(&registerPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var user = models.User{
		Name:     registerPayload.Name,
		Email:    registerPayload.Email,
		Password: registerPayload.Password,
		RoleID:   2,
	}

	// validate user data
	if err := uh.validateUser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// email check
	existingUser, err := uh.UserRepository.FindByEmail(user.Email)
	if err == nil && existingUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email already registered"})
	}

	// Hash the user's password
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Printf("user.Password: %v\n", user.Password)
	user.Password = hashed
	fmt.Printf("user.Password: %v\n", user.Password)

	if err := uh.UserRepository.Create(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (uh *UserHandler) UserLogin(c *fiber.Ctx) error {
	var loginPayload models.UserLogin
	if err := c.BodyParser(&loginPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var userLogin = models.User{
		Email:    loginPayload.Email,
		Password: loginPayload.Password,
	}

	// validate user data
	if err := uh.validateUser(userLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if loginPayload.Email == "" || loginPayload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Login data is empty"})
	}

	existingUser, err := uh.UserRepository.FindByEmail(loginPayload.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Check the provided password against the stored hashed password
	if !utils.CheckPassword(existingUser.Password, loginPayload.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// user credentials have been verified
	// generate jwt token
	token, err := middleware.CreateToken(&existingUser.ID, &existingUser.Email, &existingUser.RoleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    *token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "login successful"})
}

func (uh *UserHandler) UserLogout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
	})
	return c.SendString("logout successful")
}

func (uh *UserHandler) validateUser(user models.User) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return errors.New("invalid email format")
	}

	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(user.Password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(user.Password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(user.Password)

	if !hasUppercase || !hasLowercase || !hasNumber {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, and one number")
	}

	return nil
}
