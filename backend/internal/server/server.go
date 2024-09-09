package server

import (
	"github.com/arizdn234/EvoPay/internal/handlers"
	"github.com/arizdn234/EvoPay/internal/middleware"
	"github.com/arizdn234/EvoPay/internal/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RunServer(app *fiber.App, db *gorm.DB, port string) *fiber.App {
	// Initialize repositories and handlers
	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	// Define routes
	app.Get("/", userHandler.Welcome)

	// User routes
	userRoute := app.Group("/users")
	userRoute.Post("/login", userHandler.UserLogin)
	userRoute.Post("/register", userHandler.UserRegister)
	userRoute.Get("/logout", userHandler.UserLogout)

	// Use middleware for authentication
	userRoute.Use(middleware.RequireAuth())
	userRoute.Use(middleware.RequireAdminRole())

	// Routes that require authentication
	userRoute.Get("", userHandler.GetAllUsers)
	userRoute.Get("/:id", userHandler.GetUserByID)
	userRoute.Post("", userHandler.CreateUser)
	userRoute.Put("/:id", userHandler.UpdateUser)
	userRoute.Delete("/:id", userHandler.DeleteUser)

	// Start the server on the specified port
	app.Listen(":" + port)

	return app
}
