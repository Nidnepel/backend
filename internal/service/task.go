package service

import (
	"context"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/Nidnepel/backend/internal/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, task entity.Task) (int, error) {
	return s.repo.CreateTask(ctx, task)
}
func (s *TaskService) ReadTask(ctx context.Context, id int) (*entity.Task, error) {
	return s.repo.ReadTask(ctx, id)
}

func (s *TaskService) CloseTask(ctx context.Context, id int) (bool, error) {
	return s.repo.Close(ctx, id)
}

func (s *TaskService) CreateReport(ctx context.Context, taskId int, newReport entity.TaskReport) (int, error) {
	id, err := s.repo.CreateReport(ctx, newReport)
	if err != nil {
		return 0, err
	}
	err = s.repo.AddTaskReportForTask(ctx, taskId, id)
	return id, err
}

func (s *TaskService) ReadAllReports(ctx context.Context, taskId int) ([]*entity.TaskReport, error) {
	return s.repo.ReadReports(ctx, taskId)
}
