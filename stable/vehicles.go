package stable

import "fmt"

type Vehicle int

const (
	undefinedVehicle Vehicle = iota
	Cart
	HorseCart
	Carriage
	Droshsky
	Handcart
	Wagon
	PeatBoat
	Plow
	_ // invisible vehicles for haulsize
	noSpace
	twoSpace
	twoSpaceOnly
)

func (v *Vehicle) Flip() *Vehicle {
	if vehicle, ok := flipTable[*v]; ok {
		v = &vehicle
		return v
	}
	panic("Got unknown vehicle")
}

var flipTable = map[Vehicle]Vehicle{
	Cart:      HorseCart,
	HorseCart: Cart,
	Carriage:  Droshsky,
	Droshsky:  Carriage,
	Handcart:  Wagon,
	Wagon:     Handcart,
	PeatBoat:  Plow,
	Plow:      PeatBoat,
}

func (v Vehicle) GetAvailableHauls() [][]int {
	switch v {
	case Cart:
		return [][]int{
			[]int{1, 1, 1},
			[]int{2, 1},
			[]int{3},
		}
	case HorseCart:
		return [][]int{
			[]int{1, 1, 1, 1},
			[]int{2, 1, 1},
			[]int{2, 2},
			[]int{3, 1},
			[]int{4},
		}
	case Carriage:
		return [][]int{
			[]int{2, 1},
			[]int{3},
		}
	case Droshsky:
		return [][]int{
			[]int{2, 1, 1},
			[]int{3, 1},
			[]int{4},
		}
	case Handcart:
		return [][]int{[]int{1}}
	case Wagon:
		return [][]int{[]int{1, 1}}
	case PeatBoat, Plow:
		return [][]int{}
	default:
		panic("Got unknown vehicle")
	}
	return [][]int{}
}

//var TwoOnly =[][]int{2}
//var TwoPace =[][]int{2},[]int{1,1}

var haul4 = map[Vehicle]Vehicle{
	Droshsky:  noSpace,
	HorseCart: noSpace,
}

var haul3 = map[Vehicle]Vehicle{
	Cart:      noSpace,
	HorseCart: Handcart,
	Carriage:  noSpace,
	Droshsky:  Handcart,
}

var haul2 = map[Vehicle]Vehicle{
	Cart:      Handcart,
	HorseCart: twoSpace,
	Carriage:  Handcart,
	Droshsky:  Wagon,
}

var haul1 = map[Vehicle]Vehicle{
	Cart:      twoSpace,
	HorseCart: Cart,
	Carriage:  twoSpaceOnly,
	Droshsky:  Carriage,
	Handcart:  noSpace,
	Wagon:     Handcart,
}

func (v Vehicle) GetVp() int {
	switch v {
	case Handcart:
		return 0
	case Cart, Wagon, PeatBoat:
		return 1
	case HorseCart:
		return 2
	case Plow:
		return 3
	case Carriage:
		return 4
	case Droshsky:
		return 5
	default:
		fmt.Println(v)
		panic("Getting Vp for unknown vehicle")
	}
}

func (v Vehicle) IsSmall() bool {
	return v > Droshsky
}
