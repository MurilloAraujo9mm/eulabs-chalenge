package repository

import (
    "database/sql"
    "myapp/core/orders/domain/model"
    "myapp/core/orders/domain/repository"
)

type orderRepository struct {
    db *sql.DB
}

func NewOrderRepository(db *sql.DB) repository.OrderRepository {
    return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *model.Order) error {
    _, err := r.db.Exec("INSERT INTO orders (id, user_id, product_id, quantity, total, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)", order.ID, order.UserID, order.ProductID, order.Quantity, order.Total, order.CreatedAt, order.UpdatedAt)
    return err
}

func (r *orderRepository) GetByID(id string) (*model.Order, error) {
    order := &model.Order{}
    err := r.db.QueryRow("SELECT id, user_id, product_id, quantity, total, created_at, updated_at FROM orders WHERE id = ?", id).Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Total, &order.CreatedAt, &order.UpdatedAt)
    return order, err
}

func (r *orderRepository) Update(order *model.Order) error {
    _, err := r.db.Exec("UPDATE orders SET user_id = ?, product_id = ?, quantity = ?, total = ?, updated_at = ? WHERE id = ?", order.UserID, order.ProductID, order.Quantity, order.Total, order.UpdatedAt, order.ID)
    return err
}

func (r *orderRepository) Delete(id string) error {
    _, err := r.db.Exec("DELETE FROM orders WHERE id = ?", id)
    return err
}
