package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/gomodul/fn"
	"github.com/mahbubzulkarnain/ex-go-crud/config"
	"github.com/mahbubzulkarnain/ex-go-crud/delivery/pb"
	"github.com/mahbubzulkarnain/ex-go-crud/repository"
	"github.com/mahbubzulkarnain/ex-go-crud/service"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open(config.DB.MySQL.DriverName, config.DB.MySQL.DataSourceName)
	if err != nil {
		panic(err)
	}
	defer fn.Check(db.Close)

	db.SetMaxIdleConns(config.DB.MySQL.MaxIdleConns)
	db.SetMaxOpenConns(config.DB.MySQL.MaxOpenConns)
	db.SetConnMaxLifetime(config.DB.MySQL.ConnMaxLifetime)

	if err = db.Ping(); err != nil {
		panic(err)
	}

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
