package time

import . "github.com/binje/Fields/actions"

//TODO handle crosscalls
// Laborer2 -> builder
// DikeWarden -> warden
func OffSeason(m Month, a Action) bool {
	if _, ok := SummerEmployment[a]; ok {
		if m.Season() == DecemberPreperations {
			return true
		}
	}
	if _, ok := WinterEmployment[a]; ok {
		if m.Season() == JunePreperations {
			return true
		}
	}
	return false
}

var SummerEmployment = map[Action]struct{}{
	Fisherman:   struct{}{},
	Grocer:      struct{}{},
	WoolWeaver:  struct{}{},
	Colonist:    struct{}{},
	PeatCutter:  struct{}{},
	DikeBuilder: struct{}{},
	ClayWorker:  struct{}{},
	Farmer:      struct{}{},
	Forester:    struct{}{},
	Woodcutter:  struct{}{},
	Master:      struct{}{},
	Carpenter:   struct{}{},
	Builder:     struct{}{},
	Warden:      struct{}{},
	Laborer:     struct{}{},
}

var SummerEmploymentArray = []Action{
	Fisherman,
	Grocer,
	WoolWeaver,
	Colonist,
	PeatCutter,
	DikeBuilder,
	ClayWorker,
	Farmer,
	Forester,
	Woodcutter,
	Master,
	Carpenter,
	Builder,
	Warden,
	Laborer,
}

var WinterEmployment = map[Action]struct{}{
	PeatBoatman:      struct{}{},
	Tanner:           struct{}{},
	LinenWeaver:      struct{}{},
	Butcher:          struct{}{},
	CattleTrader:     struct{}{},
	Grocer2:          struct{}{},
	BuildersMerchant: struct{}{},
	Potter:           struct{}{},
	Baker:            struct{}{},
	WoodTrader:       struct{}{},
	Master2:          struct{}{},
	Wainwright:       struct{}{},
	DikeWarder:       struct{}{},
	Carpenter2:       struct{}{},
	Laborer2:         struct{}{},
}

var WinterEmploymentArray = []Action{
	PeatBoatman,
	Tanner,
	LinenWeaver,
	Butcher,
	CattleTrader,
	Grocer2,
	BuildersMerchant,
	Potter,
	Baker,
	WoodTrader,
	Master2,
	Wainwright,
	DikeWarder,
	Carpenter2,
	Laborer2,
}

var AllEmploymentArray = append(SummerEmploymentArray, WinterEmploymentArray...)

var AllEmployment = map[Action]struct{}{
	Fisherman:        struct{}{},
	Grocer:           struct{}{},
	WoolWeaver:       struct{}{},
	Colonist:         struct{}{},
	PeatCutter:       struct{}{},
	DikeBuilder:      struct{}{},
	ClayWorker:       struct{}{},
	Farmer:           struct{}{},
	Forester:         struct{}{},
	Woodcutter:       struct{}{},
	Master:           struct{}{},
	Carpenter:        struct{}{},
	Builder:          struct{}{},
	Warden:           struct{}{},
	Laborer:          struct{}{},
	PeatBoatman:      struct{}{},
	Tanner:           struct{}{},
	LinenWeaver:      struct{}{},
	Butcher:          struct{}{},
	CattleTrader:     struct{}{},
	Grocer2:          struct{}{},
	BuildersMerchant: struct{}{},
	Potter:           struct{}{},
	Baker:            struct{}{},
	WoodTrader:       struct{}{},
	Master2:          struct{}{},
	Wainwright:       struct{}{},
	DikeWarder:       struct{}{},
	Carpenter2:       struct{}{},
	Laborer2:         struct{}{},
}
