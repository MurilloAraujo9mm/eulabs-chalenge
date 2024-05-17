package repository

import "myapp/core/users/domain/model"

type UserRepository interface {
    Create(user *model.User) error
    GetByUsername(username string) (*model.User, error)
}
