package repository

import (
	"context"
	"github.com/Nidnepel/backend/internal/database"
	"github.com/Nidnepel/backend/internal/entity"
)

type Authorization interface {
	Check(ctx context.Context, login string, password string) (*entity.User, error)
}

type Project interface {
}

type Task interface {
}

type User interface {
	Create(ctx context.Context, user entity.User) (int, error)
	Read(ctx context.Context, userId int) (*entity.User, error)
	ReadAll(ctx context.Context) ([]*entity.User, error)
}

type Repository struct {
	Authorization
	Project
	Task
	User
}

func NewRepository(db database.Queryable) *Repository {
	return &Repository{User: NewUsersRepo(db), Authorization: NewAuthRepo(db)}
}
