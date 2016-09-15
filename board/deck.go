package board

import "fmt"

type Deck struct {
	tiles []int
}

var Values = []int{1, 2, 3, 6, 12, 24, 48, 96, 192, 384, 768, 1536, 3072, 6144, 12288}

func (d *Deck) print() {
	fmt.Println(d.tiles)
}
