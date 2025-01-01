package main

import (
	"log"
	"net/http"
	"bulletin-board/internal/handler"
  "bulletin-board/internal/repository"
  "bulletin-board/internal/usecase"
	"bulletin-board/pkg/config"
  "bulletin-board/pkg/db"
	"github.com/gorilla/mux"
)

func main() {
  database := db.Init()
  defer database.Close()

	cfg := config.LoadConfig()
	r := mux.NewRouter()

  boardRepo := repository.NewBoardRepository(database)
  boardUsecase := usecase.NewBoardUsecase(boardRepo)
  boardHandler := handler.NewBoardHandler(boardUsecase)

	r.HandleFunc("/boards", boardHandler.GetBoards).Methods("GET")

	log.Printf("Server started at :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}

