package repository

import "myapp/core/products/domain/model"

type ProductRepository interface {
    Create(product *model.Product) error
    GetByID(id string) (*model.Product, error)
    Update(product *model.Product) error
    Delete(id string) error
}
