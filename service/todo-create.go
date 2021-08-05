package service

import (
	"context"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

// TodoCreate godoc.
func (s service) TodoCreate(ctx context.Context, todo model.Todo) (lastInsertId int64, err error) {
	todoCreate, err := s.repo.TodoCreate(ctx, todo)
	if err != nil {
		return 0, err
	}
	return todoCreate.LastInsertId()
}
