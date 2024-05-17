package repository

import (
    "database/sql"
    "myapp/core/products/model"
)

type ProductRepository interface {
    Create(product *model.Product) error
    GetByID(id string) (*model.Product, error)
    Update(id string, product *model.Product) error
    Delete(id string) error
}

type productRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
    return &productRepository{db: db}
}

func (r *productRepository) Create(product *model.Product) error {
    _, err := r.db.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", product.ID, product.Name, product.Price)
    return err
}

func (r *productRepository) GetByID(id string) (*model.Product, error) {
    product := &model.Product{}
    err := r.db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price)
    return product, err
}

func (r *productRepository) Update(id string, product *model.Product) error {
    _, err := r.db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, id)
    return err
}

func (r *productRepository) Delete(id string) error {
    _, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
    return err
}
