package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Nidnepel/backend/internal/database"
	"github.com/Nidnepel/backend/internal/entity"
	"github.com/jackc/pgx/pgtype"
	"time"
)

type UsersRepo struct {
	db database.Queryable
}

func NewUsersRepo(db database.Queryable) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) Read(ctx context.Context, userId int) (*entity.User, error) {
	query := database.PSQL.
		Select("id",
			"login",
			"password",
			"status",
			"role",
		).
		From(database.TableUser).
		Where(squirrel.Eq{
			"id": userId,
		})

	var u entity.User
	err := r.db.Get(ctx, &u, query)
	if err != nil {
		return nil, fmt.Errorf("получение User: %w", err)
	}
	return &u, nil
}

func (r *UsersRepo) Create(ctx context.Context, newUser entity.User) (int, error) {
	query := database.PSQL.
		Insert(database.TableUser).
		Columns(
			"login",
			"password",
			"role",
		).
		Values(
			newUser.Login,
			newUser.Password,
			newUser.Role,
		).
		Suffix("RETURNING id")

	var id int
	err := r.db.Get(ctx, &id, query)
	if err != nil {
		return id, fmt.Errorf("создание User: %w", err)
	}

	return id, nil
}

func (r *UsersRepo) ReadAll(ctx context.Context) ([]*entity.User, error) {
	var items []*entity.User
	query := database.PSQL.Select("id",
		"login",
		"password",
		"status",
		"role",
	).From(database.TableUser)
	err := r.db.Select(ctx, &items, query)
	if err != nil {
		return nil, fmt.Errorf("получение all Users: %w", err)
	}
	return items, nil
}

func (r *UsersRepo) DeleteUserInProject(ctx context.Context, projectId, userId int) (bool, error) {
	query := database.PSQL.Delete(database.TableUserProjectList).
		Where(squirrel.Eq{
			"user_id":    userId,
			"project_id": projectId,
		})
	result, err := r.db.Exec(ctx, query)
	if err != nil {
		return false, fmt.Errorf("delete User from Project: %w", err)
	}
	cntUpdateRows, _ := result.RowsAffected()
	return cntUpdateRows > 0, nil
}

func (r *UsersRepo) AddTaskInProject(ctx context.Context, projectId, userId, taskId int) error {
	query := database.PSQL.
		Insert(database.TableUserTaskList).
		Columns(
			"project_id",
			"user_id",
			"task_id",
		).
		Values(
			projectId,
			userId,
			taskId,
		)

	_, err := r.db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("added task in project: %w", err)
	}

	return nil
}

func (r *UsersRepo) ReadAllProjects(ctx context.Context, userId int) ([]*entity.Project, error) {
	var items []*entity.Project
	query := database.PSQL.Select(
		"DISTINCT projects.id",
		"title",
		"description",
		"status",
	).From(database.TableUserProjectList).Join(database.TableProject +
		" ON user_project_list.project_id = projects.id").Where(squirrel.Eq{
		"user_id": userId,
	})
	err := r.db.Select(ctx, &items, query)
	if err != nil {
		return nil, fmt.Errorf("getting all Projects for User: %w", err)
	}
	return items, nil
}

func (r *UsersRepo) GetTasks(ctx context.Context, projectId, userId int) ([]*entity.Task, error) {
	var items []*entity.Task
	query := database.PSQL.Select(
		"DISTINCT tasks.id",
		"title",
		"description",
		"progress_status",
	).From(database.TableUserTaskList).Join(database.TableTask +
		" ON user_task_list.task_id = tasks.id").Where(squirrel.Eq{
		"project_id": projectId,
		"user_id":    userId,
	})
	err := r.db.Select(ctx, &items, query)
	if err != nil {
		return nil, fmt.Errorf("getting all Tasks for User in Project: %w", err)
	}
	return items, nil
}

func (r *UsersRepo) GetActivity(ctx context.Context, projectId, userId int) (*entity.Session, error) {
	query := database.PSQL.
		Select("id",
			"project_id",
			"user_id",
			"keylog",
			"screens",
			"start",
			"finish",
		).
		From(database.TableSession).
		Where(squirrel.Eq{
			"project_id": projectId,
			"user_id":    userId,
		})

	var u entity.Session
	err := r.db.Get(ctx, &u, query)
	if err != nil {
		return nil, fmt.Errorf("get Activity: %w", err)
	}
	return &u, nil
}

func (r *UsersRepo) CreateSession(ctx context.Context, session entity.Session) (int, error) {
	query := database.PSQL.
		Insert(database.TableSession).
		Columns("project_id",
			"user_id",
			"keylog",
			"screens",
			"start",
			"finish",
		).Values(
		session.ProjectId,
		session.UserId,
		session.Keylog,
		session.Screens,
		pgtype.Timestamp{Time: time.Now()},
		pgtype.Timestamp{Time: time.Now()},
	).Suffix("RETURNING id")

	var id int
	err := r.db.Get(ctx, &id, query)
	if err != nil {
		return id, fmt.Errorf("send Session: %w", err)
	}

	return id, nil
}
