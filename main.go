package main

import (
	"fmt"
	"log"
	"net"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gomodul/fn"
	"github.com/mahbubzulkarnain/ex-go-crud/config"
	"github.com/mahbubzulkarnain/ex-go-crud/delivery/pb"
	"github.com/mahbubzulkarnain/ex-go-crud/repository"
	"github.com/mahbubzulkarnain/ex-go-crud/service"
	"google.golang.org/grpc"
)

func main() {
	db, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer fn.Check(db.Close)

	repo := repository.New(db)
	srv := service.New(db, repo)

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, pb.TodoHandler{Srv: srv})

	lis, err := net.Listen("tcp", config.Server.ADDR)
	if err != nil {
		panic(err)
	}

	log.Println(fmt.Sprintf("%v 🚀 server started... ", config.Server.ADDR))
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
