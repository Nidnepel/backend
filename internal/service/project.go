package service

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/Nidnepel/backend/internal/repository"
)

type ProjectService struct {
	repo repository.Project
}

func NewProjectService(repo repository.Project) *ProjectService {
	return &ProjectService{repo: repo}
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
