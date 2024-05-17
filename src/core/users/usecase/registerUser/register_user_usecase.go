package registerUser

import (
    "myapp/core/users/dto"
    "myapp/core/users/model"
    "myapp/core/users/repository"
    "golang.org/x/crypto/bcrypt"
)

type RegisterUserUsecase interface {
    Execute(user *dto.RegisterDTO) error
}

type registerUserUsecase struct {
    repo repository.UserRepository
}

func NewRegisterUserUsecase(repo repository.UserRepository) RegisterUserUsecase {
    return &registerUserUsecase{repo: repo}
}

func (u *registerUserUsecase) Execute(user *dto.RegisterDTO) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    newUser := &model.User{
        Username: user.Username,
        Password: string(hashedPassword),
    }

    return u.repo.Create(newUser)
}
