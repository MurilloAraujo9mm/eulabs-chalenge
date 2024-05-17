package repository

import (
    "database/sql"
    "eulabsmyapp/core/products/domain/model"
    "eulabsmyapp/core/products/domain/repository"
)

type productRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) repository.ProductRepository {
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

func (r *productRepository) Update(product *model.Product) error {
    _, err := r.db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, product.ID)
    return err
}

func (r *productRepository) Delete(id string) error {
    _, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
    return err
}
