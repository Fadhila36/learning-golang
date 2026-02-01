# 🍱 Kasir API

A simple and lightweight Point of Sales (POS) API Service built with Go. This project is part of the "Bootcamp Golang CodeWithUmam".

## 🌐 Live Demo

� **[https://kasir-app.fadhilaabiyyu.my.id](https://kasir-app.fadhilaabiyyu.my.id)**

## �🚀 Features

- **Product Management**: Full CRUD operations for products with category support
- **Category Management**: Full CRUD operations for product categories
- **PostgreSQL Database**: Persistent storage using Supabase
- **Layered Architecture**: Clean separation (Handler → Service → Repository → Model)
- **Swagger Documentation**: Interactive API documentation
- **Health Check**: Endpoint to verify system status

## 🏗️ Architecture

```
kasir-api/
├── database/          # Database connection
├── models/            # Data structures
├── repositories/      # Database operations
├── services/          # Business logic
├── handlers/          # HTTP request handlers
├── docs/              # Swagger documentation
├── .env               # Environment config
└── main.go            # Entry point + DI
```

## 🛠️ Tech Stack

- **Language**: Go (Golang) v1.25+
- **Database**: PostgreSQL (Supabase)
- **Driver**: pgx v5
- **Config**: Viper
- **Documentation**: Swaggo
- **Deployment**: Railway

## 📋 Prerequisites

- [Go](https://go.dev/dl/) v1.25+
- [Supabase](https://supabase.com/) account (free tier works)

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

3. **Setup environment**
   ```bash
   cp .env.example .env
   # Edit .env with your Supabase credentials
   ```

4. **Create database tables** (run in Supabase SQL Editor)
   ```sql
   CREATE TABLE categories (
     id SERIAL PRIMARY KEY,
     name VARCHAR NOT NULL,
     description VARCHAR,
     created_at TIMESTAMP DEFAULT NOW(),
     updated_at TIMESTAMP
   );

   CREATE TABLE products (
     id SERIAL PRIMARY KEY,
     name VARCHAR NOT NULL,
     price INT NOT NULL,
     stock INT NOT NULL,
     category_id INT REFERENCES categories(id),
     created_at TIMESTAMP DEFAULT NOW(),
     updated_at TIMESTAMP
   );
   ```

## ▶️ Running the Application

```bash
go run main.go
```

Server will start on `http://localhost:8080`

## 📖 API Documentation

- **Local**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- **Production**: [https://kasir-app.fadhilaabiyyu.my.id/swagger/index.html](https://kasir-app.fadhilaabiyyu.my.id/swagger/index.html)

## 🔗 Endpoints

### 📦 Products
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/api/produk` | Get all products (with category_name) |
| `GET` | `/api/produk/{id}` | Get product by ID (with category_name) |
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
| `GET` | `/` | Home Dashboard |
| `GET` | `/health` | Health Check |

## 🚀 Deployment

This app is deployed on Railway with auto-deploy from GitHub.

**Environment Variables** (set in Railway):
```
PORT=8080
DB_CONN=postgresql://postgres.[PROJECT_ID]:[PASSWORD]@aws-1-ap-southeast-2.pooler.supabase.com:6543/postgres?sslmode=require
```

## 📝 License

This project is licensed under the Apache 2.0 License.
