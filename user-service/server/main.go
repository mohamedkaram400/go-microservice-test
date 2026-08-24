package main

import (
	"log"
	"net"

	"github.com/mohamed/microservices/user-service/handler"
	dbconn "github.com/mohamed/microservices/user-service/conn"
	pb "github.com/mohamed/microservices/user-service/proto"
	"github.com/mohamed/microservices/user-service/repository"
	"github.com/mohamed/microservices/user-service/service"
	"google.golang.org/grpc"
)

func main() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatal(err)
    }

    mysql, err := dbconn.ConnectMySQL(
        "root:qazwsx123@tcp(localhost:3306)/user_service?charset=utf8mb4&parseTime=True&loc=Local",
    )

    if err != nil {
        log.Fatal(err)
    }

    grpcServer := grpc.NewServer()
    userRepository := repository.NewUserRepository(mysql)
    userService := service.NewUserService(userRepository)
    grpcHandler := handler.NewUserHandler(
        userService,
    )

    pb.RegisterUserServiceServer(grpcServer, grpcHandler)


    log.Println("User gRPC Service running on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
} 