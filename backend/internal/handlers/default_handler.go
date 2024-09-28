package handlers

import "github.com/gofiber/fiber/v2"

func Welcome(c *fiber.Ctx) error {
	// General service information
	info := map[string]interface{}{
		"service": "EvoPay",
		"version": "1.0",
	}

	// User routes
	userRoutes := []map[string]interface{}{
		{
			"method":      "GET",
			"path":        "/",
			"description": "Welcome & backend routes documentation",
		},
		{
			"method":      "POST",
			"path":        "/users/login",
			"description": "User login",
		},
		{
			"method":      "POST",
			"path":        "/users/register",
			"description": "Register new user",
		},
		{
			"method":      "GET",
			"path":        "/users/logout",
			"description": "User logout",
		},
		{
			"method":      "POST",
			"path":        "/users/send-verification-email",
			"description": "Send verification email",
		},
		{
			"method":      "GET",
			"path":        "/users/verify-email",
			"description": "Verify user email",
		},
		{
			"method":      "POST",
			"path":        "/users/send-reset-password-email",
			"description": "Send password reset email",
		},
		{
			"method":      "POST",
			"path":        "/users/reset-password",
			"description": "Reset user password",
		},
		{
			"method":      "GET",
			"path":        "/users",
			"description": "Get all users",
			"note":        "Requires authentication",
		},
		{
			"method":      "POST",
			"path":        "/users",
			"description": "Create a new user",
			"note":        "Requires authentication",
		},
		{
			"method":      "GET",
			"path":        "/users/{id}",
			"description": "Get user by ID",
			"note":        "Requires authentication",
		},
		{
			"method":      "PUT",
			"path":        "/users/{id}",
			"description": "Update user by ID",
			"note":        "Requires authentication",
		},
		{
			"method":      "DELETE",
			"path":        "/users/{id}",
			"description": "Delete user by ID",
			"note":        "Requires authentication",
		},
		{
			"method":      "POST",
			"path":        "/transactions",
			"description": "Create a new transaction",
			"note":        "Requires authentication",
		},
		{
			"method":      "GET",
			"path":        "/transactions",
			"description": "Get all transactions",
			"note":        "Requires authentication",
		},
		{
			"method":      "GET",
			"path":        "/transactions/{id}",
			"description": "Get a transaction by ID",
			"note":        "Requires authentication",
		},
		{
			"method":      "PUT",
			"path":        "/transactions/{id}",
			"description": "Update a transaction by ID",
			"note":        "Requires authentication",
		},
		{
			"method":      "POST",
			"path":        "/balances",
			"description": "Create a new balance",
			"note":        "Requires authentication",
		},
		{
			"method":      "GET",
			"path":        "/balances/{userID}",
			"description": "Get balance by user ID",
			"note":        "Requires authentication",
		},
		{
			"method":      "PUT",
			"path":        "/balances/{userID}",
			"description": "Update a balance",
			"note":        "Requires authentication",
		},
		{
			"method":      "POST",
			"path":        "/balances/{userID}/add",
			"description": "Add to a user's balance",
			"note":        "Requires authentication",
		},
		{
			"method":      "POST",
			"path":        "/balances/{userID}/subtract",
			"description": "Subtract from a user's balance",
			"note":        "Requires authentication",
		},
		{
			"method":      "POST",
			"path":        "/deposits",
			"description": "Create a new deposit",
			"note":        "Requires authentication",
		},
		{
			"method":      "GET",
			"path":        "/deposits/{id}",
			"description": "Get deposit by ID",
			"note":        "Requires authentication",
		},
		{
			"method":      "GET",
			"path":        "/deposits/users/{userID}",
			"description": "Get all deposits by user ID",
			"note":        "Requires authentication",
		},
		{
			"method":      "POST",
			"path":        "/withdrawals",
			"description": "Create a new withdrawal",
			"note":        "Requires authentication",
		},
		{
			"method":      "GET",
			"path":        "/withdrawals/{id}",
			"description": "Get withdrawal by ID",
			"note":        "Requires authentication",
		},
		{
			"method":      "GET",
			"path":        "/withdrawals/users/{userID}",
			"description": "Get all withdrawals by user ID",
			"note":        "Requires authentication",
		},
		{
			"method":      "POST",
			"path":        "/transaction-logs",
			"description": "Create a new transaction log",
			"note":        "Requires authentication",
		},
		{
			"method":      "GET",
			"path":        "/transaction-logs/transaction/{transactionID}",
			"description": "Get logs by transaction ID",
			"note":        "Requires authentication",
		},
		{
			"method":      "GET",
			"path":        "/transaction-logs/user/{userID}",
			"description": "Get logs by user ID",
			"note":        "Requires authentication",
		},
	}

	// Admin routes
	adminRoutes := []map[string]interface{}{
		{
			"method":      "GET",
			"path":        "/admin/users",
			"description": "Admin can get all users",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "GET",
			"path":        "/admin/users/{id}",
			"description": "Admin can get user by ID",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "PUT",
			"path":        "/admin/users/{id}",
			"description": "Admin can update user",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "DELETE",
			"path":        "/admin/users/{id}",
			"description": "Admin can delete user",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "POST",
			"path":        "/admin/transactions",
			"description": "Admin can create a transaction",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "GET",
			"path":        "/admin/transactions",
			"description": "Admin can get all transactions",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "GET",
			"path":        "/admin/transactions/{id}",
			"description": "Admin can get transaction by ID",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "PUT",
			"path":        "/admin/transactions/{id}",
			"description": "Admin can update transaction",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "POST",
			"path":        "/admin/balances",
			"description": "Admin can create a balance",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "GET",
			"path":        "/admin/balances/{userID}",
			"description": "Admin can get balance by user ID",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "PUT",
			"path":        "/admin/balances/{userID}",
			"description": "Admin can update a balance",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "POST",
			"path":        "/admin/deposits",
			"description": "Admin can create a deposit",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "GET",
			"path":        "/admin/deposits/{id}",
			"description": "Admin can get deposit by ID",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "GET",
			"path":        "/admin/deposits/users/{userID}",
			"description": "Admin can get all deposits by user ID",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "POST",
			"path":        "/admin/withdrawals",
			"description": "Admin can create a withdrawal",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "GET",
			"path":        "/admin/withdrawals/{id}",
			"description": "Admin can get withdrawal by ID",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "GET",
			"path":        "/admin/withdrawals/users/{userID}",
			"description": "Admin can get all withdrawals by user ID",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "POST",
			"path":        "/admin/transaction-logs",
			"description": "Admin can create a new transaction log",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "GET",
			"path":        "/admin/transaction-logs/transaction/{transactionID}",
			"description": "Admin can get logs by transaction ID",
			"note":        "Requires admin authentication",
		},
		{
			"method":      "GET",
			"path":        "/admin/transaction-logs/user/{userID}",
			"description": "Admin can get logs by user ID",
			"note":        "Requires admin authentication",
		},
	}

	// Combine the routes into the info map
	info["user_routes"] = userRoutes
	info["admin_routes"] = adminRoutes
	info["note"] = "Routes marked with 'note' require valid user or admin credentials."

	return c.JSON(info)
}
