package createProduct

import (
    "myapp/core/products/model"
    "myapp/core/products/repository"
    "myapp/core/products/dto"

    "github.com/google/uuid"
)

type CreateProductUsecase interface {
    Execute(dto dto.CreateProductDTO) (*model.Product, error)
}

type createProductUsecase struct {
    repo repository.ProductRepository
}

func NewCreateProductUsecase(repo repository.ProductRepository) CreateProductUsecase {
    return &createProductUsecase{repo: repo}
}

func (u *createProductUsecase) Execute(dto dto.CreateProductDTO) (*model.Product, error) {
    product := &model.Product{
        ID:    uuid.New().String(), 
        Name:  dto.Name,
        Price: dto.Price,
    }

    err := u.repo.Create(product)
    if err != nil {
        return nil, err
    }
    return product, nil
}
