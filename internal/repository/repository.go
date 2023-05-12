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
	Create(ctx context.Context, project entity.Project) (int, error)
	Read(ctx context.Context, projectId int) (*entity.Project, error)
	Close(ctx context.Context, projectId int) (bool, error)
	ReadAll(ctx context.Context) ([]*entity.Project, error)
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
	return &Repository{User: NewUsersRepo(db), Authorization: NewAuthRepo(db), Project: NewProjectsRepo(db)}
}
