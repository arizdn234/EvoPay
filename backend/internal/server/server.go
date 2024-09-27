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
	logHandler := handlers.NewTransactionLogHandler(logRepo)

	// Transaction repository with Redis client
	transactionRepo := repositories.NewTransactionRepository(db, logRepo, redis.RedisClient)
	transactionHandler := handlers.NewTransactionHandler(transactionRepo)

	// Balance repository with Redis caching
	balanceRepo := repositories.NewBalanceRepository(db, redis.RedisClient)
	balanceHandler := handlers.NewBalanceHandler(balanceRepo)

	// Deposit repository and handler
	depositRepo := repositories.NewDepositRepository(db, balanceRepo, logRepo, redis.RedisClient)
	depositHandler := handlers.NewDepositHandler(depositRepo)

	// Withdrawal repository and handler
	withdrawalRepo := repositories.NewWithdrawalRepository(db, balanceRepo, logRepo, redis.RedisClient)
	withdrawalHandler := handlers.NewWithdrawalHandler(withdrawalRepo)

	// <<<<<<<<<<====[ Define routes (API docs) ]====>>>>>>>>>>>
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

	// =============== Deposit routes ================
	depositRoute := app.Group("/deposits")
	depositRoute.Use(middleware.RequireAuth())

	// Deposit routes
	depositRoute.Post("/", depositHandler.CreateDeposit)                   // Create a deposit
	depositRoute.Get("/:id", depositHandler.GetDepositByID)                // Get deposit by ID
	depositRoute.Get("/users/:userID", depositHandler.GetDepositsByUserID) // Get all deposits by user ID

	// =============== Withdrawal routes ================
	withdrawalRoute := app.Group("/withdrawals")
	withdrawalRoute.Use(middleware.RequireAuth())

	// Withdrawal routes
	withdrawalRoute.Post("/", withdrawalHandler.CreateWithdrawal)                   // Create a withdrawal
	withdrawalRoute.Get("/:id", withdrawalHandler.GetWithdrawalByID)                // Get withdrawal by ID
	withdrawalRoute.Get("/users/:userID", withdrawalHandler.GetWithdrawalsByUserID) // Get all withdrawals by user ID

	// =============== Transaction Log routes ================
	logRoute := app.Group("/transaction-logs")
	logRoute.Use(middleware.RequireAuth())

	// Transaction Log routes
	logRoute.Post("/", logHandler.CreateTransactionLog)                                       // Create a new transaction log
	logRoute.Get("/transaction/:transactionID", logHandler.GetTransactionLogsByTransactionID) // Get logs by transaction ID
	logRoute.Get("/user/:userID", logHandler.GetTransactionLogsByUserID)                      // Get logs by user ID

	// Start the server on the specified port
	app.Listen(":" + port)

	return app
}
