package main

import (
	. "github.com/binje/Fields/actions"
	. "github.com/binje/Fields/goods"
	. "github.com/binje/Fields/homeBoard"
	. "github.com/binje/Fields/player"
	. "github.com/binje/Fields/stable"
	. "github.com/binje/Fields/time"
)

type Game struct {
	player        *Player
	home          *HomeBoard
	goods         *Goods
	stable        *Stable
	calendar      *Calendar
	usedActions   map[Action]struct{}
	otherSideUsed bool
	bottleneck    int
	choices       [][]Action
}

func NewGame() Game {
	return Game{
		player:        NewPlayer(),
		home:          NewHomeBoard(),
		goods:         NewGoods(),
		stable:        NewStable(),
		calendar:      NewCalendar(),
		otherSideUsed: false,
		usedActions:   make(map[Action]struct{}),
	}
}

func (g *Game) AvailableActions() (actions []Action) {
	if len(g.choices) != 0 {
		return g.choices[0]
	}
	for g.home.IsOverPopulated() {
		return g.home.GetSlaughterOptions()
	}
	actions = getEmploymentActions(g.calendar.Season(), g.otherSideUsed)

	// remove used actions
	for a, _ := range g.usedActions {
		actions = remove(actions, a)
	}
	if g.stable.NumPeatBoats() > 0 {
		actions = append(actions, UsePeatBoat)
	}
	return
}

func getEmploymentActions(season Season, otherSideUsed bool) []Action {
	switch season {
	case NovemberInventorying:
		panic("No November Inventory Tasks")
	case MayInventorying:
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

func (g *Game) IsEnd() bool {
	return g.calendar.EndOfTheWorld()
}

func (g *Game) DoAction(action Action) {
	// if multple actions must be made in sequence, iterate to next choice
	if len(g.choices) != 0 {
		g.choices = g.choices[1:]
	}

	choices := action.NextActions()
	if tool := action.Tool(); tool != NoTool {
		choices.UseTool(g.player.GetToolCount(tool))
	}
	g.choices = choices
	g.DoGoodsAction(action)

	//TODO remember last employment? for things like master or imitator

	// Employment actions cannot be deuplicated (imitator aside)
	// Must be done at end becasue increasing month prints the new date
	if _, ok := AllEmployment[action]; ok {
		g.usedActions[action] = struct{}{}
		g.calendar.NextMonth()
		if OffSeason(g.calendar.Season(), action) {
			g.otherSideUsed = true
		}
	}

	var animalsToKill, bottleneck int
	if g.calendar.Season() == NovemberInventorying {
		food, grain, flax, wood := g.home.NovemberInventory()
		animalsToKill, bottleneck = g.goods.NovemberInventorying(food, grain, flax, wood)
		g.bottleneck += bottleneck
		g.calendar.NextMonth()
	} else if g.calendar.Season() == MayInventorying {
		wool := g.home.MayInventory()
		animalsToKill = g.goods.MayInventorying(wool)
		g.calendar.NextMonth()
	}
	if animalsToKill > 0 {
		numAnimals := g.home.NumAnimals()
		if animalsToKill > numAnimals {
			g.bottleneck += animalsToKill - numAnimals
			g.home.KillAllAnimals()
		} else {
			//TODO add slaughter animal n times g.choices =
		}
	}
}

func (g *Game) DoGoodsAction(a Action) {
	switch a {
	case Fisherman:
		g.goods.DoAction(a, g.player.GetToolCount(FishTrap))
	case PeatCutter:
		g.goods.DoAction(a, g.home.CutPeat(g.player.GetToolCount(Spade)))
	case ClayWorker:
		g.goods.DoAction(a, g.player.GetToolCount(Shovel))
	case Woodcutter:
		g.goods.DoAction(a, g.player.GetToolCount(Axe))
	case PeatBoatman:
		g.goods.DoAction(a, g.stable.NumPeatBoats())
	case Grocer2:
		g.goods.DoAction(a, g.home.CutPeat(1))
	default:
		g.goods.DoAction(a, 0)
	}
}
