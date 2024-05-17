package updateProduct

import (
    "myapp/core/products/model"
    "myapp/core/products/repository"
    "myapp/core/products/dto"
)

type UpdateProductUsecase interface {
    Execute(id string, dto dto.UpdateProductDTO) (*model.Product, error)
}

type updateProductUsecase struct {
    repo repository.ProductRepository
}

func NewUpdateProductUsecase(repo repository.ProductRepository) UpdateProductUsecase {
    return &updateProductUsecase{repo: repo}
}

func (u *updateProductUsecase) Execute(id string, dto dto.UpdateProductDTO) (*model.Product, error) {
    product := &model.Product{
        Name:  dto.Name,
        Price: dto.Price,
    }

    err := u.repo.Update(id, product)
    if err != nil {
        return nil, err
    }
    return product, nil
}
