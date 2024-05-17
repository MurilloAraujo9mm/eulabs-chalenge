package getOrder

import (
    "myapp/core/orders/domain/model"
    "myapp/core/orders/domain/repository"
)

type GetOrderUsecase interface {
    Execute(id string) (*model.Order, error)
}

type getOrderUsecase struct {
    repo repository.OrderRepository
}

func NewGetOrderUsecase(repo repository.OrderRepository) GetOrderUsecase {
    return &getOrderUsecase{repo: repo}
}

func (u *getOrderUsecase) Execute(id string) (*model.Order, error) {
    return u.repo.GetByID(id)
}
