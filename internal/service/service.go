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
	CreateProject(ctx context.Context, project entity.Project) (int, error)
	ReadProject(ctx context.Context, id int) (*entity.Project, error)
	CloseProject(ctx context.Context, id int) (bool, error)
	ReadAllProjects(ctx context.Context) ([]*entity.Project, error)
	ReadAllUsers(ctx context.Context, projectId int) ([]*entity.User, error)
	AddUserInProject(ctx context.Context, projectId, userId int) (bool, error)
}

type Task interface {
}

type User interface {
	CreateUser(ctx context.Context, worker entity.User) (int, error)
	ReadUser(ctx context.Context, id int) (*entity.User, error)
	ReadAllUsers(ctx context.Context) ([]*entity.User, error)
	DeleteUserInProject(ctx context.Context, projectId, userId int) (bool, error)
	ReadAllProjects(ctx context.Context, userId int) ([]*entity.Project, error)
	CreateTask(ctx context.Context, projectId, userId int, newTask entity.Task) (int, error)
	GetTasksInProject(ctx context.Context, projectId, userId int) ([]*entity.Task, error)
	GetUserActivityInProject(ctx context.Context, projectId, userId int) (*entity.Session, error)
	CreateSession(ctx context.Context, session entity.Session) (int, error)
}

type Service struct {
	Authorization
	Project
	Task
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:          NewUserService(repos.User),
		Authorization: NewAuthService(repos.Authorization),
		Project:       NewProjectService(repos.Project, repos.User),
	}
}
