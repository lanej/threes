package board

import (
	"fmt"
	"math/rand"
)

type Board struct {
	tiles [][]int
	deck  Deck
}

func newBoard(starter int) (*Board, error) {
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

	insertErr := board.insert(startCornerRow, startCornerCol, starter)

	// Of the 8 starter tiles, there are between 1-4 {1,2,3}
	return &board, insertErr
}

func (b *Board) print() {
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			fmt.Printf("|% 4d", b.tiles[row][col])
		}
		fmt.Println("|")
	}
}

type InvalidValueError struct {
	Value int
}

func (e InvalidValueError) Error() string {
	return fmt.Sprintf("Invalid value: %d", e.Value)
}

func (b *Board) insert(x, y, v int) error {
	valid := false

	for _, a := range Values {
		valid = valid || (a == v)
	}

	if !valid {
		return InvalidValueError{Value: v}
	}

	b.tiles[x][y] = v

	return nil
}
