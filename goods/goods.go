package goods

import (
	. "github.com/binje/Fields/actions"
)

type Good int

const (
	Undefined Good = iota
	Wood
	Timber
	Clay
	Brick
	Peat
	Food
	Grain
	Hide
	Flax
	Wool
	Linen
	Woolen
	Leather
	SummerWear
	WinterWear
	LeatherWear
)

var maxGood int = 15

var goodsTrack = map[Good]struct{}{
	Grain: struct{}{},
	Hide:  struct{}{},
	Flax:  struct{}{},
	Wool:  struct{}{}}

type Goods struct {
	m map[Good]int
}

func NewGoods() *Goods {
	return &Goods{map[Good]int{
		Wood:  4,
		Clay:  4,
		Peat:  3,
		Food:  5,
		Wool:  4,
		Flax:  3,
		Hide:  2,
		Grain: 1,
	}}
}

func (g *Goods) add(good Good) {
	g.increase(good, 1)
}

//TODO put cap on goods
func (g Goods) increase(good Good, i int) {
	newVal := g.m[good] + i
	if good == Food {
		if newVal > 2*maxGood {
			newVal = 2 * maxGood
		}
	} else if _, ok := goodsTrack[good]; ok {
		if newVal > maxGood {
			newVal = maxGood
		}
	}
	g.m[good] = newVal
}

func (g *Goods) useN(good Good, n int) {
	for i := 0; i < n; i++ {
		g.use(good)
	}
}

func (g *Goods) use(good Good) {
	if good == Wood {
		g.useWithSubstitute(Wood, Timber)
		return
	} else if good == Clay {
		g.useWithSubstitute(Clay, Brick)
		return
	}
	if g.m[good] < 1 {
		panic("YOU CAN'T DO THIS, YOU NEEDED that resource")
	}
	g.m[good] -= 1
}

func (g *Goods) useWithSubstitute(good Good, sub Good) {
	if g.m[good] == 0 {
		good = sub
	}
	if g.m[good] < 1 {
		panic("YOU CAN'T DO THIS, YOU NEEDED that resource")
	}
	g.m[good] -= 1
}

func (g *Goods) useWithSubstitutes(good Good, sub Good, sub2 Good) {
	if g.m[good] == 0 {
		good = sub
		if g.m[good] == 0 {
			good = sub2
		}
	}
	if g.m[good] < 1 {
		panic("YOU CAN'T DO THIS, YOU NEEDED that resource")
	}
	g.m[good] -= 1
}

func (g *Goods) NovemberInventorying(food, grain, flax, wood int) (animals int, bottleNeck int) {
	g.increase(Food, food)
	g.increase(Grain, grain)
	g.increase(Flax, flax)
	g.increase(Wood, wood)
	for i := 0; i < 2; i++ {
		if g.m[Peat]+g.m[Wood] == 0 {
			bottleNeck++
		} else {
			g.useWithSubstitute(Peat, Wood)
		}
	}
	animals = g.foodSustenance()
	return
}

func (g *Goods) MayInventorying(wool int) (animals int) {
	g.increase(Wool, wool)
	animals = g.foodSustenance()
	return
}

func (g *Goods) foodSustenance() (animals int) {
	for i := 0; i < 3; i++ {
		if g.m[Food]+g.m[Grain] == 0 {
			animals++
		} else {
			g.useWithSubstitute(Food, Grain)
		}
	}
	return
}

func (g *Goods) DoAction(a Action, i int) {
	g.gainGoods(a, i)
	g.loseGoods(a, i)
}

func (g *Goods) gainGoods(a Action, i int) {
	switch a {
	case GetWood:
		g.add(Wood)
	case Get4Wood:
		g.increase(Wood, 4)
	case GetClay:
		g.add(Clay)
	case GetTimber:
		g.add(Timber)
	case GetBrick:
		g.add(Brick)
	case Fisherman:
		g.increase(Food, i)
	case Grocer:
		g.add(Grain)
		g.add(Leather)
	case WeaveWool:
		g.add(Woolen)
	case PeatCutter:
		g.increase(Peat, i)
	case ClayWorker:
		g.increase(Clay, i)
	case Woodcutter:
		g.increase(Wood, i)
	case PeatBoatman:
		g.increase(Peat, 3+i)
	case TanHide:
		g.add(Leather)
	case WeaveLinen:
		g.add(Linen)
	case ButcherCow:
		g.add(Food)
		fallthrough
	case ButcherSheep, ButcherHorse:
		g.increase(Food, 3)
		g.increase(Hide, 2)
	case CattleTrader:
		g.increase(Grain, 2)
	case Grocer2:
		g.increase(Peat, i)
		g.add(Wood)
		g.add(Clay)
	case BuildersMerchant:
		g.increase(Hide, 2)
	case MakePot:
		g.add(Peat)
		g.increase(Food, 3)
	// Travel
	case TravelHage, PeatForFood:
		g.add(Food)
	case TravelBeemoor, TradeHide, TradeFlax:
		g.increase(Food, 2)
	case TradePeat, TradeLinen:
		g.increase(Food, 3)
	case TradeSheep, TradeLeather, TradeAnimal, Trade2Grain, TradeWoolen:
		g.increase(Food, 4)
	case TradePeatBoat, TradeHorse, TradeCow, TradeTimber:
		g.increase(Food, 5)
	case Bake, TradeClothing, TradeSummerWear:
		g.increase(Food, 6)
	case TradeWinterWear, TradeLeatherWear:
		g.increase(Food, 7)
	case TravelDornum:
		g.increase(Food, 8)
	case Trade2Animal:
		g.increase(Food, 9)
	case TradeLinenSet:
		g.increase(Food, 12)
	case TradeClothingSet:
		g.increase(Food, 30)
	//PeatBoat
	case PeatForWool:
		g.add(Wool)
	case PeatForGrain:
		g.add(Grain)
	case PeatForFlax:
		g.add(Flax)
	case PeatForHide:
		g.add(Hide)
	}
}

