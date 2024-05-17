package repository

import (
    "database/sql"
    "myapp/core/users/model"
    "github.com/google/uuid"
)

type UserRepository interface {
    Create(user *model.User) error
    GetByUsername(username string) (*model.User, error)
}

type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
    user.ID = uuid.New().String()
    _, err := r.db.Exec("INSERT INTO users (id, username, password) VALUES (?, ?, ?)", user.ID, user.Username, user.Password)
    return err
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
    user := &model.User{}
    err := r.db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
    return user, err
}
