package createProduct

import (
    "myapp/core/products/domain/model"
    "myapp/core/products/domain/repository"
    "myapp/go_modules/products/dto"
    "github.com/google/uuid"
)

type CreateProductUsecase interface {
    Execute(input dto.CreateProductDTO) (*model.Product, error)
}

type createProductUsecase struct {
    repo repository.ProductRepository
}

func NewCreateProductUsecase(repo repository.ProductRepository) CreateProductUsecase {
    return &createProductUsecase{repo: repo}
}

func (u *createProductUsecase) Execute(input dto.CreateProductDTO) (*model.Product, error) {
    product := &model.Product{
        ID:    uuid.New().String(),
        Name:  input.Name,
        Price: input.Price,
    }
    err := u.repo.Create(product)
    if err != nil {
        return nil, err
    }
    return product, nil
}
