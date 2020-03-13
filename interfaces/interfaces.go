package interfaces

type Size int

const (
	Small Size = iota
	Large
)

type Tile interface {
	Vp() int
	Size() Size
	AnimalRoom() int
	Flip() Tile
}
