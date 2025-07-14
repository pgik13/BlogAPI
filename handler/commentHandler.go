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

type CommentHandler struct {
	Service *services.CommentService
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}

	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "failed to get userID", 500)
		return
	}

	comment.UserID = userID

	err = h.Service.CreateComment(&comment)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(comment)
}

func (h *CommentHandler) GetCommentByPostID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "post ID required", http.StatusBadRequest)
	}

	idConv, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid post ID", http.StatusBadRequest)
	}

	comments, err := h.Service.GetCommentByPostID(uint(idConv))
	if err != nil {
		http.Error(w, "could not load comments", 400)
		return
	}

	var response []models.CommentContent
	for _, c := range comments {
		response = append(response, models.CommentContent{Username: c.User.Username, Content: c.Content})
	}

	json.NewEncoder(w).Encode(response)
}
