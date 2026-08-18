package main

import (
	"log"
	"net"

	pb "github.com/mohamed/microservices/user-service/proto"
	"github.com/mohamed/microservices/user-service/service"
	"google.golang.org/grpc"
)

type User struct {
    ID   int
    Name string
    Email string
}

func main() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatal(err)
    }

    grpcServer := grpc.NewServer()
    userService := service.NewUserService()

    pb.RegisterUserServiceServer(grpcServer, userService)

    log.Println("User gRPC Service running on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}