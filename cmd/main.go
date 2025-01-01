package main

import (
	"log"
	"net/http"
	"bulletin-board/internal/handler"
	"bulletin-board/pkg/config"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	r := mux.NewRouter()
	h := handler.NewUserHandler()

	r.HandleFunc("/users", h.GetUsers).Methods("GET")

	log.Printf("Server started at :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
