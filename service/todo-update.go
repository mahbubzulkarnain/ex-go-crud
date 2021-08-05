package service

import (
	"context"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

// TodoUpdate godoc.
func (s service) TodoUpdate(ctx context.Context, todo model.Todo) (rowsAffected int64, err error) {
	todoUpdate, err := s.repo.TodoUpdate(ctx, todo)
	if err != nil {
		return 0, err
	}
	return todoUpdate.RowsAffected()
}
