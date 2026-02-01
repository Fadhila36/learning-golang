package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"kasir-api/database"
	_ "kasir-api/docs"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"

	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

const homeHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kasir API Service</title>
    <style>
        :root {
            --primary: #2563eb;
            --secondary: #475569;
            --bg: #f8fafc;
            --card-bg: #ffffff;
            --text: #1e293b;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
            line-height: 1.6;
            color: var(--text);
            background: var(--bg);
            margin: 0;
            padding: 2rem;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
        }
        header {
            text-align: center;
            margin-bottom: 3rem;
        }
        h1 {
            color: var(--primary);
            font-size: 2.5rem;
            margin-bottom: 0.5rem;
        }
        .subtitle {
            color: var(--secondary);
            font-size: 1.1rem;
        }
        .grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
            gap: 1.5rem;
            margin-bottom: 2rem;
        }
        .card {
            background: var(--card-bg);
            padding: 1.5rem;
            border-radius: 8px;
            box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1);
            transition: transform 0.2s;
        }
        .card:hover {
            transform: translateY(-2px);
        }
        .card h2 {
            margin-top: 0;
            color: var(--primary);
            font-size: 1.25rem;
            display: flex;
            align-items: center;
            gap: 0.5rem;
        }
        .badge {
            background: #e0e7ff;
            color: var(--primary);
            padding: 0.25rem 0.75rem;
            border-radius: 9999px;
            font-size: 0.75rem;
            font-weight: 600;
        }
        code {
            background: #f1f5f9;
            padding: 0.2rem 0.4rem;
            border-radius: 4px;
            font-size: 0.875rem;
            color: #ec4899;
        }
        .endpoint {
            margin-bottom: 0.5rem;
            font-family: monospace;
        }
        a {
            color: var(--primary);
            text-decoration: none;
            font-weight: 500;
        }
        a:hover {
            text-decoration: underline;
        }
        .method {
            font-weight: bold;
            display: inline-block;
            width: 50px;
        }
        .get { color: #059669; }
        .post { color: #2563eb; }
        .put { color: #d97706; }
        .delete { color: #dc2626; }
        footer {
            text-align: center;
            margin-top: 3rem;
            color: var(--secondary);
            font-size: 0.875rem;
        }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1>🍱 Kasir API</h1>
            <div class="subtitle">Simple Point of Sales Service built with Go</div>
        </header>

        <div class="grid">
            <div class="card">
                <h2>Category Service <span class="badge">CRUD</span></h2>
                <p>Manage product categories.</p>
                <div class="endpoint"><span class="method get">GET</span> <a href="/api/categories">/api/categories</a></div>
                <div class="endpoint"><span class="method post">POST</span> /api/categories</div>
                <div class="endpoint"><span class="method put">PUT</span> /api/categories/:id</div>
                <div class="endpoint"><span class="method delete">DEL</span> /api/categories/:id</div>
            </div>

            <div class="card">
                <h2>Product Service <span class="badge">CRUD</span></h2>
                <p>Manage inventory items.</p>
                <div class="endpoint"><span class="method get">GET</span> <a href="/api/produk">/api/produk</a></div>
                <div class="endpoint"><span class="method post">POST</span> /api/produk</div>
                <div class="endpoint"><span class="method put">PUT</span> /api/produk/:id</div>
                <div class="endpoint"><span class="method delete">DEL</span> /api/produk/:id</div>
            </div>
        </div>

        <div class="card">
            <h2>📦 System Info</h2>
            <p><strong>Status:</strong> <span style="color: #059669; font-weight: bold;">Operational</span></p>
            <p><strong>Version:</strong> v1.0.0</p>
            <p><strong>Documentation:</strong> <a href="/swagger/index.html">/swagger/index.html</a></p>
            <p><strong>Health Check:</strong> <a href="/health">/health</a></p>
        </div>

        <footer>
            <p>Bootcamp Golang CodeWithUmam &copy; 2026</p>
        </footer>
    </div>
</body>
</html>`

// @title Kasir API
// @version 1.0
// @description Simple Point of Sales API (Bootcamp Golang)
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host kasir-app.fadhilaabiyyu.my.id
// @schemes https http
// @BasePath /api
func main() {
	// Load config dengan Viper
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Default port jika tidak di-set
	if config.Port == "" {
		config.Port = "8080"
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Dependency Injection
	// Product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Setup routes - Products
	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	// Setup routes - Categories
	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Api Running",
		})
	})

	// Home Page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, homeHTML)
	})

	// Swagger UI
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Start server
	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}
