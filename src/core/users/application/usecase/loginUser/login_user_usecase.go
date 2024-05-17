package loginUser

import (
    "eulabsmyapp/core/users/domain/repository"
    "time"

    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("supersecretkey")

type LoginUserUsecase interface {
    Execute(username, password string) (string, error)
}

type loginUserUsecase struct {
    repo repository.UserRepository
}

func NewLoginUserUsecase(repo repository.UserRepository) LoginUserUsecase {
    return &loginUserUsecase{repo: repo}
}

func (u *loginUserUsecase) Execute(username, password string) (string, error) {
    user, err := u.repo.GetByUsername(username)
    if err != nil {
        return "", err
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", err
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
