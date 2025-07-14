package routes

import (
	"bloggers/handler"
	"bloggers/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handler.UserHandler, postHandler *handler.PostHandler, commentHandler *handler.CommentHandler) *mux.Router {
	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LogIn).Methods("POST")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	//Editor-only routes
	protected.Handle("/posts", middleware.EditorOnlyMiddleware(http.HandlerFunc(postHandler.CreatePost))).Methods("POST")
	protected.Handle("/posts/{id}", middleware.EditorOnlyMiddleware(http.HandlerFunc(postHandler.EditPost))).Methods("PATCH") //instead of put because we won't always update the whole post
	protected.Handle("/delete/{id}", middleware.EditorOnlyMiddleware(http.HandlerFunc(postHandler.DeletePost))).Methods("DELETE")

	//Authenticated user routes
	protected.HandleFunc("/post/{id}", postHandler.GetPostByID).Methods("GET")
	protected.HandleFunc("/comment", commentHandler.CreateComment).Methods("POST")
	protected.HandleFunc("/post/{id}/comments", commentHandler.GetCommentByPostID).Methods("GET")
	return r

}
