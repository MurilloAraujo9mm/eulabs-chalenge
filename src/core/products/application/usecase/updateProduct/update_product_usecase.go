package updateProduct

import (
    "myapp/core/products/domain/model"
    "myapp/core/products/domain/repository"
    "myapp/go_modules/products/dto"
)

type UpdateProductUsecase interface {
    Execute(id string, input dto.UpdateProductDTO) (*model.Product, error)
}

type updateProductUsecase struct {
    repo repository.ProductRepository
}

func NewUpdateProductUsecase(repo repository.ProductRepository) UpdateProductUsecase {
    return &updateProductUsecase{repo: repo}
}

func (u *updateProductUsecase) Execute(id string, input dto.UpdateProductDTO) (*model.Product, error) {
    product, err := u.repo.GetByID(id)
    if err != nil {
        return nil, err
    }

    product.Name = input.Name
    product.Price = input.Price

    err = u.repo.Update(product)
    if err != nil {
        return nil, err
    }

    return product, nil
}
