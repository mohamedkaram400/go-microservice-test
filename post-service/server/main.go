package main

import (
	"log"
	"net/http"
	userpb "github.com/mohamed/microservices/user-service/proto"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/mohamed/microservices/post-service/handler"
	"google.golang.org/grpc"
)


func main() {
    // ------------------------------------
	// 1. Connect to User Service
	// ------------------------------------

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

    // ------------------------------------
	// 2. Create User Service gRPC client
	// ------------------------------------

	userClient := userpb.NewUserServiceClient(conn)

	// ------------------------------------
	// 3. Create Post Handler
	// ------------------------------------

	postHandler := handler.NewPostHandler(userClient)

    // ------------------------------------
	// 4. Register REST endpoint
	// ------------------------------------

	http.HandleFunc("/posts/1", postHandler.GetPost)

	// ------------------------------------
	// 5. Start REST server
	// ------------------------------------

	log.Println("Post Service REST running on :8082")

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}

