package handler

import (
	"encoding/json"
	"net/http"
	"bulletin-board/internal/domain"
)

type UserHandler struct {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []domain.User{
		{ID: 1, Name: "John", Age: 30},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
