package usecase

import (
	"database/sql"
	"myapp/core/orders/model"
)

type OrderUsecase interface {
	CreateOrder(order *model.Order) error
	GetOrder(id int) (*model.Order, error)
}

type orderUsecase struct {
	db *sql.DB
}

func NewOrderUsecase(db *sql.DB) OrderUsecase {
	return &orderUsecase{db: db}
}

func (u *orderUsecase) CreateOrder(order *model.Order) error {
	_, err := u.db.Exec("INSERT INTO orders (user_id, quantity) VALUES (?, ?)", order.UserID, order.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (u *orderUsecase) GetOrder(id int) (*model.Order, error) {
	order := &model.Order{}
	err := u.db.QueryRow("SELECT id, user_id, quantity FROM orders WHERE id = ?", id).Scan(&order.ID, &order.UserID, &order.Quantity)
	if err != nil {
		return nil, err
	}
	return order, nil
}
