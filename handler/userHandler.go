package handler

import (
	"bloggers/models"
	"bloggers/services"
	"encoding/json"
	"net/http"
)

//handler layer - parse http requests & responses and call the service layer
//
// service layer - business logic and calls the repo layer
//
// repo layer - direct database operation

type UserHandler struct {
	Service *services.UserServices
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	//call the service layer
	err = h.Service.RegisterUser(&user)
	if err != nil {
		http.Error(w, "could not register user", http.StatusBadRequest)
		return
	}

	//response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	var request models.LoginRequest
	err := json.NewDecoder(r.Body).Decode((&request))
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.Service.LogInUser(request)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusInternalServerError)
		return
	}

	//response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}
