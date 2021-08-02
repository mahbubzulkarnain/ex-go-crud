package service

import (
	"context"
	"log"
)

func (s service) TodoDelete(ctx context.Context, id int64) (int64, error) {
	log.Println("TodoDelete", id)
	return 1, nil
}
