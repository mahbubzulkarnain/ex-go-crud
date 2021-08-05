package service

import (
	"context"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

// TodoList godoc.
func (s service) TodoList(ctx context.Context) ([]*model.Todo, error) {
	return s.repo.TodoList(ctx)
}
