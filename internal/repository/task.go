package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Nidnepel/backend/internal/database"
	"github.com/Nidnepel/backend/internal/entity"
)

type TasksRepo struct {
	db database.Queryable
}

func NewTasksRepo(db database.Queryable) *TasksRepo {
	return &TasksRepo{db: db}
}

func (r *TasksRepo) CreateTask(ctx context.Context, newTask entity.Task) (int, error) {
	query := database.PSQL.
		Insert(database.TableTask).
		Columns(
			"title",
			"description",
		).
		Values(
			newTask.Title,
			newTask.Description,
		).
		Suffix("RETURNING id")

	var id int
	err := r.db.Get(ctx, &id, query)
	if err != nil {
		return id, fmt.Errorf("создание Task: %w", err)
	}

	return id, nil
}

func (r *TasksRepo) ReadTask(ctx context.Context, id int) (*entity.Task, error) {
	query := database.PSQL.
		Select("id",
			"title",
			"description",
			"progress_status",
		).
		From(database.TableTask).
		Where(squirrel.Eq{
			"id": id,
		})

	var u entity.Task
	err := r.db.Get(ctx, &u, query)
	if err != nil {
		return nil, fmt.Errorf("получение Task: %w", err)
	}
	return &u, nil
}

func (r *TasksRepo) Close(ctx context.Context, taskId int) (bool, error) {
	query := database.PSQL.Update(database.TableTask).
		Set("progress_status", false).
		Where(squirrel.Eq{
			"id": taskId,
		})

	result, err := r.db.Exec(ctx, query)
	cntUpdate, _ := result.RowsAffected()
	return cntUpdate > 0, err
}

func (r *TasksRepo) CreateReport(ctx context.Context, newReport entity.TaskReport) (int, error) {
	query := database.PSQL.
		Insert(database.TableReports).
		Columns(
			"title",
			"description",
		).
		Values(
			newReport.Title,
			newReport.Description,
		).
		Suffix("RETURNING id")

	var id int
	err := r.db.Get(ctx, &id, query)
	if err != nil {
		return id, fmt.Errorf("создание Report: %w", err)
	}

	return id, nil
}

func (r *TasksRepo) AddTaskReportForTask(ctx context.Context, taskId, id int) error {
	query := database.PSQL.
		Insert(database.TableTaskReportList).
		Columns(
			"task_id",
			"report_id",
		).
		Values(
			taskId,
			id,
		)

	_, err := r.db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("added report in task: %w", err)
	}

	return nil
}

func (r *TasksRepo) ReadReports(ctx context.Context, taskId int) ([]*entity.TaskReport, error) {
	var items []*entity.TaskReport
	query := database.PSQL.Select(
		"DISTINCT task_reports.id",
		"title",
		"description",
	).From(database.TableTaskReportList).Join(database.TableReports +
		" ON task_report_list.report_id = task_reports.id").Where(squirrel.Eq{
		"task_id": taskId,
	})
	err := r.db.Select(ctx, &items, query)
	if err != nil {
		return nil, fmt.Errorf("getting all Reports for Task: %w", err)
	}
	return items, nil
}
