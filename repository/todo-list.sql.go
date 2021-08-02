package repository

import (
	"context"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

const TodoListQuery = `SELECT id, title, description FROM todo`

// TodoList godoc.
func (q *Queries) TodoList(ctx context.Context) ([]*model.Todo, error) {
	rows, err := q.query(ctx, q.todoListStmt, TodoListQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]*model.Todo, 0)
	for rows.Next() {
		var i *model.Todo
		if err = rows.Scan(&i.Id, &i.Title, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
