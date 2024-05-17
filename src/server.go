package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var db *sql.DB

func initDB() {
	var err error
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
}

func main() {
	initDB()
	defer db.Close()

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "API is running now"})
	})

	e.POST("/products", createProduct)
	e.GET("/products/:id", getProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}

func createProduct(c echo.Context) error {
	product := &Product{}
	if err := c.Bind(product); err != nil {
		return err
	}

	result, err := db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = int(id)

	return c.JSON(http.StatusCreated, product)
}

func getProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product := &Product{}
	err := db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
		}
		return err
	}

	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product := &Product{ID: id}
	if err := c.Bind(product); err != nil {
		return err
	}

	_, err := db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
