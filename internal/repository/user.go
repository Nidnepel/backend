package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Nidnepel/backend/internal/database"
	"github.com/Nidnepel/backend/internal/entity"
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
