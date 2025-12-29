package main

import (
	"api-service/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "Product API",
			"version": "1.0.0",
		})
	})

	// Initialize handlers
	productHandler := handlers.NewProductHandler()

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Product CRUD endpoints
		products := v1.Group("/products")
		{
			products.GET("", productHandler.ListProducts)                          // List all products (with optional ?category= filter)
			products.GET("/:id", productHandler.GetProduct)                        // Get single product
			products.POST("", productHandler.CreateProduct)                        // Create product
			products.PUT("/:id", productHandler.UpdateProduct)                     // Update product
			products.DELETE("/:id", productHandler.DeleteProduct)                  // Delete product
			products.GET("/category/:category", productHandler.GetProductsByCategory) // Get by category
		}
	}

	log.Println("🚀 Product API Server starting on :8080")
	log.Println("📝 Endpoints:")
	log.Println("   GET    /health")
	log.Println("   GET    /api/v1/products")
	log.Println("   GET    /api/v1/products/:id")
	log.Println("   POST   /api/v1/products")
	log.Println("   PUT    /api/v1/products/:id")
	log.Println("   DELETE /api/v1/products/:id")
	log.Println("   GET    /api/v1/products/category/:category")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
