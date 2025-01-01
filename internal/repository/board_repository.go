package repository

import (
    "bulletin-board/internal/domain"
    "database/sql"
)

type BoardRepository struct {
    db *sql.DB
}

func NewBoardRepository(db *sql.DB) *BoardRepository {
    return &BoardRepository{db: db}
}

func (r *BoardRepository) Create(board *domain.Board) error {
    _, err := r.db.Exec("INSERT INTO boards (name) VALUES ($1)", board.Name)
    return err
}

func (r *BoardRepository) GetAll() ([]domain.Board, error) {
    rows, err := r.db.Query("SELECT id, name FROM boards")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var boards []domain.Board
    for rows.Next() {
        var board domain.Board
        err := rows.Scan(&board.ID, &board.Name)
        if err != nil {
            return nil, err
        }
        boards = append(boards, board)
    }
    return boards, nil
}
