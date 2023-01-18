package main

import (
	"log"
	"net"

	"github.com/91diego/grpc/database"
	"github.com/91diego/grpc/server"
	"github.com/91diego/grpc/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	list, err := net.Listen("tcp", ":5065")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPosgresRepository("postgres://postgres:postgres@localhost:54323/postgres?sslmode=disable")
	server := server.NewStudentServer(repo)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
