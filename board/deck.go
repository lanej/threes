package board

type Deck struct {
	tiles []int
}

func (d *Deck) print() {
	fmt.Println(d.tiles)
}
