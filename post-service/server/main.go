package main

import (
	"log"

	"github.com/gin-gonic/gin"
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
	// Connect to User Service
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
	// Connect to DB
    // ------------------------------------
	mysql, err := dbconn.ConnectMySQL("root:qazwsx123@tcp(localhost:3306)/post_service?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := mysql.DB()
	if err != nil {
		log.Fatal("❌ Failed to get sql.DB: ", err)
	}

	defer sqlDB.Close()

    // ------------------------------------
	// Create User Service gRPC client
	// ------------------------------------

	userClient := userpb.NewUserServiceClient(conn)

	// ------------------------------------
	// Create Post Handler, Service, and Repository
	// ------------------------------------

	postRepository := repository.NewPostRepository(mysql)
	postService := service.NewPostService(postRepository, userClient)
	postHandler := handler.NewPostHandler(postService)

    // ------------------------------------
	// Register REST endpoint
	// ------------------------------------

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger(), gin.Recovery())

	// 7. Versioned API group
	v1 := router.Group("/api/v1")

	v1.GET("/get-post/:id", postHandler.GetPost)
	v1.POST("/create-post", postHandler.CreatePost)

	// ------------------------------------
	// Start REST server
	// ------------------------------------

	log.Println("Post Service REST running on :8082")

	if err := router.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}

