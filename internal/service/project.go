package service

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/Nidnepel/backend/internal/repository"
)

type ProjectService struct {
	repo     repository.Project
	userRepo repository.User
}

func NewProjectService(repo repository.Project, userRepo repository.User) *ProjectService {
	return &ProjectService{repo: repo, userRepo: userRepo}
}

func (s *ProjectService) CreateProject(ctx context.Context, project entity.Project) (int, error) {
	return s.repo.Create(ctx, project)
}

func (s *ProjectService) ReadProject(ctx context.Context, id int) (*entity.Project, error) {
	return s.repo.Read(ctx, id)
}

func (s *ProjectService) CloseProject(ctx context.Context, id int) (bool, error) {
	return s.repo.Close(ctx, id)
}

func (s *ProjectService) ReadAllProjects(ctx context.Context) ([]*entity.Project, error) {
	return s.repo.ReadAll(ctx)
}

func (s *ProjectService) AddUserInProject(ctx context.Context, projectId, userId int) (bool, error) {
	user, err := s.userRepo.Read(ctx, userId)
	if err != nil || user.Role == entity.AdminRole || user.Status == entity.NotActiveUser {
		return false, err
	}
	isOk, err := s.repo.AddUser(ctx, userId, projectId)
	return isOk, err
}

func (s *ProjectService) ReadAllUsers(ctx context.Context, projectId int) ([]*entity.User, error) {
	return s.repo.ReadAllUsers(ctx, projectId)
}
