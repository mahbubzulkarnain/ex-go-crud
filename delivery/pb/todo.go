package pb

import (
	"context"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
	"github.com/mahbubzulkarnain/ex-go-crud/service"
)

// TodoHandler godoc.
type TodoHandler struct {
	Srv service.Service
}

// List godoc.
func (h TodoHandler) List(ctx context.Context, _ *Empty) (*Todos, error) {
	list, err := h.Srv.TodoList(ctx)
	if err != nil {
		return nil, err
	}
	todos := make([]*Todo, 0)
	for _, i := range list {
		todos = append(todos, &Todo{
			Id:          i.Id,
			Title:       i.Title,
			Description: i.Description,
		})
	}
	return &Todos{Todos: todos}, nil
}

// Create godoc.
func (h TodoHandler) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	todo := req.GetTodo()
	todoID, err := h.Srv.TodoCreate(ctx, model.Todo{
		Title:       todo.GetTitle(),
		Description: todo.GetDescription(),
	})
	if err != nil {
		return nil, err
	}
	todo.Id = todoID
	return &CreateResponse{
		Todo: todo,
	}, nil
}

// Read godoc.
func (h TodoHandler) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	todo, err := h.Srv.TodoRead(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &ReadResponse{
		Todo: &Todo{
			Id:          todo.Id,
			Title:       todo.Title,
			Description: todo.Description,
		},
	}, nil
}

// Update godoc.
func (h TodoHandler) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	todo := req.GetTodo()
	rowsAffected, err := h.Srv.TodoUpdate(ctx, model.Todo{
		Id:          todo.GetId(),
		Title:       todo.GetTitle(),
		Description: todo.GetDescription(),
	})
	if err != nil {
		return nil, err
	}
	return &UpdateResponse{
		Id:           todo.GetId(),
		RowsAffected: rowsAffected,
	}, nil
}

// Delete godoc.
func (h TodoHandler) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	rowsAffected, err := h.Srv.TodoDelete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &DeleteResponse{
		Id:           req.GetId(),
		RowsAffected: rowsAffected,
	}, nil
}