func (g *Goods) loseGoods(a Action, i int) {
	switch a {
	case WeaveWool:
		g.use(Wool)
	case Laborer, Laborer2:
		g.use(Food)
		g.use(Food)
	case TanHide:
		g.use(Hide)
	case WeaveLinen:
		g.use(Flax)
	case UseClay, MakePot:
		g.use(Clay)
	case Baker:
		g.useWithSubstitute(Grain, Flax)
		g.useWithSubstitute(Peat, Wood)
	case WoodTrader:
		g.useWithSubstitute(Food, Grain)
	case BuildStall:
		g.use(Clay)
		g.use(Clay)
		g.use(Grain)
	case BuildStable:
		g.use(Brick)
		g.use(Brick)
	case BuildCart:
		g.use(Wood)
		fallthrough
	case BuildDroshsky:
		g.use(Wood)
		fallthrough
	case BuildHorseCart:
		g.use(Wood)
		fallthrough
	case BuildCarriage, BuildWagon:
		g.use(Wood)
		g.use(Wood)
		fallthrough
	case BuildHandcart:
		g.use(Wood)
		fallthrough
	case BuildPeatBoat, BuildPlow, UseWood:
		g.use(Wood)
	case TravelBeemoor:
		g.use(Peat)
	case TradeLeather:
		g.use(Leather)
	case TradeHide:
		g.use(Hide)
	case Trade2Grain:
		g.use(Grain)
		g.use(Grain)
	case TradeWoolen:
		g.use(Woolen)
	case TradePeat, PeatForWool, PeatForGrain, PeatForFood, PeatForFlax, PeatForHide:
		g.use(Peat)
	case TradeWinterWear, UseWinterWear:
		g.use(WinterWear)
	case TradeSummerWear, UseSummerWear:
		g.use(SummerWear)
	case TradeLeatherWear, UseLeatherWear:
		g.use(LeatherWear)
	case TradeFlax:
		g.use(Flax)
	case TradeLinen:
		g.use(Linen)
	case UseTimber, TradeTimber, BuildMill:
		g.use(Timber)
	case TradeLinenSet:
		g.use(Woolen)
		g.use(Leather)
		g.use(Linen)
	case TradeClothingSet:
		g.use(LeatherWear)
		g.use(WinterWear)
		g.use(SummerWear)
	case UseBrick, BuildTextileHouse:
		g.use(Brick)
		// TODO buildings
	case BuildFarmersHouse, BuildPlowMakersWorkShop, BuildNovicesHut, BuildWorkShop, BuildWeavingParlor, BuildColonistsHouse, BuildCarpentersWorkshop, BuildSchnappsDistillery, BuildLoadingStation, BuildLitterStorage, BuildWoodTrader:
		g.use(Grain)
	case BuildTurnery, BuildSmokehouse, BuildSmithy, BuildCooperage, BuildBakehouse:
		g.use(Timber)
		g.use(Brick)
	case BuildWeavingMill:
		g.useN(Brick, 2)
	case Use8Flax:
		g.useN(Flax, 8)
	case Use10Flax:
		g.useN(Flax, 10)
	case Use8Grain:
		g.useN(Grain, 8)
	case Use10Wool:
		g.useN(Wool, 10)
	case BuildSaddlery:
		g.useN(Timber, 2)
		g.useN(Leather, 3)
	case BuildWaterfrontHouse:
		g.useN(Brick, 2)
		g.useN(Food, 25)
	case BuildJoinery:
		g.useN(Timber, 2)
		g.useN(Grain, 5)
	case BuildPottersInn, BuildFarmersInn, BuildJunkDealersInn, BuildGulfHouseInn, BuildMilkHouseInn:
		g.useN(Food, 9)
	case BuildVillageChurch, BuildLutetsburgCastle, BuildBerumCastle:
		g.useN(Timber, 3)
		g.useN(Brick, 3)
		g.useN(Food, 15)
	}
}
