package server

import (
	"github.com/arizdn234/EvoPay/internal/handlers"
	"github.com/arizdn234/EvoPay/internal/middleware"
	"github.com/arizdn234/EvoPay/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RunServer(app *fiber.App, db *gorm.DB, port string) *fiber.App {
	// Initialize repositories and handlers
	// user===============
	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)
	// transaction===============
	TransactionRepo := repositories.NewTransactionRepository(db)
	TransactionHandler := handlers.NewTransactionHandler(TransactionRepo)

	// Define routes (api docs)
	app.Get("/", userHandler.Welcome)

	// ===============User routes===============
	userRoute := app.Group("/users")
	userRoute.Post("/login", userHandler.UserLogin)
	userRoute.Post("/register", userHandler.UserRegister)
	userRoute.Post("/logout", userHandler.UserLogout)

	// Email verification and password recovery routes
	userRoute.Post("/send-verification-email", userHandler.SendVerificationEmail)
	userRoute.Get("/verify-email", userHandler.VerifyEmail)
	userRoute.Post("/send-reset-password-email", userHandler.SendResetPasswordEmail)
	userRoute.Post("/reset-password", userHandler.ResetPassword)

	// Use middleware for authentication
	userRoute.Use(middleware.RequireAuth())
	userRoute.Use(middleware.RequireAdminRole())

	// User routes that require authentication
	userRoute.Get("", userHandler.GetAllUsers)
	userRoute.Get("/:id", userHandler.GetUserByID)
	userRoute.Post("", userHandler.CreateUser)
	userRoute.Put("/:id", userHandler.UpdateUser)
	userRoute.Delete("/:id", userHandler.DeleteUser)

	// ===============Transaction routes===============
	transactionRoute := app.Group("/transactions")
	transactionRoute.Use(middleware.RequireAuth())
	transactionRoute.Post("/", TransactionHandler.CreateTransaction)
	transactionRoute.Get("/", TransactionHandler.GetAllTransactions)
	transactionRoute.Get("/:id", TransactionHandler.GetTransactionByID)
	transactionRoute.Put("/:id", TransactionHandler.UpdateTransaction)
	transactionRoute.Delete("/:id", TransactionHandler.DeleteTransaction)

	// Start the server on the specified port
	app.Listen(":" + port)

	return app
}
