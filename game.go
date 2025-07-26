package main

import (
	. "github.com/binje/Fields/actions"
	. "github.com/binje/Fields/goods"
	. "github.com/binje/Fields/homeBoard"
	. "github.com/binje/Fields/player"
	. "github.com/binje/Fields/stable"
	. "github.com/binje/Fields/time"
)

// Game represents the main game state
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

// NewGame creates a new game instance
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

// VP calculates the total victory points
func (g *Game) VP() int {
	vp := 0
	vp += g.goods.Vp()
	vp += g.stable.Vp()
	vp += g.player.ToolsVp()
	vp += g.goods.GoodsTrackVp()
	
	// Bonus VP for depot
	if g.home.HasDepot() {
		vp += g.goods.GoodsTrackVp()
	}
	
	vp += g.home.Vp()
	vp += g.home.AnimalVp()
	vp -= g.bottleneck
	return vp
}

// AvailableActions returns all available actions for the current game state
func (g *Game) AvailableActions() []Action {
	if len(g.choices) != 0 {
		return g.choices[0]
	}
	
	// Handle overpopulation
	if g.home.IsOverPopulated() {
		return g.home.GetSlaughterOptions()
	}
	
	actions := getEmploymentActions(g.calendar.Season(), g.otherSideUsed)
	
	// Add peat boat action if available
	if g.stable.NumPeatBoats() > 0 {
		actions = append(actions, UsePeatBoat)
	}
	
	return removeDuplicates(actions)
}

// getEmploymentActions returns employment actions based on season and usage
func getEmploymentActions(season Season, otherSideUsed bool) []Action {
	switch season {
	case NovemberInventorying:
		return []Action{} // No November Inventory Tasks
	case MayInventorying:
		return []Action{} // No May Inventory Tasks
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

// removeDuplicates efficiently removes duplicate actions
func removeDuplicates(actions []Action) []Action {
	if len(actions) == 0 {
		return actions
	}
	
	seen := make(map[Action]bool)
	result := make([]Action, 0, len(actions))
	
	for _, action := range actions {
		if !seen[action] {
			seen[action] = true
			result = append(result, action)
		}
	}
	
	return result
}

// IsEnd checks if the game has ended
func (g *Game) IsEnd() bool {
	return g.calendar.EndOfTheWorld()
}

// DoAction executes an action and updates the game state
func (g *Game) DoAction(action Action) {
	// Handle sequential actions
	if len(g.choices) != 0 {
		g.choices = g.choices[1:]
	}

	// Process action choices
	choices := action.NextActions()
	if tool := action.Tool(); tool != NoTool {
		choices.UseTool(g.player.GetToolCount(tool))
	}
	g.choices = choices
	
	// Execute goods action
	g.DoGoodsAction(action)

	// Handle employment actions
	if _, isEmployment := AllEmployment[action]; isEmployment {
		g.usedActions[action] = struct{}{}
		g.calendar.NextMonth()
		if OffSeason(g.calendar.Season(), action) {
			g.otherSideUsed = true
		}
	}

	// Handle inventory actions
	g.handleInventoryActions()
}

// handleInventoryActions processes November and May inventory actions
func (g *Game) handleInventoryActions() {
	switch g.calendar.Season() {
	case NovemberInventorying:
		food, grain, flax, wood := g.home.NovemberInventory()
		animalsToKill, bottleneck := g.goods.NovemberInventorying(food, grain, flax, wood)
		g.bottleneck += bottleneck
		g.calendar.NextMonth()
		
		// Handle animal slaughter
		if animalsToKill > 0 {
			g.handleAnimalSlaughter(animalsToKill)
		}
		
	case MayInventorying:
		wool := g.home.MayInventory()
		animalsToKill := g.goods.MayInventorying(wool)
		g.calendar.NextMonth()
		
		// Handle animal slaughter
		if animalsToKill > 0 {
			g.handleAnimalSlaughter(animalsToKill)
		}
		
	default:
		return
	}
}

// handleAnimalSlaughter processes animal slaughter based on available animals
func (g *Game) handleAnimalSlaughter(animalsToKill int) {
	numAnimals := g.home.NumAnimals()
	if animalsToKill > numAnimals {
		g.bottleneck += animalsToKill - numAnimals
		g.home.KillAllAnimals()
	} else {
		// TODO: Implement selective animal slaughter
		// g.choices = append(g.choices, createSlaughterChoices(animalsToKill))
	}
}

// DoGoodsAction executes goods-related actions
func (g *Game) DoGoodsAction(action Action) {
	switch action {
	case Fisherman:
		g.goods.DoAction(action, g.player.GetToolCount(FishTrap))
	case PeatCutter:
		g.goods.DoAction(action, g.home.CutPeat(g.player.GetToolCount(Spade)))
	case ClayWorker:
		g.goods.DoAction(action, g.player.GetToolCount(Shovel))
	case Woodcutter:
		g.goods.DoAction(action, g.player.GetToolCount(Axe))
	case PeatBoatman:
		g.goods.DoAction(action, g.stable.NumPeatBoats())
	case Grocer2:
		g.goods.DoAction(action, g.home.CutPeat(1))
	default:
		g.goods.DoAction(action, 0)
	}
}
