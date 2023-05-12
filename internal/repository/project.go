package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Nidnepel/backend/internal/database"
	"github.com/Nidnepel/backend/internal/entity"
)

type ProjectsRepo struct {
	db database.Queryable
}

func NewProjectsRepo(db database.Queryable) *ProjectsRepo {
	return &ProjectsRepo{db: db}
}

func (r *ProjectsRepo) Read(ctx context.Context, projectId int) (*entity.Project, error) {
	query := database.PSQL.
		Select("id",
			"title",
			"description",
			"status",
		).
		From(database.TableProject).
		Where(squirrel.Eq{
			"id": projectId,
		})

	var u entity.Project
	err := r.db.Get(ctx, &u, query)
	if err != nil {
		return nil, fmt.Errorf("получение Project: %w", err)
	}
	return &u, nil
}

func (r *ProjectsRepo) Close(ctx context.Context, projectId int) (bool, error) {
	query := database.PSQL.Update(database.TableProject).
		Set("status", false).
		Where(squirrel.Eq{
			"id": projectId,
		})

	result, err := r.db.Exec(ctx, query)
	cntUpdate, _ := result.RowsAffected()
	return cntUpdate > 0, err
}

func (r *ProjectsRepo) Create(ctx context.Context, newProject entity.Project) (int, error) {
	query := database.PSQL.
		Insert(database.TableProject).
		Columns(
			"title",
			"description",
		).
		Values(
			newProject.Title,
			newProject.Description,
		).
		Suffix("RETURNING id")

	var id int
	err := r.db.Get(ctx, &id, query)
	if err != nil {
		return id, fmt.Errorf("создание Project: %w", err)
	}

	return id, nil
}

func (r *ProjectsRepo) ReadAll(ctx context.Context) ([]*entity.Project, error) {
	var items []*entity.Project
	query := database.PSQL.Select(
		"id",
		"title",
		"description",
		"status",
	).From(database.TableProject)
	err := r.db.Select(ctx, &items, query)
	if err != nil {
		return nil, fmt.Errorf("получение all Project: %w", err)
	}
	return items, nil
}

func (r *ProjectsRepo) AddUser(ctx context.Context, userId, projectId int) (bool, error) {
	query := database.PSQL.
		Insert(database.TableUserProjectList).
		Columns(
			"user_id",
			"project_id",
		).
		Values(
			userId,
			projectId,
		)

	_, err := r.db.Exec(ctx, query)
	if err != nil {
		return false, fmt.Errorf("add User in project: %w", err)
	}

	return true, nil
}

func (r *ProjectsRepo) ReadAllUsers(ctx context.Context, projectId int) ([]*entity.User, error) {
	var items []*entity.User
	query := database.PSQL.Select(
		"DISTINCT users.id",
		"login",
		"password",
		"role",
		"status",
	).From(database.TableUserProjectList).Join(database.TableUser + " ON user_project_list.user_id = users.id").Where(squirrel.Eq{
		"project_id": projectId,
	})
	err := r.db.Select(ctx, &items, query)
	if err != nil {
		return nil, fmt.Errorf("получение all Users in Project: %w", err)
	}
	return items, nil
}
