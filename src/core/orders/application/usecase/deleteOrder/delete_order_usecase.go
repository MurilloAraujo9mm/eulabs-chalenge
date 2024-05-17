package deleteOrder

import "myapp/core/orders/domain/repository"

type DeleteOrderUsecase interface {
    Execute(id string) error
}

type deleteOrderUsecase struct {
    repo repository.OrderRepository
}

func NewDeleteOrderUsecase(repo repository.OrderRepository) DeleteOrderUsecase {
    return &deleteOrderUsecase{repo: repo}
}

func (u *deleteOrderUsecase) Execute(id string) error {
    return u.repo.Delete(id)
}
