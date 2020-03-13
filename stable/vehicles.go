package stable

type VehicleName int

const (
	Cart VehicleName = iota
	HorseCart
	Carriage
	Droshsky
	Handcart
	Wagon
	PeatBoat
	Plow
)

type Size bool

const (
	Small = true
	Large = false
)

type Vehicle struct {
	name VehicleName
	size Size
	vp   int
}

func (v *Vehicle) Flip() *Vehicle {
	switch v.name {
	case Cart:
		return GetHorseCart()
	case HorseCart:
		return GetCart()
	case Carriage:
		return GetDroshsky()
	case Droshsky:
		return GetCarriage()
	case Handcart:
		return GetWagon()
	case Wagon:
		return GetHandcart()
	case PeatBoat:
		return GetPlow()
	case Plow:
		return GetPeatBoat()
	default:
		panic("Got unknown vehicle")
	}
}

func GetVehicle(n VehicleName) *Vehicle {
	switch n {
	case Cart:
		return GetCart()
	case HorseCart:
		return GetHorseCart()
	case Carriage:
		return GetCarriage()
	case Droshsky:
		return GetDroshsky()
	case Handcart:
		return GetHandcart()
	case Wagon:
		return GetWagon()
	case PeatBoat:
		return GetPeatBoat()
	case Plow:
		return GetPlow()
	default:
		panic("Got unknown vehicle")
	}
}

func (v *Vehicle) GetSize() Size {
	return v.size
}

func (v *Vehicle) IsPlow() Size {
	return v.name == Plow
}

func (v *Vehicle) GetVp() int {
	return v.vp
}

func (v Vehicle) GetAvailableHauls() [][]int {
	switch v.name {
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

func GetCart() *Vehicle {
	return &Vehicle{
		name: Cart,
		size: Large,
		vp:   1,
	}
}

func GetHorseCart() *Vehicle {
	return &Vehicle{
		name: HorseCart,
		size: Large,
		vp:   2,
	}
}

func GetCarriage() *Vehicle {
	return &Vehicle{
		name: Carriage,
		size: Large,
		vp:   4,
	}
}

func GetDroshsky() *Vehicle {
	return &Vehicle{
		name: Droshsky,
		size: Large,
		vp:   5,
	}
}

func GetHandcart() *Vehicle {
	return &Vehicle{
		name: Handcart,
		size: Small,
		vp:   0,
	}
}
func GetWagon() *Vehicle {
	return &Vehicle{
		name: Wagon,
		size: Small,
		vp:   1,
	}
}

func GetPeatBoat() *Vehicle {
	return &Vehicle{
		name: PeatBoat,
		size: Small,
		vp:   1,
	}
}

func GetPlow() *Vehicle {
	return &Vehicle{
		name: Plow,
		size: Small,
		vp:   3,
	}
}
