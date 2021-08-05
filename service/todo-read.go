package service

import (
	"context"
	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

// TodoRead godoc.
func (s service) TodoRead(ctx context.Context, todoID int64) (model.Todo, error) {
	return s.repo.TodoRead(ctx, todoID)
}
