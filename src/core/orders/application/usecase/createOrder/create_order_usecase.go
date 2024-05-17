package createOrder

import (
    "eulabsmyapp/core/orders/domain/model"
    "eulabsmyapp/core/orders/domain/repository"
    "eulabsmyapp/go_modules/orders/dto"
    "github.com/google/uuid"
    "time"
)

type CreateOrderUsecase interface {
    Execute(input dto.CreateOrderDTO) (*model.Order, error)
}

type createOrderUsecase struct {
    repo repository.OrderRepository
}

func NewCreateOrderUsecase(repo repository.OrderRepository) CreateOrderUsecase {
    return &createOrderUsecase{repo: repo}
}

func (u *createOrderUsecase) Execute(input dto.CreateOrderDTO) (*model.Order, error) {
    order := &model.Order{
        ID:        uuid.New().String(),
        UserID:    input.UserID,
        ProductID: input.ProductID,
        Quantity:  input.Quantity,
        Total:     input.Total,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    err := u.repo.Create(order)
    if err != nil {
        return nil, err
    }
    return order, nil
}
