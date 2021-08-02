package repository

import (
	"context"
	"database/sql"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

const TodoUpdateQuery = `UPDATE todo SET title=?, description=? WHERE id=?`

// TodoUpdate godoc.
func (q *Queries) TodoUpdate(ctx context.Context, arg model.Todo) (sql.Result, error) {
	return q.exec(ctx, q.todoUpdateStmt, TodoUpdateQuery, arg.Title, arg.Description, arg.Id)
}
