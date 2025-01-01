package handler

import (
  "bulletin-board/internal/usecase"
  "encoding/json"
  "net/http"
)

type BoardHandler struct {
  usecase *usecase.BoardUsecase
}

func NewBoardHandler(u *usecase.BoardUsecase) *BoardHandler {
  return &BoardHandler{usecase: u}
}

func (h *BoardHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {
  name := r.FormValue("name")
  if err := h.usecase.CreateBoard(name); err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(map[string]string{
      "error": err.Error(),
    })
    return
  }

  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(map[string]string{
    "message": "Board created",
  })

}

func (h *BoardHandler) GetBoards(w http.ResponseWriter, r *http.Request) {
  boards, err := h.usecase.GetBoards()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(map[string]string{
      "error": err.Error(),
    })
    return
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(boards)

}

