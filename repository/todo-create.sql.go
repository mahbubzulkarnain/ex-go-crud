package repository

import (
	"context"
	"database/sql"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

const TodoCreateQuery = `INSERT INTO todo(title, description) VALUES(?,?);`

type TodoCreateParams struct {
	Title       string
	Description string
}

// TodoCreate godoc.
func (q *Queries) TodoCreate(ctx context.Context, arg model.Todo) (sql.Result, error) {
	return q.exec(ctx, q.todoCreateStmt, TodoCreateQuery, arg.Title, arg.Description)
}
