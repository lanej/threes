package board

import (
	"fmt"
	"math"
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
	remainingSeed := 8

	seeds := []int{3, 2, 1}

	for _, seed := range seeds {
		maxSeed := math.Min(float64(3), float64(remainingSeed))
		numSeed := rand.Intn(int(maxSeed)) + 1

		for i := 0; i < numSeed; i++ {
			empties := board.EmptyTiles()
			index := rand.Intn(len(empties))

			position := empties[index]

			board.insert(position[0], position[1], seed)
		}

		remainingSeed -= numSeed
	}

	return &board, insertErr
}

func (b *Board) print() {
	for x, _ := range b.tiles {
		for y, _ := range b.tiles[x] {
			fmt.Printf("|% 4d", b.tiles[x][y])
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

func (b *Board) EmptyTiles() [][]int {
	empties := make([][]int, 0)

	for x, row := range b.tiles {
		for y, v := range row {
			if v == 0 {
				empties = append(empties, []int{x, y})
			}
		}
	}

	return empties
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
