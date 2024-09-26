package server

import (
	"github.com/arizdn234/EvoPay/internal/handlers"
	"github.com/arizdn234/EvoPay/internal/middleware"
	"github.com/arizdn234/EvoPay/internal/redis"
	"github.com/arizdn234/EvoPay/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RunServer(app *fiber.App, db *gorm.DB, port string) *fiber.App {
	// Initialize repositories and handlers
	// User repository
	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	// Transaction log repository
	logRepo := repositories.NewTransactionLogRepository(db, redis.RedisClient)

	// Transaction repository with redis client
	transactionRepo := repositories.NewTransactionRepository(db, logRepo, redis.RedisClient)
	transactionHandler := handlers.NewTransactionHandler(transactionRepo)

	// Balance repository with Redis caching
	balanceRepo := repositories.NewBalanceRepository(db, redis.RedisClient)
	balanceHandler := handlers.NewBalanceHandler(balanceRepo)

	// Define routes (API docs)
	app.Get("/", userHandler.Welcome)

	// =============== User routes ================
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

	// =============== Transaction routes ================
	transactionRoute := app.Group("/transactions")
	transactionRoute.Use(middleware.RequireAuth())
	transactionRoute.Post("/", transactionHandler.CreateTransaction)
	transactionRoute.Get("/", transactionHandler.GetAllTransactions)
	transactionRoute.Get("/:id", transactionHandler.GetTransactionByID)
	transactionRoute.Put("/:id", transactionHandler.UpdateTransaction)

	// =============== Balance routes ================
	balanceRoute := app.Group("/balances")
	balanceRoute.Use(middleware.RequireAuth())

	// Balance routes
	balanceRoute.Post("/", balanceHandler.CreateBalance)                       // Create a balance
	balanceRoute.Get("/:userID", balanceHandler.GetBalanceByUserID)            // Get balance by user ID
	balanceRoute.Put("/:userID", balanceHandler.UpdateBalance)                 // Update a balance
	balanceRoute.Post("/:userID/add", balanceHandler.AddToBalance)             // Add to a balance
	balanceRoute.Post("/:userID/subtract", balanceHandler.SubtractFromBalance) // Subtract from a balance

	// Start the server on the specified port
	app.Listen(":" + port)

	return app
}
