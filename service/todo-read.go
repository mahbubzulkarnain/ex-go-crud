package service

import (
	"context"
	"log"

	"github.com/bxcodec/faker/v3"
	"github.com/mahbubzulkarnain/ex-go-crud/model"
)

func (s service) TodoRead(ctx context.Context, id int64) (*model.Todo, error) {
	var todo *model.Todo
	if err := faker.FakeData(&todo); err != nil {
		return nil, err
	}
	log.Println("TodoRead", id)
	return todo, nil
}
