package service

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/Nidnepel/backend/internal/repository"
)

type Authorization interface {
}

type Project interface {
}

type Task interface {
}

type Manager interface {
}

type Worker interface {
	CreateWorker(ctx context.Context, worker entity.User) (int, error)
	ReadWorker(ctx context.Context, id int) (*entity.User, error)
}

type Service struct {
	Authorization
	Project
	Task
	Manager
	Worker
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Worker: NewWorkerService(repos.Worker)}
}
