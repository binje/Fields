package stable

import . "github.com/binje/Fields/actions"

type Stable struct {
	vehicles          map[Vehicle]int
	smallSpacesLeft   int
	largeSpacesLeft   int
	wildcardSpaceOpen bool
}

var numSmallSpaces = 4
var numLargeSpaces = 2

func NewStable() *Stable {
	return &Stable{
		vehicles:          make(map[Vehicle]int),
		smallSpacesLeft:   numSmallSpaces,
		largeSpacesLeft:   numLargeSpaces,
		wildcardSpaceOpen: true,
	}
}

func (s *Stable) DoAction(action Action) {
	if vehicle, ok := buildVehicle[action]; ok {
		s.addVehicle(vehicle)
		return
	}
	switch action {
	case TravelDornum:
		s.removeSmallVehicle(Plow)
	case TradePeatBoat:
		s.removeSmallVehicle(PeatBoat)
	}
}

var buildVehicle = map[Action]Vehicle{
	BuildHorseCart: HorseCart,
	BuildCarriage:  Carriage,
	BuildDroshsky:  Droshsky,
	BuildHandcart:  Handcart,
	BuildWagon:     Wagon,
	BuildCart:      Cart,
	BuildPeatBoat:  PeatBoat,
	BuildPlow:      Plow,
}

func (s *Stable) canAddVehicle(v Vehicle) bool {
	if s.wildcardSpaceOpen {
		return true
	}
	if v.IsSmall() {
		return s.smallSpacesLeft > 0
	} else {
		return s.largeSpacesLeft > 0
	}
}

func (s *Stable) addVehicle(v Vehicle) {
	s.vehicles[v]++
	if v.IsSmall() {
		if s.smallSpacesLeft > 0 {
			s.smallSpacesLeft--
		} else {
			s.wildcardSpaceOpen = false
		}
	} else {
		if s.largeSpacesLeft > 0 {
			s.largeSpacesLeft--
		} else {
			s.wildcardSpaceOpen = false
		}
	}
}

func (s *Stable) removeSmallVehicle(v Vehicle) {
	if s.vehicles[Cart]+s.vehicles[Wagon]+s.vehicles[PeatBoat]+s.vehicles[Plow] > numSmallSpaces {
		s.wildcardSpaceOpen = true
	} else {
		s.smallSpacesLeft++
	}
	s.vehicles[v]--
}

func (s *Stable) removeLargeVehicle(v Vehicle) {
	if s.vehicles[Handcart]+s.vehicles[Droshsky]+s.vehicles[Carriage]+s.vehicles[HorseCart] > numLargeSpaces {
		s.wildcardSpaceOpen = true
	} else {
		s.smallSpacesLeft++
	}
	s.vehicles[v]--
}

func (s *Stable) NumPlows() int {
	return s.vehicles[Plow]
}

func (s *Stable) NumPeatBoats() int {
	return s.vehicles[PeatBoat]
}

func (s *Stable) Vp() (vp int) {
	for v, qty := range s.vehicles {
		vp += v.GetVp() * qty
	}
	if s.largeSpacesLeft == numLargeSpaces {
		vp -= 3
	}
	return
}
