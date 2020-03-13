package action

type Action int

const (
	_ Action = iota
	_
	Fisherman
	Grocer
	WoolWeaver
	Colonist
	PeatCutter
	DikeBuilder
	ClayWorker
	Farmer
	Forester
	Woodcutter
	Master
	Carpenter
	Builder
	Warden
	Laborer
	_
	PeatBoatman
	Tanner
	LinenWeaver
	Butcher
	CattleTrader
	Grocer2
	BuildersMerchant
	Potter
	Baker
	WoodTrader
	Master2
	Wainwright
	DikeWarder
	Carpenter2
	Laborer2
	_
	Finish
	_
	GetSheep
	GetCow
	GetHorse
	GetWood
	Get4Wood
	GetTimber
	GetClay
	GetBrick
	GetFlaxField
	GetGrainField
	GetForest
	_
	UseSheep
	UseCow
	UseHorse
	UseWood
	UseTimber
	UseClay
	UseBrick
	UseFlaxField
	UseGrainField
	UseForest
	_
	BuildVehicle
	BuildHorseCart
	BuildCarriage
	BuildDroshsky
	BuildHandcart
	BuildWagon
	BuildPeatBoat
	BuildPlow
)

/*
type Choice []Action

func (a Action) GetNextActions() []Choices {
	c := make([]Choice, 0)
	switch a {
	case Grocer:
		return
	default:
		return c
	}
}
*/

/*

var ToolActions = map[Action]Tool{
	WoolWeaver:  WeavingLooms,
	Master:      Workbenches,
	Tanner:      FleshingBeams,
	LinenWeaver: WeavingLooms,
	Butcher:     SlaughteringTables,
	Potter:      PotteryWheels,
	Baker:       Ovens,
	Master2:     Workbenches,
}
*/
