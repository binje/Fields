package goods

import (
	"fmt"

	. "github.com/binje/Fields/actions"
)

// Good represents a type of resource in the game.
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

// GoodsTrack defines which goods are tracked for VP calculation
var GoodsTrack = map[Good]struct{}{
	Grain: {},
	Hide:  {},
	Flax:  {},
	Wool:  {},
}

const maxGood = 15

// Goods holds the player's resources.
type Goods struct {
	resources map[Good]int
}

// NewGoods creates a new Goods struct with starting resources.
func NewGoods() *Goods {
	return &Goods{
		resources: map[Good]int{
			Brick:       0,
			Clay:        4,
			Flax:        3,
			Food:        5,
			Grain:       1,
			Hide:        2,
			Leather:     0,
			LeatherWear: 0,
			Linen:       0,
			Peat:        3,
			SummerWear:  0,
			Timber:      0,
			WinterWear:  0,
			Wood:        4,
			Wool:        4,
			Woolen:      0,
		},
	}
}

// Resource Management Methods

// add increases the count of a good by 1.
func (g *Goods) add(good Good) {
	g.increase(good, 1)
}

// increase increases the count of a good by i, respecting max limits.
func (g *Goods) increase(good Good, i int) {
	newVal := g.resources[good] + i
	switch good {
	case Food:
		if newVal > 2*maxGood {
			g.resources[good] = 2 * maxGood
			return
		}
	case Grain, Hide, Flax, Wool:
		if newVal > maxGood {
			g.resources[good] = maxGood
			return
		}
	}
	g.resources[good] = newVal
}

// useN uses n units of a good, returning an error if not enough resources.
func (g *Goods) useN(good Good, n int) error {
	for i := 0; i < n; i++ {
		if err := g.use(good); err != nil {
			return err
		}
	}
	return nil
}

// use uses one unit of a good, with substitutes for Wood/Timber and Clay/Brick.
func (g *Goods) use(good Good) error {
	switch good {
	case Wood:
		return g.useWithSubstitute(Wood, Timber)
	case Clay:
		return g.useWithSubstitute(Clay, Brick)
	}
	if g.resources[good] < 1 {
		return fmt.Errorf("not enough %v", good)
	}
	g.resources[good]--
	return nil
}

// useWithSubstitute tries to use good, or sub if good is unavailable.
func (g *Goods) useWithSubstitute(good Good, sub Good) error {
	if g.resources[good] == 0 {
		good = sub
	}
	if g.resources[good] < 1 {
		return fmt.Errorf("not enough %v or substitute %v", good, sub)
	}
	g.resources[good]--
	return nil
}

// useWithSubstitutes tries to use good, or sub, or sub2 if previous are unavailable.
func (g *Goods) useWithSubstitutes(good Good, sub Good, sub2 Good) error {
	if g.resources[good] == 0 {
		good = sub
		if g.resources[good] == 0 {
			good = sub2
		}
	}
	if g.resources[good] < 1 {
		return fmt.Errorf("not enough %v, %v, or %v", good, sub, sub2)
	}
	g.resources[good]--
	return nil
}

// Inventory Methods

// NovemberInventorying processes the November inventory phase.
func (g *Goods) NovemberInventorying(food, grain, flax, wood int) (animals, bottleNeck int) {
	g.increase(Food, food)
	g.increase(Grain, grain)
	g.increase(Flax, flax)
	g.increase(Wood, wood)
	
	for i := 0; i < 2; i++ {
		if g.resources[Peat]+g.resources[Wood] == 0 {
			bottleNeck++
		} else {
			_ = g.useWithSubstitute(Peat, Wood)
		}
	}
	animals = g.foodSustenance()
	return
}

// MayInventorying processes the May inventory phase.
func (g *Goods) MayInventorying(wool int) (animals int) {
	g.increase(Wool, wool)
	animals = g.foodSustenance()
	return
}

// foodSustenance checks if there is enough food/grain for animals.
func (g *Goods) foodSustenance() (animals int) {
	for i := 0; i < 3; i++ {
		if g.resources[Food]+g.resources[Grain] == 0 {
			animals++
		} else {
			_ = g.useWithSubstitute(Food, Grain)
		}
	}
	return
}

// Action Processing

// DoAction applies the effects of an action on goods.
func (g *Goods) DoAction(a Action, i int) {
	g.gainGoods(a, i)
	g.loseGoods(a, i)
}

