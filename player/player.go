package player

import "fmt"

type Player struct {
	tools map[Tool]int
}

func NewPlayer() *Player {
	return &Player{tools: make(map[Tool]int)}
}

type Tool int

const (
	NoTool Tool = iota
	FishTrap
	FleshingBeam
	WeavingLoom
	SlaughteringTable
	Spade
	ShovelPair
	Shovel
	PotteryWheel
	Oven
	Axe
	Workbench
)

var toolRamp = map[Tool][]int{
	FishTrap:          []int{2, 3, 4, 5, 6},
	FleshingBeam:      []int{3, 5, 6},
	WeavingLoom:       []int{2, 3, 4, 5},
	SlaughteringTable: []int{2, 3, 4},
	Spade:             []int{3, 5, 7},
	ShovelPair:        []int{1, 2, 2, 3},
	Shovel:            []int{3, 4, 5, 6},
	PotteryWheel:      []int{2, 3, 4},
	Oven:              []int{1, 2, 3, 4},
	Axe:               []int{3, 4, 5, 6},
	Workbench:         []int{2, 3, 4}}

func (p *Player) IncreaseTool(t Tool) {
	if p.tools[t]+1 == len(toolRamp[t]) {
		fmt.Printf("%s tool already maxed\n", t)
		return
	}
	p.tools[t]++
	fmt.Printf("%s tool increased to %d\n", t, p.tools[t])
}

func (p *Player) GetToolCount(t Tool) int {
	return toolRamp[t][p.tools[t]]
}
