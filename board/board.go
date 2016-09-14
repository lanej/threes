package board

import (
	"fmt"
	"math/rand"
)

type Board struct {
	tiles [][]int
	deck  Deck
}

func newBoard(starter int) *Board {
	// Initialize Board
	board := Board{
		tiles: [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		deck: Deck{},
	}

	// Starter value goes in a random corner, seemingly rotates
	startCornerRow := rand.Intn(1) * 3
	startCornerCol := rand.Intn(1) * 3

	board.tiles[startCornerRow][startCornerCol] = starter

	// Of the 8 starter tiles, there are between 1-4 {1,2,3}
	return new(Board)
}

func (b *Board) print() {
	fmt.Println("something")
}
