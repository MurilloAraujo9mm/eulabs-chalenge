package getProduct

import (
    "myapp/core/products/model"
    "myapp/core/products/repository"
)

type GetProductUsecase interface {
    Execute(id string) (*model.Product, error)
}

type getProductUsecase struct {
    repo repository.ProductRepository
}

func NewGetProductUsecase(repo repository.ProductRepository) GetProductUsecase {
    return &getProductUsecase{repo: repo}
}

func (u *getProductUsecase) Execute(id string) (*model.Product, error) {
    return u.repo.GetByID(id)
}
