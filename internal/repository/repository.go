package repository

import (
	"context"
	"github.com/Nidnepel/backend/internal/database"
	"github.com/Nidnepel/backend/internal/entity"
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
	Create(ctx context.Context, worker entity.User) (int, error)
	Read(ctx context.Context, workerId int) (*entity.User, error)
}

type Repository struct {
	Authorization
	Project
	Task
	Manager
	Worker
}

func NewRepository(db database.Queryable) *Repository {
	return &Repository{Worker: NewWorkersRepo(db)}
}
