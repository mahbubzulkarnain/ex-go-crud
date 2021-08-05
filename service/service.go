package service

import (
	"context"
	"database/sql"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
	"github.com/mahbubzulkarnain/ex-go-crud/repository"
)

// Service godoc.
type Service interface {
	TodoList(ctx context.Context) ([]*model.Todo, error)
	TodoCreate(ctx context.Context, todo model.Todo) (lastInsertId int64, err error)
	TodoRead(ctx context.Context, todoID int64) (model.Todo, error)
	TodoUpdate(ctx context.Context, todo model.Todo) (rowsAffected int64, err error)
	TodoDelete(ctx context.Context, todoID int64) (rowsAffected int64, err error)
}

type service struct {
	db   *sql.DB
	repo *repository.Queries
}

// New godoc.
func New(db *sql.DB, repo *repository.Queries) Service {
	return &service{db: db, repo: repo}
}
