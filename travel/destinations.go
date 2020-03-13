package Travel

type Destination int

const (
	Norden Destination = iota
	Hage
	Beemoor
	Dornum
	Esens
	Aurich
	Emden
	Leer
	Bremen
)

func (d Destination) HaulSize() int {
	switch d {
	case Hage, Beemoor, Dornum:
		return 1
	case Norden, Aurich, Esens:
		return 2
	case Emden, Leer:
		return 3
	case Bremen:
		return 4
	default:
		panic("Unknown destination")
	}
}

func TravelVp(destinations []Destination) int {
	distance := 0
	for _, d := range destinations {
		distance += d.HaulSize()
	}
	switch {
	case distance > 8:
		return distance/2 + 1
	case distance > 5:
		return 4
	case distance > 2:
		return 3
	case distance > 1:
		return 1
	default:
		return 0
	}
}
