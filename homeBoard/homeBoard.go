package home

import (
	. "github.com/binje/Fields/actions"
)

var maxDikeMoves = 7

type HomeBoard struct {
	dikeMoves  int
	openSpaces int
	forests    int
	smallMoor  moor
	largeMoors []moor
	animals    map[animal]int
	landscapes map[landscape]int
}

func NewHomeBoard() *HomeBoard {
	return &HomeBoard{
		dikeMoves:  0,
		openSpaces: 2,
		animals:    make(map[animal]int),
		landscapes: make(map[landscape]int),
	}
}

func (h *HomeBoard) DoAction(action Action, i int) {
	switch action {
	// Dike Actions
	case DikeWarden:
		h.moveDike(1)
	case DikeBuilder:
		//TODO get shovelpairs
		h.moveDike(i)
	// landscape actions
	case PlowGrainField:
		h.addLandscape(GrainField)
	case PlowFlaxField:
		h.addLandscape(FlaxField)
	case PlantForest:
		h.addLandscape(forest)
	case BuildStall:
		h.addLandscape(stall)
	case BuildStable:
		h.landscapes[stall]--
		h.landscapes[stable]++
	//Vehicles
	case BuildCarriage, BuildDroshsky:
		h.useAnimal(horse)
		h.useAnimal(horse)
	// Other
	case Fisherman, CattleTrader:
		h.animals[sheep]++
	case Colonist:
		h.flipMoor()
		h.animals[horse]++
	case Builder:
		h.openSpaces--
	// Travel
	// TODO handle trade moor
	case TravelHage:
		h.openSpaces++
	case SellGrainField:
		h.landscapes[GrainField]--
	case SellFlaxField:
		h.landscapes[FlaxField]--
	case ButcherSheep, SlaughterSheep, TradeSheep:
		h.useAnimal(sheep)
	case ButcherCow, SlaughterCow, TradeCow:
		h.useAnimal(cow)
	case ButcherHorse, SlaughterHorse, TradeHorse, BuildCart, BuildHorseCart:
		h.useAnimal(horse)
	}
}

func (h *HomeBoard) useAnimal(a animal) {
	h.animals[a]++
	//fmt.Printf("You now have %d %s\n", h.animals[a], a)
}

// Dikes
func (h *HomeBoard) moveDike(i int) {
	h.dikeMoves++
	switch h.dikeMoves {
	case 1, 4, 7:
		h.openSpaces += 3
	}
}

// Landscapes
type landscape int

const (
	forest landscape = iota
	park
	stall
	depot
	stable
	doubleStall
	Grainfield
	FlaxField
	GrainField
)

func (h *HomeBoard) addLandscape(l landscape) {
	h.landscapes[l]++
	h.openSpaces--
}

func (h *HomeBoard) flip(l landscape) {
	h.landscapes[l]--
	h.landscapes[flipTable[l]]++
}

var flipTable = map[landscape]landscape{
	forest:      park,
	park:        forest,
	stall:       depot,
	depot:       stall,
	stable:      doubleStall,
	doubleStall: stable,
	GrainField:  FlaxField,
	FlaxField:   GrainField,
}

// animals
type animal int

const (
	sheep animal = iota
	cow
	horse
)

func (h *HomeBoard) NumAnimals() int {
	return h.animals[sheep] + h.animals[cow] + h.animals[horse]
}

func (h *HomeBoard) KillAllAnimals() {
	h.animals[sheep] = 0
	h.animals[cow] = 0
	h.animals[horse] = 0
}

func (h *HomeBoard) GetSlaughterOptions() []Action {
	actions := make([]Action, 3)
	if h.animals[sheep] > 0 {
		actions = append(actions, SlaughterSheep)
	}
	if h.animals[cow] > 0 {
		actions = append(actions, SlaughterCow)
	}
	if h.animals[horse] > 0 {
		actions = append(actions, SlaughterHorse)
	}
	return actions
}

func (h *HomeBoard) Breed(a animal) {
	h.animals[a]++
}

func (h *HomeBoard) GetBreedingNumbers() (independent, repetitive int) {
	for landscape, qty := range h.landscapes {
		switch landscape {
		case stall:
			independent += qty
		case doubleStall:
			independent += 2 * qty
		case stable:
			repetitive += qty
		}
	}
	return
}

// Space

func (h *HomeBoard) IsOverPopulated() bool {
	sheeps := h.animals[sheep]
	cows := h.animals[cow]
	horses := h.animals[horse]
	for i := 0; i < h.landscapes[stable]; i++ {
		sheeps, cows, horses = removeFromLargest(6, sheeps, cows, horses)
	}
	for i := 0; i < 2*h.landscapes[doubleStall]; i++ {
		sheeps, cows, horses = removeFromLargest(3, sheeps, cows, horses)
	}
	for i := 0; i < h.landscapes[stall]; i++ {
		sheeps, cows, horses = removeFromLargest(3, sheeps, cows, horses)
	}
	for i := 0; i < 2*h.landscapes[park]; i++ {
		sheeps, cows, horses = removeFromLargest(1, sheeps, cows, horses)
	}
	return sheeps < 0 || cows < 0 || horses < 0
}

func removeFromLargest(qty, a, b, c int) (int, int, int) {
	if a >= b && a >= c {
		a -= qty
	} else if b >= a && b >= c {
		b -= qty
	} else {
		c -= qty
	}
	return a, b, c
}

// Moor

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

func (h *HomeBoard) flipMoor() {
	for _, m := range h.largeMoors {
		if !m.Dehydrated {
			m.Dehydrated = true
			return
		}
	}
}

func (h *HomeBoard) CutPeat(i int) (cut int) {
	//TODO give option of what to cut
	for ; i > 0 && !h.smallMoor.isClear(); i-- {
		h.smallMoor.cut()
		cut++
	}
	for _, m := range h.largeMoors {
		if !m.Dehydrated {
			return
		}
		for ; i > 0 && !m.isClear(); i-- {
			m.cut()
			cut++
		}
	}
	return
}

func (h *HomeBoard) NovemberInventory() (food, grain, flax, wood int) {
	sheep := h.animals[sheep]
	switch {
	case sheep >= 8:
		food = 3
	case sheep >= 5:
		food = 2
	case sheep >= 2:
		food = 1
	}

	cow := h.animals[cow]
	switch {
	case cow >= 5:
		food += 3
	case cow >= 3:
		food += 2
	case cow >= 1:
		food += 1
	}

	return food,
		h.landscapes[GrainField],
		h.landscapes[FlaxField],
		h.landscapes[forest]
}

func (h *HomeBoard) MayInventory() (wool int) {
	//TODO get wool after breeding
	sheep := h.animals[sheep]
	switch {
	case sheep >= 6:
		return 3
	case sheep >= 4:
		return 2
	case sheep >= 1:
		return 1
	default:
		return
	}
	//TODO breed animals
}

func (h *HomeBoard) Vp() int {
	return h.landscapeVp() + dikeVp(h.dikeMoves)
}

func (h *HomeBoard) landscapeVp() (vp int) {
	for landscape, qty := range h.landscapes {
		vp += qty * landscapePoints[landscape]
	}
	return
}

var landscapePoints = map[landscape]int{
	forest:      2,
	park:        6,
	stall:       2,
	depot:       2,
	stable:      6,
	doubleStall: 9,
}

func dikeVp(dikeMoves int) int {
	switch dikeMoves {
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
