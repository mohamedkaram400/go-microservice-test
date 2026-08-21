package main

import (
	"log"
	"net"

	dbconn "github.com/mohamed/microservices/post-service/conn"
	pb "github.com/mohamed/microservices/user-service/proto"
	"github.com/mohamed/microservices/user-service/service"
	"google.golang.org/grpc"
)

type User struct {
    ID    int       `json:"id"`
    Name  string    `json:"name"`
    Email string    `json:"email"`
}

func main() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatal(err)
    }

    grpcServer := grpc.NewServer()
    userService := service.NewUserService()

    pb.RegisterUserServiceServer(grpcServer, userService)


	dbconn.ConnectMySQL("root:qazwsx123@tcp(localhost:3306)/user_service?charset=utf8mb4&parseTime=True&loc=Local")

    log.Println("User gRPC Service running on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}