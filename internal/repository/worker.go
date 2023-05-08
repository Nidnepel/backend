package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Nidnepel/backend/internal/database"
	"github.com/Nidnepel/backend/internal/entity"
)

type WorkersRepo struct {
	db database.Queryable
}

func NewWorkersRepo(db database.Queryable) *WorkersRepo {
	return &WorkersRepo{db: db}
}

func (r *WorkersRepo) Read(ctx context.Context, workerId int) (*entity.User, error) {
	query := database.PSQL.
		Select("id",
			"login",
			"password",
		).
		From(database.TableUser).
		Where(squirrel.Eq{
			"id": workerId,
		})

	var u entity.User
	err := r.db.Get(ctx, &u, query)
	if err != nil {
		return nil, fmt.Errorf("получение worker: %w", err)
	}
	return &u, nil
}

func (r *WorkersRepo) Create(ctx context.Context, newWorker entity.User) (int, error) {
	query := database.PSQL.
		Insert(database.TableUser).
		Columns(
			"id",
			"login",
			"password",
			"role",
			"status",
		).
		Values(
			newWorker.Id,
			newWorker.Login,
			newWorker.Password,
			newWorker.Role,
			newWorker.Status,
		).
		Suffix("RETURNING id")

	var id int
	err := r.db.Get(ctx, &id, query)
	if err != nil {
		return id, fmt.Errorf("создание worker: %w", err)
	}

	return id, nil
}
