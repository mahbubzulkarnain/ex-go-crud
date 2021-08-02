package service

import (
	"context"
	"database/sql"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
	"github.com/mahbubzulkarnain/ex-go-crud/repository"
)

type Service interface {
	TodoList(ctx context.Context) ([]*model.Todo, error)
	TodoCreate(ctx context.Context, todo model.Todo) (int64, error)
	TodoRead(ctx context.Context, id int64) (*model.Todo, error)
	TodoUpdate(ctx context.Context, todo model.Todo) (int64, error)
	TodoDelete(ctx context.Context, id int64) (int64, error)
}

type service struct {
	db   *sql.DB
	repo *repository.Queries
}

func New(db *sql.DB, repo *repository.Queries) Service {
	return &service{db: db, repo: repo}
}
