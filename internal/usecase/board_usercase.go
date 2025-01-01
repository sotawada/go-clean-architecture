package usecase

import "bulletin-board/internal/domain"

type BoardRepository interface {
    Create(board *domain.Board) error
    GetAll() ([]domain.Board, error)
}

type BoardUsecase struct {
    repo BoardRepository
}

func NewBoardUsecase(r BoardRepository) *BoardUsecase {
    return &BoardUsecase{repo: r}
}

func (u *BoardUsecase) CreateBoard(name string) error {
    board := &domain.Board{Name: name}
    return u.repo.Create(board)
}

func (u *BoardUsecase) GetBoards() ([]domain.Board, error) {
    return u.repo.GetAll()
}
