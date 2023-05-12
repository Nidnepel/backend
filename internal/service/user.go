package service

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/Nidnepel/backend/internal/repository"
)

type UserService struct {
	repo     repository.User
	taskRepo repository.Task
}

func NewUserService(repo repository.User, taskRepo repository.Task) *UserService {
	return &UserService{
		repo:     repo,
		taskRepo: taskRepo,
	}
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

func (s *UserService) DeleteUserInProject(ctx context.Context, projectId, userId int) (bool, error) {
	return s.repo.DeleteUserInProject(ctx, projectId, userId)
}

func (s *UserService) ReadAllProjects(ctx context.Context, userId int) ([]*entity.Project, error) {
	return s.repo.ReadAllProjects(ctx, userId)
}

func (s *UserService) CreateTask(ctx context.Context, projectId, userId int, newTask entity.Task) (int, error) {
	id, err := s.taskRepo.CreateTask(ctx, newTask)
	if err != nil {
		return 0, err
	}
	err = s.repo.AddTaskInProject(ctx, projectId, userId, id)
	return id, err
}

func (s *UserService) GetTasksInProject(ctx context.Context, projectId, userId int) ([]*entity.Task, error) {
	return s.repo.GetTasks(ctx, projectId, userId)
}

func (s *UserService) GetUserActivityInProject(ctx context.Context, projectId, userId int) (*entity.Session, error) {
	return s.repo.GetActivity(ctx, projectId, userId)
}

func (s *UserService) CreateSession(ctx context.Context, newSession entity.Session) (int, error) {
	return s.repo.CreateSession(ctx, newSession)
}
