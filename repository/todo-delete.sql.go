package repository

import (
	"context"
	"database/sql"
)

const TodoDeleteQuery = `DELETE FROM todo WHERE id=?`

// TodoDelete godoc.
func (q *Queries) TodoDelete(ctx context.Context, todoID int64) (sql.Result, error) {
	return q.exec(ctx, q.todoDeleteStmt, TodoDeleteQuery, todoID)
}
