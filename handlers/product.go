package handlers

import (
	"api-service/database"
	"api-service/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	db *database.DB
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		db: database.GetDB(),
	}
}

// ListProducts retrieves all products with optional category filter
// @Summary List all products
// @Description Get a paginated list of all products, optionally filtered by category
// @Tags Products
// @Accept json
// @Produce json
// @Param category query string false "Filter by category"
// @Success 200 {object} models.ProductListResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	category := c.Query("category")

	var products []models.Product
	if category != "" {
		products = h.db.SearchProducts(category)
	} else {
		products = h.db.GetAllProducts()
	}

	response := models.ProductListResponse{
		Total:    len(products),
		Page:     1,
		PageSize: len(products),
		Products: products,
	}

	c.JSON(http.StatusOK, response)
}

// GetProduct retrieves a single product by ID
// @Summary Get product by ID
// @Description Get detailed information about a specific product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string "Invalid product ID"
// @Failure 404 {object} map[string]string "Product not found"
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, exists := h.db.GetProduct(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product
// @Summary Create a new product
// @Description Add a new product to the catalog
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.CreateProductRequest true "Product details"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string
// @Router /api/v1/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
	}

	createdProduct := h.db.CreateProduct(product)
	c.JSON(http.StatusCreated, createdProduct)
}

// UpdateProduct updates an existing product
// @Summary Update a product
// @Description Update product details by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.UpdateProductRequest true "Updated product details"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 404 {object} map[string]string "Product not found"
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var req models.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, exists := h.db.UpdateProduct(id, &req)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct deletes a product
// @Summary Delete a product
// @Description Remove a product from the catalog
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string "Product deleted successfully"
// @Failure 400 {object} map[string]string "Invalid product ID"
// @Failure 404 {object} map[string]string "Product not found"
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	success := h.db.DeleteProduct(id)
	if !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// GetProductsByCategory filters products by category
// @Summary Get products by category
// @Description Retrieve all products in a specific category
// @Tags Products
// @Accept json
// @Produce json
// @Param category path string true "Product category"
// @Success 200 {object} map[string]interface{} "Category products list"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/products/category/{category} [get]
func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")
	products := h.db.SearchProducts(category)

	c.JSON(http.StatusOK, gin.H{
		"category": category,
		"total":    len(products),
		"products": products,
	})
}
