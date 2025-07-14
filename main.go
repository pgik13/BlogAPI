package main

import (
	"bloggers/config"
	"bloggers/database"
	"bloggers/handler"
	"bloggers/repo"
	"bloggers/routes"
	"bloggers/services"
	"fmt"
	"net/http"
	"os"
)

func main() {

	//load variables
	config.LoadEnv()

	//connect to database
	database.ConnectDB()

	//initialise repo
	userRepo := &repo.UserRepo{}
	postRepo := &repo.PostRepo{}
	commentRepo := &repo.CommentRepo{}

	//initialise service
	userService := &services.UserServices{Repo: *userRepo}
	postService := &services.PostService{Repo: *postRepo}
	commentService := &services.CommentService{Repo: *commentRepo}

	//initialise handler
	userHandler := &handler.UserHandler{Service: userService}
	postHandler := &handler.PostHandler{Service: postService}
	commentHandler := &handler.CommentHandler{Service: commentService}

	//define routes
	router := routes.SetupRouter(userHandler, postHandler, commentHandler)

	//start server
	fmt.Println("Server running on " + os.Getenv("DB_HOST"))
	http.ListenAndServe(":8080", router)
}
