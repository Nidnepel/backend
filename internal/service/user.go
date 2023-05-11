package service

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/Nidnepel/backend/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user entity.User) (int, error) {
	return s.repo.Create(ctx, user)
}

func (s *UserService) ReadUser(ctx context.Context, id int) (*entity.User, error) {
	return s.repo.Read(ctx, id)
}

func (s *UserService) ReadAllUsers(ctx context.Context) ([]*entity.User, error) {
	return s.repo.ReadAll(ctx)
}
