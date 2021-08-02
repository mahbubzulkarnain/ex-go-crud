package repository

import (
	"context"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

const ReadQuery = `SELECT id, title, description FROM todo WHERE id=?`

// TodoRead godoc.
func (q *Queries) TodoRead(ctx context.Context, todoID string) (model.Todo, error) {
	row := q.queryRow(ctx, q.todoReadStmt, ReadQuery, todoID)
	var i model.Todo
	err := row.Scan(
		&i.Id,
		&i.Title,
		&i.Description,
	)
	return i, err
}
