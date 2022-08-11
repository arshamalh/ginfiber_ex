package main

func main() {
	// Create new instance of gin app
	gin := NewGinApp()
	// Add routes to gin app
	gin.RegisterRoutes()
	// Start gin server
	gin.Start()

	// Create new instance of fiber app
	fiber := NewFiberApp()
	// Add routes to fiber app
	fiber.RegisterRoutes()
	// Start fiber server
	fiber.Start()
}
