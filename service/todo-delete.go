package service

import (
	"context"
)

// TodoDelete godoc.
func (s service) TodoDelete(ctx context.Context, todoID int64) (rowsAffected int64, err error) {
	todoDelete, err := s.repo.TodoDelete(ctx, todoID)
	if err != nil {
		return 0, err
	}
	return todoDelete.RowsAffected()
}
