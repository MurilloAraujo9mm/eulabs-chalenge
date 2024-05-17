package updateOrder

import (
    "myapp/core/orders/domain/model"
    "myapp/core/orders/domain/repository"
    "myapp/go_modules/orders/dto"
    "time"
)

type UpdateOrderUsecase interface {
    Execute(id string, input dto.UpdateOrderDTO) (*model.Order, error)
}

type updateOrderUsecase struct {
    repo repository.OrderRepository
}

func NewUpdateOrderUsecase(repo repository.OrderRepository) UpdateOrderUsecase {
    return &updateOrderUsecase{repo: repo}
}

func (u *updateOrderUsecase) Execute(id string, input dto.UpdateOrderDTO) (*model.Order, error) {
    order, err := u.repo.GetByID(id)
    if err != nil {
        return nil, err
    }

    order.Quantity = input.Quantity
    order.Total = input.Total
    order.UpdatedAt = time.Now()

    err = u.repo.Update(order)
    if err != nil {
        return nil, err
    }

    return order, nil
}
