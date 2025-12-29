# Product API Service

A RESTful API for product CRUD operations built with Go and Gin framework.

## Features

- ✅ Complete CRUD operations for products
- ✅ In-memory database with sample data
- ✅ Input validation
- ✅ Category-based filtering
- ✅ RESTful API design

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/api/v1/products` | List all products |
| GET | `/api/v1/products?category=Electronics` | Filter products by category |
| GET | `/api/v1/products/:id` | Get product by ID |
| POST | `/api/v1/products` | Create new product |
| PUT | `/api/v1/products/:id` | Update product |
| DELETE | `/api/v1/products/:id` | Delete product |
| GET | `/api/v1/products/category/:category` | Get products by category |

## Quick Start

```bash
# Install dependencies
go mod download

# Run the server
go run main.go
```

Server will start on `http://localhost:8080`

## License

MIT
