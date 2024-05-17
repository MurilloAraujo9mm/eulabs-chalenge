package repository

import "myapp/core/orders/domain/model"

type OrderRepository interface {
    Create(order *model.Order) error
    GetByID(id string) (*model.Order, error)
    Update(order *model.Order) error
    Delete(id string) error
}
