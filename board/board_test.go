package board

import "testing"

func TestBoardPrint(t *testing.T) {
	board, err := newBoard(96)
	if err != nil {
		t.Error(err.Error())
	}
	board.print()
}

func TestBoardEmptyTiles(t *testing.T) {
	board, _ := newBoard(96)
	empties := cap(board.EmptyTiles())

	if empties != 8 {
		t.Errorf("Expected 8 empty tiles, got %d", empties)
	}
}
