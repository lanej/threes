package board

import "testing"

func TestBoardPrint(t *testing.T) {
	board, err := newBoard(96)
	if err != nil {
		t.Error(err.Error())
	}
	board.print()
}