// gainGoods adds goods based on the action.
func (g *Goods) gainGoods(a Action, i int) {
	switch a {
	// Resource gathering
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
	
	// Employment actions
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
	
	// Travel and trade actions
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
	
	// Peat boat actions
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

// loseGoods removes goods based on the action.
func (g *Goods) loseGoods(a Action, i int) {
	switch a {
	// Crafting actions
	case WeaveWool:
		_ = g.use(Wool)
	case TanHide:
		_ = g.use(Hide)
	case WeaveLinen:
		_ = g.use(Flax)
	case UseClay, MakePot:
		_ = g.use(Clay)
	case Baker:
		_ = g.useWithSubstitute(Grain, Flax)
		_ = g.useWithSubstitute(Peat, Wood)
	case WoodTrader:
		_ = g.useWithSubstitute(Food, Grain)
	
	// Employment actions
	case Laborer, Laborer2:
		_ = g.useN(Food, 2)
	
	// Building actions
	case BuildStall:
		_ = g.useN(Clay, 2)
		_ = g.use(Grain)
	case BuildStable:
		_ = g.useN(Brick, 2)
	case BuildCart:
		_ = g.use(Wood)
		fallthrough
	case BuildDroshsky:
		_ = g.use(Wood)
		fallthrough
	case BuildHorseCart:
		_ = g.use(Wood)
		fallthrough
	case BuildCarriage, BuildWagon:
		_ = g.useN(Wood, 2)
		fallthrough
	case BuildHandcart:
		_ = g.use(Wood)
		fallthrough
	case BuildPeatBoat, BuildPlow, UseWood:
		_ = g.use(Wood)
	
	// Travel and trade actions
	case TravelBeemoor:
		_ = g.use(Peat)
	case TradeLeather:
		_ = g.use(Leather)
	case TradeHide:
		_ = g.use(Hide)
	case Trade2Grain:
		_ = g.useN(Grain, 2)
	case TradeWoolen:
		_ = g.use(Woolen)
	case TradePeat, PeatForWool, PeatForGrain, PeatForFood, PeatForFlax, PeatForHide:
		_ = g.use(Peat)
	case TradeWinterWear, UseWinterWear:
		_ = g.use(WinterWear)
	case TradeSummerWear, UseSummerWear:
		_ = g.use(SummerWear)
	case TradeLeatherWear, UseLeatherWear:
		_ = g.use(LeatherWear)
	case TradeFlax:
		_ = g.use(Flax)
	case TradeLinen:
		_ = g.use(Linen)
	case UseTimber, TradeTimber, BuildMill:
		_ = g.use(Timber)
	case TradeLinenSet:
		_ = g.use(Woolen)
		_ = g.use(Leather)
		_ = g.use(Linen)
	case TradeClothingSet:
		_ = g.use(LeatherWear)
		_ = g.use(WinterWear)
		_ = g.use(SummerWear)
	case UseBrick, BuildTextileHouse:
		_ = g.use(Brick)
	
	// Building actions (continued)
	case BuildFarmersHouse, BuildPlowMakersWorkShop, BuildNovicesHut, BuildWorkShop, BuildWeavingParlor, BuildColonistsHouse, BuildCarpentersWorkshop, BuildSchnappsDistillery, BuildLoadingStation, BuildLitterStorage, BuildWoodTrader:
		_ = g.use(Grain)
	case BuildTurnery, BuildSmokehouse, BuildSmithy, BuildCooperage, BuildBakehouse:
		_ = g.use(Timber)
		_ = g.use(Brick)
	case BuildWeavingMill:
		_ = g.useN(Brick, 2)
	case Use8Flax:
		_ = g.useN(Flax, 8)
	case Use10Flax:
		_ = g.useN(Flax, 10)
	case Use8Grain:
		_ = g.useN(Grain, 8)
	case Use10Wool:
		_ = g.useN(Wool, 10)
	case BuildSaddlery:
		_ = g.useN(Timber, 2)
		_ = g.useN(Leather, 3)
	case BuildWaterfrontHouse:
		_ = g.useN(Brick, 2)
		_ = g.useN(Food, 25)
	case BuildJoinery:
		_ = g.useN(Timber, 2)
		_ = g.useN(Grain, 5)
	case BuildPottersInn, BuildFarmersInn, BuildJunkDealersInn, BuildGulfHouseInn, BuildMilkHouseInn:
		_ = g.useN(Food, 9)
	case BuildVillageChurch, BuildLutetsburgCastle, BuildBerumCastle:
		_ = g.useN(Timber, 3)
		_ = g.useN(Brick, 3)
		_ = g.useN(Food, 15)
	}
}

// Victory Points Calculation

// Vp returns the total victory points for goods.
func (g *Goods) Vp() int {
	return g.calculateBasicVp() + g.calculateClothingVp()
}

// calculateBasicVp calculates VP from basic goods.
func (g *Goods) calculateBasicVp() int {
	return g.resources[Timber]/2 + g.resources[Brick] + g.resources[Linen] + g.resources[Woolen] + g.resources[Leather]
}

// calculateClothingVp calculates VP from clothing items.
func (g *Goods) calculateClothingVp() int {
	return 2 * (g.resources[SummerWear] + g.resources[WinterWear] + g.resources[LeatherWear])
}

// GoodsTrackVp returns the victory points for the goods track.
func (g *Goods) GoodsTrackVp() int {
	v := 0
	if g.resources[Food] > maxGood {
		v += 3
		g.resources[Food] -= maxGood
	}
	
	// Use GoodsTrack to iterate over tracked goods
	for good := range GoodsTrack {
		v += calculateGoodsTrackVp(g.resources[good])
	}
	
	return v
}

// calculateGoodsTrackVp calculates VP for a single good based on quantity.
func calculateGoodsTrackVp(quantity int) int {
	if quantity < 7 {
		return 0
	}
	if quantity <= 10 {
		return 1
	}
	if quantity <= 14 {
		return 2
	}
	return 3
}
