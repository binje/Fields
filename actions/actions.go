package action

import . "github.com/binje/Fields/player"

type Action int

const (
	_ Action = iota
	_
	GainEmployment
	_ // Summer Employment
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
	_ //WinterEmployment
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
	DikeWarden
	Carpenter2
	Laborer2
	_ //Tools
	WeaveWool
	TanHide
	WeaveLinen
	ButcherCow
	ButcherSheep
	ButcherHorse
	MakePot
	Bake
	_
	Finish
	Flip
	Next
	_
	GetFishTraps
	GetFleshingBeams
	GetWeavingLooms
	GetSlaughteringTables
	GetSpades
	GetShovelPaires
	GetPotteryWheels
	GetOvens
	GetAxes
	GetWorkBenches
	_
	GetSheep
	GetCow
	GetHorse
	GetWood
	Get4Wood
	GetTimber
	GetClay
	GetBrick
	_
	SlaughterSheep
	SlaughterCow
	SlaughterHorse
	_
	GetPlow
	PlowFlaxField
	PlowGrainField
	PlantForest
	BuildStall
	BuildStable
	_
	BuildVehicle
	BuildHorseCart
	BuildCarriage
	BuildDroshsky
	BuildHandcart
	BuildWagon
	BuildCart
	BuildPeatBoat
	BuildPlow
	_ //TODO use peatboat
	_ // Traveling
	TravelNorden
	TravelHage
	TravelBeemoor
	TravelDornum //TODO remove plow from stable
	TravelEsens
	TravelAurich
	TravelEmden
	TravelLeer
	TravelBremen
	_
	SellGrainField
	SellFlaxField
	TradeSheep
	TradePeatBoat
	TradeWinterWear
	TradeLeather
	TradeAnimal
	TradeHorse
	TradeHide
	Trade2Grain
	TradeWoolen
	TradePeat
	TradeSummerWear
	TradeLeatherWear
	TradeFlax
	TradeLinen
	TradeCow
	TradeClothing
	TradeMoor
	TradeTimber
	Trade2Animal
	TradeLinenSet
	TradeClothingSet
	UseLeatherWear
	UseSummerWear
	UseWinterWear
	_
	UsePeatBoat
	PeatForWool
	PeatForGrain
	PeatForFood
	PeatForFlax
	PeatForHide
	_ //Building
	BuildFarmersHouse
	BuildPlowMakersWorkShop
	BuildNovicesHut
	BuildWorkShop
	BuildWeavingParlor
	BuildColonistsHouse
	BuildCarpentersWorkshop
	BuildSchnappsDistillery
	BuildLoadingStation
	BuildLitterStorage
	BuildWoodTrader
	BuildTurnery
	BuildSmokehouse
	BuildSmithy
	BuildCooperage
	BuildBakehouse
	BuildMill
	BuildWeavingMill
	BuildTextileHouse
	BuildSaddlery
	BuildJoinery
	BuildWaterfrontHouse
	BuildPottersInn
	BuildFarmersInn
	BuildJunkDealersInn
	BuildGulfHouseInn
	BuildMilkHouseInn
	BuildSluiceYardInn
	BuildVillageChurch
	BuildLutetsburgCastle
	BuildBerumCastle
	_
	UseWood
	UseClay
	UseTimber
	UseBrick
	Use8Flax
	Use8Grain
	Use10Flax
	Use10Wool
)

type Choices [][]Action

func choose(actions ...Action) Choices {
	c := make([]Action, len(actions))
	for i, a := range actions {
		c[i] = a
	}
	return Choices{c}
}

func (c Choices) and(actions ...Action) Choices {
	return append([][]Action(c), actions)
}

func (c *Choices) UseTool(i int) {
	//TODO
	//get last choice extend it i times
}

var nextActionsMap = map[Action]Choices{
	// Summer Employment
	Grocer:      choose(GetSheep, GetCow, GetHorse, GetTimber, GetBrick),
	WoolWeaver:  choose(WeaveWool, Finish),
	Colonist:    choose(Flip).and(GetHorse),
	DikeBuilder: choose(GetCow, GetSheep),
	Farmer:      choose(GetPlow, Next).and(PlowGrainField, PlowFlaxField, Finish),
	Forester:    choose(PlantForest, Builder),
	// TODO cannot incease the same thing but an imitator of the master could?
	Master:    choose(buildAnyTool()...),
	Carpenter: choose(Builder, BuildStall),
	//TODO	Builder: choose(AllBuildings()...),
	Warden:  choose(Flip),
	Laborer: choose(BuildVehicle, GainEmployment),
	// Winter Employment
	Tanner:           choose(TanHide, Finish),
	LinenWeaver:      choose(WeaveLinen, Finish),
	CattleTrader:     choose(GetCow, GetHorse),
	Grocer2:          choose(getAnyAnimal()...),
	BuildersMerchant: choose(GetWood, GetClay).and(GetTimber, GetBrick),
	Potter:           choose(MakePot, Finish),
	Baker:            choose(Bake, Finish),
	WoodTrader:       choose(Get4Wood, Builder),
	Master2:          choose(buildAnyTool()...),
	Wainwright:       choose(BuildVehicle, Finish).and(BuildPeatBoat),
	DikeWarden:       choose(Flip),
	Carpenter2:       choose(BuildStable, BuildStall),
	Laborer2:         choose(BuildVehicle, GainEmployment),
	UsePeatBoat:      choose(PeatForWool, PeatForGrain, PeatForFood, PeatForFlax, PeatForHide, Finish),
	// Building
	BuildFarmersHouse:       choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildPlowMakersWorkShop: choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildNovicesHut:         choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildWorkShop:           choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildWeavingParlor:      choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildColonistsHouse:     choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildCarpentersWorkshop: choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildSchnappsDistillery: choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildLoadingStation:     choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildLitterStorage:      choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildWoodTrader:         choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildMill:               choose(Use8Flax, Use8Grain),
	BuildWeavingMill:        choose(Use10Flax, Use10Wool),
	BuildTextileHouse:       choose(UseWinterWear, UseWinterWear, UseSummerWear).and(UseWinterWear, UseWinterWear, UseSummerWear),
	BuildPottersInn:         choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildFarmersInn:         choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildJunkDealersInn:     choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildGulfHouseInn:       choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildMilkHouseInn:       choose(UseWood, UseClay, UseTimber, UseBrick),
	BuildSluiceYardInn:      choose(UseWood, UseClay, UseTimber, UseBrick),
	// Travel
	TravelHage: choose(SellGrainField, SellFlaxField),
}

func (a Action) NextActions() Choices {
	return nextActionsMap[a]
}

var toolActions = map[Action]Tool{
	WoolWeaver:  WeavingLoom,
	Master:      Workbench,
	Tanner:      FleshingBeam,
	LinenWeaver: WeavingLoom,
	Butcher:     SlaughteringTable,
	Potter:      PotteryWheel,
	Baker:       Oven,
	Master2:     Workbench,
}

func (a Action) Tool() Tool {
	return toolActions[a]
}

func buildAnyTool() []Action {
	return []Action{GetFishTraps, GetFleshingBeams, GetWeavingLooms, GetSlaughteringTables, GetSpades, GetShovelPaires, GetPotteryWheels, GetOvens, GetAxes, GetWorkBenches, Finish}
}

func getAnyAnimal() []Action {
	return []Action{GetSheep, GetCow, GetHorse}
}
