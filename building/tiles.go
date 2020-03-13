package building

import (
	. "github.com/binje/Fields/interfaces"
)

type Field int

const (
	Forest Field = iota
	Park
	Stall
	Depot
	Stable
	DoubleStall
	GrainField
	FlaxField
)

func (f Field) Size() Size {
	return Small
}

func (f Field) AnimalRoom() {
	//TODO abide by only same animal
	switch f {
	case Park:
		return 2
	case Stall:
		return 3
	case Stable, DoubleStall:
		return 6
	default:
		return 0
	}
}

func (f Field) Vp() int {
	switch f {
	case GrainField, FlaxField:
		return 0
	case Forest, Stall, Depot:
		return 2
	case Park, Stable:
		return 6
	case DoubleStall:
		return 9
	default:
		panic("Unknown tile for vp")
	}
}
