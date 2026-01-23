# 🍱 Kasir API

A simple and lightweight Point of Sales (POS) API Service built with Go. This project is part of the "Bootcamp Golang CodeWithUmam".

## 🚀 Features

- **Product Management**: Full CRUD operations for products (Indomie Godog, Aqua, etc.).
- **Category Management**: Full CRUD operations for product categories.
- **In-Memory Storage**: Uses Go slices for fast, temporary data storage.
- **Swagger Documentation**: Interactive API documentation using Swagger UI.
- **Health Check**: Endpoint to verify system status.
- **Home Dashboard**: Simple HTML landing page with available endpoints.

## 🛠️ Tech Stack

- **Language**: Go (Golang) v1.25+
- **Router**: Standard library `net/http`
- **Documentation**: [Swaggo](https://github.com/swaggo/swag)
- **Hot Reload**: [Air](https://github.com/air-verse/air)

## 📋 Prerequisites

Ensure you have the following installed:
- [Go](https://go.dev/dl/)
- [Air](https://github.com/air-verse/air) (Optional, for hot reloading)

## ⚙️ Installation

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd kasir-api
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Generate Swagger Docs (Optional)**
   If you modify the API annotations, regenerate the docs:
   ```bash
   swag init
   ```

## ▶️ Running the Application

### Standard Run
```bash
go run main.go
```

### With Hot Reload (Recommended for Dev)
```bash
air
```

The server will start on `http://localhost:8080`.

## 📖 API Documentation

Once the server is running, you can access the interactive API documentation at:

👉 **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

## 🔗 Endpoints

### 📦 Products
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/api/produk` | Get all products |
| `GET` | `/api/produk/{id}` | Get product by ID |
| `POST` | `/api/produk` | Create a new product |
| `PUT` | `/api/produk/{id}` | Update a product |
| `DELETE` | `/api/produk/{id}` | Delete a product |

### 🏷️ Categories
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/api/categories` | Get all categories |
| `GET` | `/api/categories/{id}` | Get category by ID |
| `POST` | `/api/categories` | Create a new category |
| `PUT` | `/api/categories/{id}` | Update a category |
| `DELETE` | `/api/categories/{id}` | Delete a category |

### ⚙️ System
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/` | Home / Dashboard |
| `GET` | `/health` | Health Check |

## 📝 License
This project is licensed under the Apache 2.0 License.
