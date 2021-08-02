package service

import (
	"context"
	"log"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

// TodoCreate godoc.
func (s service) TodoCreate(ctx context.Context, todo model.Todo) (int64, error) {
	log.Println("TodoCreate", todo)
	return 1, nil
}
