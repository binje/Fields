package main

import . "github.com/binje/Fields/interfaces"

var maxDikeMoves = 7

type HomeBoard struct {
	dikeMoves   int
	openSpaces  int
	grainFields int
	flaxFields  int
	forests     int
	smallMoor   moor
	largeMoors  []moor
	animals     map[Animal]int
	tiles       []Tile
}

//TODO new homeboard
func NewHomeBoard() *HomeBoard {
	return &HomeBoard{
		dikeMoves:  0,
		openSpaces: 2,
		animals:    make(map[Animal]int),
	}
}

type Animal int

const (
	Sheep Animal = iota
	Cow
	Horse
)

type moor struct {
	Dehydrated bool
	peat       int
}

func NewMoor() moor {
	return moor{false, 4}
}

func (m *moor) Flip() {
	m.Dehydrated = true
}

func (m *moor) cut() {
	m.peat--
}
func (m *moor) isClear() bool {
	return m.peat == 0
}

func (h *HomeBoard) FlipMoor() {
	for _, m := range h.largeMoors {
		if !m.Dehydrated {
			m.Dehydrated = true
			return
		}
	}
}

func (h *HomeBoard) CutPeat(i int) {
	//TODO give option of what to cut
	for ; i > 0 && !h.smallMoor.isClear(); i-- {
		h.smallMoor.cut()
	}
	for _, m := range h.largeMoors {
		if !m.Dehydrated {
			return
		}
		for ; i > 0 && !m.isClear(); i-- {
			m.cut()
		}
	}
}

func (h *HomeBoard) Breed(a Animal) {
	h.animals[a]++
}

func (h *HomeBoard) GetAnimal(a Animal) int {
	return h.animals[a]
}

func (h *HomeBoard) MoveDike() {
	h.dikeMoves++
	switch h.dikeMoves {
	case 1, 4, 7:
		h.openSpaces += 3
	}
}

func (h *HomeBoard) AddTile() {
	h.openSpaces--
	//TODO throw error?
}

func (h *HomeBoard) VP() int {
	return dikeVP(h.dikeMoves)
	switch h.dikeMoves {
	case 0, 1, 2, 3, 4:
		return -3
	case 5:
		return -2
	case 6:
		return -1
	default:
		return 0
	}
}

func dikeVP(i int) (vp int) {
	vp = i - 7
	if vp < -3 {
		return -3
	}
	return
}
