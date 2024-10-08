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
	app.Get("/", handlers.Welcome)

	// <<=============== [[[[ User routes ]]]] ================>>
	userRoute := app.Group("/users")
	userRoute.Post("/login", userHandler.UserLogin)
	userRoute.Post("/register", userHandler.UserRegister)
	userRoute.Post("/logout", userHandler.UserLogout)

	// Email verification and password recovery routes
	userRoute.Post("/send-verification-email", userHandler.SendVerificationEmail)
	userRoute.Get("/verify-email", userHandler.VerifyEmail)
	userRoute.Post("/send-reset-password-email", userHandler.SendResetPasswordEmail)
	userRoute.Post("/reset-password", userHandler.ResetPassword)

	// Use middleware for authentication on user routes
	userRoute.Use(middleware.RequireAuth())

	// User routes that require authentication
	userRoute.Get("/me", userHandler.GetMe)
	userRoute.Get("", userHandler.GetAllUsers)       // Get all users
	userRoute.Get("/:id", userHandler.GetUserByID)   // Get user by ID
	userRoute.Post("", userHandler.CreateUser)       // Create a new user
	userRoute.Put("/:id", userHandler.UpdateUser)    // Update user by ID
	userRoute.Delete("/:id", userHandler.DeleteUser) // Delete user by ID

	// =============== Transaction routes ================
	transactionRoute := app.Group("/transactions")
	transactionRoute.Use(middleware.RequireAuth())
	transactionRoute.Post("/", transactionHandler.CreateTransaction)
	transactionRoute.Get("/", transactionHandler.GetAllTransactions)
	transactionRoute.Get("/user/:userID", transactionHandler.GetTransactionsByUserID)
	transactionRoute.Get("/:id", transactionHandler.GetTransactionByID)
	transactionRoute.Put("/:id", transactionHandler.UpdateTransaction)

	// =============== Balance routes ================
	balanceRoute := app.Group("/balances")
	balanceRoute.Use(middleware.RequireAuth())
	// balanceRoute.Post("/", balanceHandler.CreateBalance)                       // Create a balance
	balanceRoute.Get("/:userID", balanceHandler.GetBalanceByUserID)            // Get balance by user ID
	balanceRoute.Put("/:userID", balanceHandler.UpdateBalance)                 // Update a balance
	balanceRoute.Post("/:userID/add", balanceHandler.AddToBalance)             // Add to a balance
	balanceRoute.Post("/:userID/subtract", balanceHandler.SubtractFromBalance) // Subtract from a balance

	// =============== Deposit routes ================
	depositRoute := app.Group("/deposits")
	depositRoute.Use(middleware.RequireAuth())
	depositRoute.Post("/", depositHandler.CreateDeposit)                   // Create a deposit
	depositRoute.Get("/:id", depositHandler.GetDepositByID)                // Get deposit by ID
	depositRoute.Get("/users/:userID", depositHandler.GetDepositsByUserID) // Get all deposits by user ID

	// =============== Withdrawal routes ================
	withdrawalRoute := app.Group("/withdrawals")
	withdrawalRoute.Use(middleware.RequireAuth())
	withdrawalRoute.Post("/", withdrawalHandler.CreateWithdrawal)                   // Create a withdrawal
	withdrawalRoute.Get("/:id", withdrawalHandler.GetWithdrawalByID)                // Get withdrawal by ID
	withdrawalRoute.Get("/users/:userID", withdrawalHandler.GetWithdrawalsByUserID) // Get all withdrawals by user ID

	// =============== Transaction Log routes ================
	logRoute := app.Group("/transaction-logs")
	logRoute.Use(middleware.RequireAuth())
	logRoute.Post("/", logHandler.CreateTransactionLog)                                       // Create a new transaction log
	logRoute.Get("/transaction/:transactionID", logHandler.GetTransactionLogsByTransactionID) // Get logs by transaction ID
	logRoute.Get("/user/:userID", logHandler.GetTransactionLogsByUserID)                      // Get logs by user ID

	// <<=============== [[[[ Admin routes ]]]] ================>>
	adminRoute := app.Group("/admin")
	adminRoute.Use(middleware.RequireAuth())      // Require authentication for admin routes
	adminRoute.Use(middleware.RequireAdminRole()) // Require admin role for the following routes

	// =============== User Management ================
	adminRoute.Get("/users", userHandler.GetAllUsers)       // Admin can get all users
	adminRoute.Get("/users/:id", userHandler.GetUserByID)   // Admin can get user by ID
	adminRoute.Put("/users/:id", userHandler.UpdateUser)    // Admin can update user
	adminRoute.Delete("/users/:id", userHandler.DeleteUser) // Admin can delete user

	// =============== Transaction Management ================
	adminRoute.Post("/transactions", transactionHandler.CreateTransaction)     // Admin can create a transaction
	adminRoute.Get("/transactions", transactionHandler.GetAllTransactions)     // Admin can get all transactions
	adminRoute.Get("/transactions/:id", transactionHandler.GetTransactionByID) // Admin can get transaction by ID
	adminRoute.Put("/transactions/:id", transactionHandler.UpdateTransaction)  // Admin can update a transaction

	// =============== Balance Management ================
	adminRoute.Post("/balances", balanceHandler.CreateBalance)                        // Admin can create a balance
	adminRoute.Get("/balances/:userID", balanceHandler.GetBalanceByUserID)            // Admin can get balance by user ID
	adminRoute.Put("/balances/:userID", balanceHandler.UpdateBalance)                 // Admin can update a balance
	adminRoute.Post("/balances/:userID/add", balanceHandler.AddToBalance)             // Admin can add to a balance
	adminRoute.Post("/balances/:userID/subtract", balanceHandler.SubtractFromBalance) // Admin can subtract from a balance

	// =============== Deposit Management ================
	adminRoute.Post("/deposits", depositHandler.CreateDeposit)                    // Admin can create a deposit
	adminRoute.Get("/deposits/:id", depositHandler.GetDepositByID)                // Admin can get deposit by ID
	adminRoute.Get("/deposits/users/:userID", depositHandler.GetDepositsByUserID) // Admin can get all deposits by user ID

	// =============== Withdrawal Management ================
	adminRoute.Post("/withdrawals", withdrawalHandler.CreateWithdrawal)                    // Admin can create a withdrawal
	adminRoute.Get("/withdrawals/:id", withdrawalHandler.GetWithdrawalByID)                // Admin can get withdrawal by ID
	adminRoute.Get("/withdrawals/users/:userID", withdrawalHandler.GetWithdrawalsByUserID) // Admin can get all withdrawals by user ID

	// =============== Transaction Log Management ================
	adminRoute.Post("/transaction-logs", logHandler.CreateTransactionLog)                                        // Admin can create a new transaction log
	adminRoute.Get("/transaction-logs/transaction/:transactionID", logHandler.GetTransactionLogsByTransactionID) // Admin can get logs by transaction ID
	adminRoute.Get("/transaction-logs/user/:userID", logHandler.GetTransactionLogsByUserID)                      // Admin can get logs by user ID

	// Start the server on the specified port
	app.Listen(":" + port)

	return app
}
