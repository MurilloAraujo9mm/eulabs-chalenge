package repository

import (
    "database/sql"
    "log"
    "time" 

    "eulabsmyapp/core/orders/domain/model"
    "eulabsmyapp/core/orders/domain/repository"
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
    log.Printf("Fetching order with ID: %s", id) 
    order := &model.Order{}
    var createdAt, updatedAt []uint8
    err := r.db.QueryRow("SELECT id, user_id, product_id, quantity, total, created_at, updated_at FROM orders WHERE id = ?", id).Scan(
        &order.ID,
        &order.UserID,
        &order.ProductID,
        &order.Quantity,
        &order.Total,
        &createdAt,
        &updatedAt,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("No order found with ID: %s", id)
            return nil, err
        }
        log.Printf("Error fetching order with ID: %s, error: %v", id, err)
        return nil, err
    }
    order.CreatedAt, err = parseTime(createdAt)
    if err != nil {
        return nil, err
    }
    order.UpdatedAt, err = parseTime(updatedAt)
    if err != nil {
        return nil, err
    }
    return order, nil
}

func (r *orderRepository) Update(order *model.Order) error {
    _, err := r.db.Exec("UPDATE orders SET user_id = ?, product_id = ?, quantity = ?, total = ?, updated_at = ? WHERE id = ?", order.UserID, order.ProductID, order.Quantity, order.Total, order.UpdatedAt, order.ID)
    return err
}

func (r *orderRepository) Delete(id string) error {
    _, err := r.db.Exec("DELETE FROM orders WHERE id = ?", id)
    return err
}

func parseTime(b []uint8) (time.Time, error) {
    str := string(b)
    return time.Parse("2006-01-02 15:04:05", str)
}
