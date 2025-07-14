package handler

import (
	"bloggers/middleware"
	"bloggers/models"
	"bloggers/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PostHandler struct {
	Service *services.PostService
}

// handles HTTP POST requests to create a new post.
// expects a JSON body, extracts the user ID from the auth token and stores the new post in the database.
func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post // Declare a variable to hold the decoded JSON request body

	//Decode the JSON request body into the post struct
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}

	//Retrieve the user ID from the JWT token or session via middleware
	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "could not get user ID", 500)
		return
	}

	post.UserID = userID

	//Call the repository method to save the post to the database - method created in the service layer
	err = h.Service.Repo.CreatePost(&post) //try without .repo
	if err != nil {
		http.Error(w, "could not create post", 500)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {

	//using gorilla/mux to its potential
	vars := mux.Vars(r)
	id := vars["id"]

	//id check
	if id == "" {
		http.Error(w, "post ID is required", 400)
		return
	}

	//try to convert id to string
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid post ID", 400) //400 might be better if cannot convert - it was, this is a client error not server error
		return
	}

	//use converted id to get post from service layer
	post, err := h.Service.GetPostByID(uint(idInt))
	if err != nil {
		http.Error(w, "could not fetch post", 400)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) EditPost(w http.ResponseWriter, r *http.Request) {

	//Get postID from URL path
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "post ID is required", 400)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid post ID", 400)
		return
	}

	post, err := h.Service.GetPostByID(uint(idInt))
	if err != nil {
		http.Error(w, "post not found", http.StatusNotFound) //404
	}

	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "unauthorised", http.StatusUnauthorized) //401
		return
	}

	if post.UserID != userID {
		http.Error(w, "You do not have permission to edit this post", http.StatusForbidden) //403
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "invlaid request body", 400)
		return
	}

	err = h.Service.EditPost(uint(idInt), updates)
	if err != nil {
		http.Error(w, "could not update post", 500)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Post updated"))
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {

	var post models.Post

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "post ID is required", 400)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid post ID", 400)
		return
	}

	posts, err := h.Service.GetPostByID(uint(idInt))
	if err != nil {
		http.Error(w, "post not found", 404)
	}

	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "unauthorised", http.StatusUnauthorized)
		return
	}

	if posts.UserID != userID {
		http.Error(w, "You do not have permission to edit this post", http.StatusForbidden)
	}

	err = h.Service.DeletePost(&post, uint(idInt))
	if err != nil {
		http.Error(w, "unable to delete post", 500)
	}

	w.WriteHeader(200)
	w.Write([]byte("post deleted"))
}
