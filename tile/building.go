package building

type BuildingName int

const (
	_ BuildingName = iota
	_
	FarmersHouse
	PlowMakersWorkShop
	NovicesHut
	WorkShop
	_
	WeavingParlor
	ColonistsHouse
	CarpentersWorkshop
	SchnappsDistillery
	LoadingStation
	LitterStorage
	WoodTrader
	_
	Turnery
	Smokehouse
	Smithy
	Cooperage
	Bakehouse
	_
	Mill
	WeavingMill
	TextileHouse
	Saddlery
	Joinery
	WaterFrontHouse
	PottersInn
	FarmersInn
	JunkDealersInn
	GulfHouseInn
	MilkHouseInn
	SluiceYardInn
	_
	VillageChurch
	LutetsburgCastle
	BerumCastle
)

type Building struct {
	cost  [][]Good
	worth int
}

func GetBuildingVp(b BuildingName) int {
	switch b {
	case FarmersHouse, PlowMakersWorkShop, NovicesHut, WorkShop:
		return 1
	case WeavingParlor, ColonistsHouse, CarpentersWorkshop, SchnappsDistillery, LoadingStation, LitterStorage, WoodTrader:
		return 1
	case VillageChurch, LutetsburgCastle, BerumCastle:
		return 15
	default:
		panic("Unknown Building for VP")
	}

}
