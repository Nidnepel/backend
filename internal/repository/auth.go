package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Nidnepel/backend/internal/database"
	"github.com/Nidnepel/backend/internal/entity"
)

type AuthRepo struct {
	db database.Queryable
}

func NewAuthRepo(db database.Queryable) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) Check(ctx context.Context, login string, password string) (*entity.User, error) {
	query := database.PSQL.
		Select("id",
			"login",
			"password",
			"status",
			"role",
		).
		From(database.TableUser).
		Where(squirrel.Eq{
			"login":    login,
			"password": password,
		})

	var u entity.User
	err := r.db.Get(ctx, &u, query)
	if err != nil {
		return nil, fmt.Errorf("check User: %w", err)
	}
	return &u, nil
}
