package registerUser

import (
    "myapp/core/users/domain/model"
    "myapp/core/users/domain/repository"
    "myapp/go_modules/users/dto"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
)

type RegisterUserUsecase interface {
    Execute(input dto.RegisterDTO) (*model.User, error)
}

type registerUserUsecase struct {
    repo repository.UserRepository
}

func NewRegisterUserUsecase(repo repository.UserRepository) RegisterUserUsecase {
    return &registerUserUsecase{repo: repo}
}

func (u *registerUserUsecase) Execute(input dto.RegisterDTO) (*model.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &model.User{
        ID:       uuid.New().String(),
        Username: input.Username,
        Password: string(hashedPassword),
    }

    err = u.repo.Create(user)
    if err != nil {
        return nil, err
    }
    return user, nil
}
