package service

import (
	"context"
	"log"

	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

func (s service) TodoUpdate(ctx context.Context, todo model.Todo) (int64, error) {
	log.Println("TodoUpdate", todo)
	return 1, nil
}
