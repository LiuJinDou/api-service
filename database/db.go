package database

import (
	"api-service/models"
	"sync"
	"time"
)

// In-memory database for demo purposes
type DB struct {
	products map[int]*models.Product
	nextID   int
	mu       sync.RWMutex
}

var instance *DB
var once sync.Once

func GetDB() *DB {
	once.Do(func() {
		instance = &DB{
			products: make(map[int]*models.Product),
			nextID:   1,
		}
		// Seed with sample data
		instance.seedData()
	})
	return instance
}

func (db *DB) seedData() {
	sampleProducts := []models.Product{
		{
			Name:        "Laptop",
			Description: "High-performance laptop for developers",
			Price:       1299.99,
			Stock:       15,
			Category:    "Electronics",
		},
		{
			Name:        "Wireless Mouse",
			Description: "Ergonomic wireless mouse",
			Price:       29.99,
			Stock:       100,
			Category:    "Accessories",
		},
		{
			Name:        "Mechanical Keyboard",
			Description: "RGB mechanical gaming keyboard",
			Price:       149.99,
			Stock:       50,
			Category:    "Accessories",
		},
	}

	for _, p := range sampleProducts {
		db.CreateProduct(&p)
	}
}

func (db *DB) GetAllProducts() []models.Product {
	db.mu.RLock()
	defer db.mu.RUnlock()

	products := make([]models.Product, 0, len(db.products))
	for _, p := range db.products {
		products = append(products, *p)
	}
	return products
}

func (db *DB) GetProduct(id int) (*models.Product, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	product, exists := db.products[id]
	if !exists {
		return nil, false
	}
	return product, true
}

func (db *DB) CreateProduct(product *models.Product) *models.Product {
	db.mu.Lock()
	defer db.mu.Unlock()

	product.ID = db.nextID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	db.products[db.nextID] = product
	db.nextID++

	return product
}

func (db *DB) UpdateProduct(id int, updates *models.UpdateProductRequest) (*models.Product, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	product, exists := db.products[id]
	if !exists {
		return nil, false
	}

	if updates.Name != "" {
		product.Name = updates.Name
	}
	if updates.Description != "" {
		product.Description = updates.Description
	}
	if updates.Price > 0 {
		product.Price = updates.Price
	}
	if updates.Stock >= 0 {
		product.Stock = updates.Stock
	}
	if updates.Category != "" {
		product.Category = updates.Category
	}

	product.UpdatedAt = time.Now()
	return product, true
}

func (db *DB) DeleteProduct(id int) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exists := db.products[id]
	if !exists {
		return false
	}

	delete(db.products, id)
	return true
}

func (db *DB) SearchProducts(category string) []models.Product {
	db.mu.RLock()
	defer db.mu.RUnlock()

	products := make([]models.Product, 0)
	for _, p := range db.products {
		if category == "" || p.Category == category {
			products = append(products, *p)
		}
	}
	return products
}
