package service

import (
	"context"
	"log"

	"github.com/bxcodec/faker/v3"
	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

func (s service) TodoList(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := faker.FakeData(&todos); err != nil {
		return nil, err
	}
	log.Println("TodoList", todos)
	return todos, nil
}
