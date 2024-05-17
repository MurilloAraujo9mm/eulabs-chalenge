package repository

import "myapp/core/orders/model"

type OrderRepository interface {
	Create(order *model.Order) error
	FindByID(id int) (*model.Order, error)
}
