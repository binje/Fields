package main

import (
	"fmt"

	. "github.com/binje/Fields/actions"
	. "github.com/binje/Fields/goods"
	. "github.com/binje/Fields/stable"
	. "github.com/binje/Fields/time"
)

type Game struct {
	player        *Player
	homeBoard     *HomeBoard
	goods         *Goods
	stable        *Stable
	halfYear      int
	month         Month
	usedActions   map[Action]struct{}
	otherSideUsed bool
	choices       [][]Action
}

func NewGame() Game {
	return Game{
		player:        NewPlayer(),
		homeBoard:     NewHomeBoard(),
		goods:         NewGoods(),
		stable:        NewStable(),
		halfYear:      1,
		month:         0,
		otherSideUsed: false,
	}
}

func (g *Game) AvailableActions() (actions []Action) {
	if len(g.choices) != 0 {
		return g.choices[0]
	}
	actions = getEmploymentActions(g.month, g.otherSideUsed)
	// remove used actions
	//TODO test for sequential actions removal
	for a, _ := range g.usedActions {
		actions = remove(actions, a)
	}
	return
}

func getEmploymentActions(m Month, otherSideUsed bool) []Action {
	switch m.Season() {
	case NovemberInventorying:
		//TODO choose how to pay
		panic("No November Inventory Tasks")
	case MayInventorying:
		//TODO choose how to breed animals
		panic("No May Inventory Tasks")
	case JunePreperations:
		if otherSideUsed {
			return SummerEmploymentArray
		}
	case DecemberPreperations:
		if otherSideUsed {
			return WinterEmploymentArray
		}
	}
	return AllEmploymentArray

}

func remove(s []Action, action Action) []Action {
	for i, a := range s {
		if a == action {
			s[i] = s[len(s)-1]
			return s[:len(s)-1]
		}
	}
	return s
}

func (g *Game) NextMonth() {
	g.month++
	if g.month > May {
		g.month = July
	}
	if g.month == November {
		g.NovemberInventory()
	}
	if g.month == May {
		g.MayInventory()
	}
}

func (g *Game) NovemberInventory() {
	g.halfYear++
	g.otherSideUsed = false
	g.usedActions = map[Action]struct{}{}

	sheep := g.homeBoard.GetAnimal(Sheep)
	if sheep > 2 {
		g.goods.Add(Food)
	}
	if sheep > 5 {
		g.goods.Add(Food)
	}
	if sheep > 8 {
		g.goods.Add(Food)
	}

	cows := g.homeBoard.GetAnimal(Cow)
	if cows > 1 {
		g.goods.Add(Food)
	}
	if cows > 3 {
		g.goods.Add(Food)
	}
	if cows > 5 {
		g.goods.Add(Food)
	}

	//TODO catch error here
	//TODO can spend animal?
	g.goods.UseWithSubstitute(Food, Grain, 3)
	g.goods.UseWithSubstitute(Peat, Wood, 2)

}

func (g *Game) MayInventory() {
	g.halfYear++
	g.otherSideUsed = false
	g.usedActions = map[Action]struct{}{}

	//TODO stables := g.homeBoard.getNumStables()
}

func (g *Game) IsEnd() bool {
	return g.halfYear == 10
}

func (g *Game) PrintDate() {
	fmt.Println(g.halfYear, g.month)
}

func (g *Game) DoAction(action Action) {
	if len(g.choices) != 0 {
		g.choices = g.choices[1:]
	}

	// Employment actions cannot be deuplicated (imitator aside)
	if _, ok := AllEmployment[action]; ok {
		g.usedActions[action] = struct{}{}
		if OffSeason(g.month, action) {
			g.otherSideUsed = true
		}
	}

	switch action {
	case GetSheep, GetCow, GetHorse, GetWood, Get4Wood, GetTimber, GetClay, GetBrick, GetFlaxField, GetGrainField:
		g.getResource(action)
	case Fisherman:
		g.fisherman()
	case Grocer:
		g.grocer()
	case Colonist:
		g.colonist()
	case PeatCutter:
		g.peatCutter()
	default:
		fmt.Println(action)
		panic("Action Not Known")
	}
}

func (g *Game) addChoice(choices ...Action) {
	g.choices = append(g.choices, append(choices, Finish))
}

