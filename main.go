package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "kasir-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{}

var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 15000, Stok: 10},
	{ID: 2, Nama: "Aqua 600ml", Harga: 3000, Stok: 20},
	{ID: 3, Nama: "Ayam Goreng", Harga: 25000, Stok: 30},
}

// @Summary Get product by ID
// @Description Get details of a single product
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} Produk
// @Failure 400,404 {string} string "Invalid ID or Product not found"
// @Router /produk/{id} [get]
func getProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "produk Belum Ada", http.StatusNotFound)
}

// @Summary Update product by ID
// @Description Update an existing product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body Produk true "Updated product data"
// @Success 200 {object} Produk
// @Failure 400,404 {string} string "Invalid request or Product not found"
// @Router /produk/{id} [put]
func updateProduk(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid produk ID", http.StatusBadRequest)
		return
	}
	// get data dari request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	// loop produk, cari id, ganti sesuai data dari request
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk[i])
			return
		}
	}
	http.Error(w, "produk Belum Ada", http.StatusNotFound)
}

// @Summary Delete product by ID
// @Description Delete a product
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 400,404 {string} string "Invalid ID or Product not found"
// @Router /produk/{id} [delete]
func deleteProduk(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid produk ID", http.StatusBadRequest)
		return
	}
	// loop produk, cari id, dapat index yang di hapus
	for i, p := range produk {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesuah index
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "produk berhasil dihapus",
			})
			return
		}
	}
	http.Error(w, "produk Belum Ada", http.StatusNotFound)

}

// @Summary Get all products
// @Description Get list of all products
// @Tags products
// @Produce json
// @Success 200 {array} Produk
// @Router /produk [get]
func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produk)
}

// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body Produk true "New product data"
// @Success 201 {object} Produk
// @Failure 400 {string} string "Invalid request body"
// @Router /produk [post]
func createProduct(w http.ResponseWriter, r *http.Request) {
	var produkBaru Produk
	err := json.NewDecoder(r.Body).Decode(&produkBaru)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	produkBaru.ID = len(produk) + 1
	produk = append(produk, produkBaru)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produkBaru)
}

// @Summary Get all categories
// @Description Get list of all categories
// @Tags categories
// @Produce json
// @Success 200 {array} Category
// @Router /categories [get]
func getAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// @Summary Get category by ID
// @Description Get details of a single category
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} Category
// @Failure 400,404 {string} string "Invalid ID or Category not found"
// @Router /categories/{id} [get]
func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	for _, c := range categories {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

// @Summary Create a new category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body Category true "New category data"
// @Success 201 {object} Category
// @Failure 400 {string} string "Invalid request body"
// @Router /categories [post]
func createCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}

// @Summary Update category by ID
// @Description Update an existing category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body Category true "Updated category data"
// @Success 200 {object} Category
// @Failure 400,404 {string} string "Invalid request or Category not found"
// @Router /categories/{id} [put]
func updateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var updatedRec Category
	err = json.NewDecoder(r.Body).Decode(&updatedRec)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, c := range categories {
		if c.ID == id {
			categories[i].Name = updatedRec.Name
			categories[i].Description = updatedRec.Description
			// ID cannot be changed, keeping existing ID
			updatedRec.ID = id

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories[i])
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

// @Summary Delete category by ID
// @Description Delete a category
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Failure 400,404 {string} string "Invalid ID or Category not found"
// @Router /categories/{id} [delete]
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	for i, c := range categories {
		if c.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted successfully"})
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
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

// @host localhost:8080
// @BasePath /api
func main() {

	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}
	})
	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getAllProducts(w, r)
		} else if r.Method == "POST" {
			createProduct(w, r)
		}
	})

	// Categories Endpoints
	// GET /categories
	// POST /categories
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getAllCategories(w, r)
		} else if r.Method == http.MethodPost {
			createCategory(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// GET /categories/{id}
	// PUT /categories/{id}
	// DELETE /categories/{id}
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getCategoryByID(w, r)
		} else if r.Method == http.MethodPut {
			updateCategory(w, r)
		} else if r.Method == http.MethodDelete {
			deleteCategory(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	// GET localhost:8080/health
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

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server")
	}
}
