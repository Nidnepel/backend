package service

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/Nidnepel/backend/internal/repository"
)

type Authorization interface {
	CheckUser(ctx context.Context, login string, password string) (*entity.User, error)
}

type Project interface {
}

type Task interface {
}

type User interface {
	CreateUser(ctx context.Context, worker entity.User) (int, error)
	ReadUser(ctx context.Context, id int) (*entity.User, error)
	ReadAllUsers(ctx context.Context) ([]*entity.User, error)
}

type Service struct {
	Authorization
	Project
	Task
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{User: NewUserService(repos.User), Authorization: NewAuthService(repos.Authorization)}
}
