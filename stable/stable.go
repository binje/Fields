package stable

type Stable struct {
	vehicles          []Vehicle
	smallSpacesLeft   int
	largeSpacesLeft   int
	wildcardSpaceOpen bool
}

var numSmallSpaces = 4
var numLargeSpaces = 2

func NewStable() *Stable {
	return &Stable{
		vehicles:          make([]Vehicle, 0),
		smallSpacesLeft:   numSmallSpaces,
		largeSpacesLeft:   numLargeSpaces,
		wildcardSpaceOpen: true,
	}
}

func (s *Stable) getVp() int {
	vp := 0
	for _, v := range s.vehicles {
		vp += v.GetVp()
	}
	if s.largeSpacesLeft == numLargeSpaces {
		vp -= 3
	}
	return vp
}

func (s *Stable) CanAddVehicle(v Vehicle) bool {
	if s.wildcardSpaceOpen {
		return true
	}
	if v.GetSize() == Small {
		return s.smallSpacesLeft > 0
	} else {
		return s.largeSpacesLeft > 0
	}
}

func (s *Stable) AddVehicle(n VehicleName) {
	//TODO throw away equipment?
	v := *GetVehicle(n)
	if !s.CanAddVehicle(v) {
		return
	}
	s.vehicles = append(s.vehicles, v)
	if v.GetSize() == Small {
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

func (s *Stable) GetNumPlows() (sum int) {
	for _, v := range s.vehicles {
		if v.IsPlow() {
			sum++
		}
	}
	return
}
