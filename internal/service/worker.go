package service

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/Nidnepel/backend/internal/repository"
)

type WorkerService struct {
	repo repository.Worker
}

func NewWorkerService(repo repository.Worker) *WorkerService {
	return &WorkerService{repo: repo}
}

func (s *WorkerService) CreateWorker(ctx context.Context, worker entity.User) (int, error) {
	return s.repo.Create(ctx, worker)
}

func (s *WorkerService) ReadWorker(ctx context.Context, id int) (*entity.User, error) {
	return s.repo.Read(ctx, id)
}
