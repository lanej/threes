package board

import "testing"

func TestBoardPrint(*testing.T) {
	board := newBoard(3)
	board.print()
}
