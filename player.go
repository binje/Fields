package main

import "fmt"

type Tool int

const (
	FishTraps Tool = iota
	FleshingBeams
	WeavingLooms
	SlaughteringTables
	Spades
	ShovelPairs
	Shovels
	PotteryWheels
	Ovens
	Axes
	Workbenches
)

var toolRamp = map[Tool][]int{
	FishTraps:          []int{2, 3, 4, 5, 6},
	FleshingBeams:      []int{3, 5, 6},
	WeavingLooms:       []int{2, 3, 4, 5},
	SlaughteringTables: []int{2, 3, 4},
	Spades:             []int{3, 5, 7},
	ShovelPairs:        []int{1, 2, 2, 3},
	Shovels:            []int{3, 4, 5, 6},
	PotteryWheels:      []int{2, 3, 4},
	Ovens:              []int{1, 2, 3, 4},
	Axes:               []int{3, 4, 5, 6},
	Workbenches:        []int{2, 3, 4}}

type Player struct {
	m map[Tool]int
}

func NewPlayer() *Player {
	return &Player{make(map[Tool]int)}
}

func (p *Player) IncreaseTool(tool Tool) {
	if p.m[tool]+1 == len(toolRamp[tool]) {
		fmt.Printf("%s tool already maxed\n", tool)
		return
	}
	p.m[tool]++
	fmt.Printf("%s tool increased to %d\n", tool, p.m[tool])
}

func (p *Player) GetToolCount(tool Tool) int {
	return toolRamp[tool][p.m[tool]]
}

/*
func (p *Player) IncreaseFishTraps() {
	p.increaseTool(FishTraps)
}

func (p *Player) getFishTraps() int {
	return p.getTool(FishTraps)
}

func (p *Player) IncreaseFleshingBeams() {
	p.increaseTool(FleshingBeams)
}

func (p *Player) getFleshingBeams() int {
	return p.getTool(FleshingBeams)
}

func (p *Player) IncreaseWeavingLooms() {
	p.increaseTool(WeavingLooms)
}

func (p *Player) getWeavingLooms() int {
	return p.getTool(WeavingLooms)
}

func (p *Player) IncreaseSlaughteringTables() {
	p.increaseTool(SlaughteringTables)
}

func (p *Player) getSlaughteringTables() int {
	return p.getTool(SlaughteringTables)
}

func (p *Player) IncreaseSpades() {
	p.increaseTool(Spades)
}

func (p *Player) getSpades() int {
	return p.getTool(Spades)
}

func (p *Player) IncreaseShovelPairs() {
	p.increaseTool(ShovelPairs)
}

func (p *Player) getShovelPairs() int {
	return p.getTool(ShovelPairs)
}

func (p *Player) IncreaseShovels() {
	p.increaseTool(Shovels)
}

func (p *Player) getShovels() int {
	return p.getTool(Shovels)
}

func (p *Player) IncreasePotteryWheels() {
	p.increaseTool(PotteryWheels)
}

func (p *Player) getPotteryWheels() int {
	return p.getTool(PotteryWheels)
}

func (p *Player) IncreaseOvens() {
	p.increaseTool(Ovens)
}

func (p *Player) getOvens() int {
	return p.getTool(Ovens)
}

func (p *Player) IncreaseAxes() {
	p.increaseTool(Axes)
}

func (p *Player) getAxes() int {
	return p.getTool(Axes)
}

func (p *Player) IncreaseWorkBenches() {
	p.increaseTool(WorkBenches)
}

func (p *Player) getWorkBenches() int {
	return p.getTool(WorkBenches)
}
*/
