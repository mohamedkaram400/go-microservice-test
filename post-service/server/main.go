package main

import (
	"log"
	"net/http"

	userpb "github.com/mohamed/microservices/user-service/proto"
	"google.golang.org/grpc/credentials/insecure"

	dbconn "github.com/mohamed/microservices/post-service/conn"
	"github.com/mohamed/microservices/post-service/handler"
	"github.com/mohamed/microservices/post-service/repository"
	"github.com/mohamed/microservices/post-service/service"
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


	// Connect to DB
	mysql, err := dbconn.ConnectMySQL("root:qazwsx123@tcp(localhost:3306)/post_service?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := mysql.DB()
	if err != nil {
		log.Fatal("❌ Failed to get sql.DB: ", err)
	}
    // ------------------------------------
	// 2. Create User Service gRPC client
	// ------------------------------------

	userClient := userpb.NewUserServiceClient(conn)

	// ------------------------------------
	// 3. Create Post Handler, Service, and Repository
	// ------------------------------------

	postRepository := repository.NewPostRepository(sqlDB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(userClient, postService)

    // ------------------------------------
	// 4. Register REST endpoint
	// ------------------------------------

	http.HandleFunc("/posts/", postHandler.GetAllPosts)
	http.HandleFunc("/posts/:id", postHandler.GetPost)
	http.HandleFunc("/posts", postHandler.CreatePost)
	http.HandleFunc("/posts/:id", postHandler.DeletePost)

	// ------------------------------------
	// 5. Start REST server
	// ------------------------------------

	log.Println("Post Service REST running on :8082")

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}

