package getOrder

import (
    "eulabsmyapp/core/orders/domain/model"
    "eulabsmyapp/core/orders/domain/repository"
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