func (g *Game) addMandatoryChoice(choices ...Action) {
	g.choices = append(g.choices, choices)
}

func (g *Game) fisherman() {
	g.player.IncreaseTool(FishTraps)
	g.goods.Increase(Food, g.player.GetToolCount(FishTraps))
}

func (g *Game) grocer() {
	g.goods.Add(Grain)
	g.goods.Add(Hide)
	g.addChoice(GetSheep, GetCow, GetHorse, GetTimber, GetBrick)
}

func (g *Game) WoolWeaver() {
	//TODO add weave wool
	//g.addChoice(WeaveWool)
}

func (g *Game) colonist() {
	g.homeBoard.Breed(Cow)
	g.homeBoard.FlipMoor()
}

func (g *Game) peatCutter() {
	g.homeBoard.CutPeat(g.player.GetToolCount(Spades))
}

func (g *Game) dikeBuilder() {
	for i := 0; i < g.player.GetToolCount(Spades)/2; i++ {
		g.homeBoard.MoveDike()
	}
	g.addChoice(GetSheep, GetCow)
}

func (g *Game) clayWorker() {
	g.goods.Increase(Clay, g.player.GetToolCount(Shovels))
}

func (g *Game) farmer() {
	g.stable.AddVehicle(Plow)
	for i := 0; i < g.stable.GetNumPlows(); i++ {
		g.addChoice(GetFlaxField, GetGrainField)
	}
}

func (g *Game) forester() {
	g.goods.Use(Food, 1)
	g.addChoice(GetForest, Builder)
}

func (g *Game) woodcutter() {
	g.goods.Increase(Wood, g.player.GetToolCount(Axes))
}

//TODO master

func (g *Game) carpenter() {

	// TODO
}

func (g *Game) builder() {
	g.addChoice(g.getBuildings()...)
}

func (g *Game) warden() {
	//TODO
}

func (g *Game) laborer() {
	//TODO
}

func (g *Game) peatBoatman() {
	g.goods.Increase(Peat, 3)
	//TODO add peat per peat boat
}

//TODO tanner
//TODO linenweaver
//TODO butcher

func (g *Game) cattleTrader() {
	g.goods.Increase(Grain, 2)
	g.homeBoard.Breed(Sheep)
	g.addChoice(GetCow, GetHorse)
}

func (g *Game) grocer2() {
	g.homeBoard.CutPeat(1)
	g.goods.Add(Wood)
	g.goods.Add(Clay)
	g.addChoice(GetSheep, GetCow, GetHorse)
}

func (g *Game) buildersMerchant() {
	g.goods.Increase(Hide, 2)
	g.addChoice(GetWood, GetClay)
	g.addChoice(GetTimber, GetBrick)
}

//TODO potter
//TODO baker

func (g *Game) woodTrader() {
	g.goods.UseWithSubstitute(Food, Grain, 1)
	g.addChoice(Get4Wood, Builder)
}

func (g *Game) wainwright() {
	g.addChoice(BuildVehicle)
	g.addChoice(BuildPeatBoat)
}

func (g *Game) buildVehicle() {
	g.addChoice()
}

func (g *Game) buildPeatBoat() {
	g.goods.Use(Wood, 1)
	g.addMandatoryChoice(UseCow, UseHorse)
}

func (g *Game) dikeWarder() {
	g.homeBoard.MoveDike()
	g.warden()
}

func (g *Game) carpenter2() {
	g.carpenter()
}

func (g *Game) laborer2() {
	//TODO NOT THE SAME
	g.laborer()
}

func (g *Game) getBuildings() (buildings []Action) {
	//TODO have buildings
	return
}

func (g *Game) getResource(a Action) {
	switch a {
	case GetSheep:
		g.homeBoard.Breed(Sheep)
	case GetCow:
		g.homeBoard.Breed(Cow)
	case GetHorse:
		g.homeBoard.Breed(Horse)
	case GetWood:
		g.goods.Add(Wood)
	case Get4Wood:
		g.goods.Increase(Wood, 4)
	case GetTimber:
		g.goods.Add(Timber)
	case GetClay:
		g.goods.Add(Clay)
	case GetBrick:
		g.goods.Add(Brick)
	case GetFlaxField, GetGrainField, GetForest:
		g.homeBoard.AddTile()
	default:
		panic("Got a nonexistant resource?")
	}
}
