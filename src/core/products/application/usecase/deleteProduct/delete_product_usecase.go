package deleteProduct

import "myapp/core/products/domain/repository"

type DeleteProductUsecase interface {
    Execute(id string) error
}

type deleteProductUsecase struct {
    repo repository.ProductRepository
}

func NewDeleteProductUsecase(repo repository.ProductRepository) DeleteProductUsecase {
    return &deleteProductUsecase{repo: repo}
}

func (u *deleteProductUsecase) Execute(id string) error {
    return u.repo.Delete(id)
}
