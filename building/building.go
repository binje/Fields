package building

type Building int

const (
	FarmersHouse Building = iota
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

/*:GoB
func (b *Building) Cost() [][]Action {
	switch b {
	case FarmersHouse, PlowMakersWorkShop, NovicesHut, WorkShop, WeavingParlor, ColonistsHouse, CarpentersWorkshop, SchnappsDistillery, LoadingStation, LitterStorage, WoodTrader:
		return [][]Action{
			[]Action{UseWood, UseClay, UseTimber, UseBrick},
			[]Action{UseGrain}}
	case Turnery, Smokehouse, Smithy, Cooperage, Bakehouse:
		return [][]Action{
			[]Action{UseTimber},
			[]Action{UseBrick}}
	}
	return [][]Action{}
}
*/

func GetBuildingVp(b Building) int {
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
