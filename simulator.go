package main

import (
	"fmt"
	"log"

	. "github.com/binje/Fields/actions"
	. "github.com/binje/Fields/state"
	"github.com/pkg/profile"
)

// GameSimulator handles the game simulation logic
type GameSimulator struct {
	root *State
}

// NewGameSimulator creates a new game simulator
func NewGameSimulator() *GameSimulator {
	return &GameSimulator{
		root: Root(),
	}
}

// RunSimulation executes the game simulation
func (gs *GameSimulator) RunSimulation() (int, int) {
	defer profile.Start().Stop()
	
	gameCount := 0
	for !gs.root.RootFinished() {
		gameCount++
		game := NewGame()
		state := gs.root
		
		for !game.IsEnd() {
			availableActions := game.AvailableActions()
			state.LoadActions(availableActions)
			
			action, err := gs.selectAction(state, availableActions)
			if err != nil {
				log.Printf("Error selecting action: %v", err)
				break
			}
			
			game.DoAction(action)
			state = state.TakeAction(action)
		}
		
		state.MarkFinished(game.VP())
	}
	
	return gameCount, gs.root.GetVp()
}

// selectAction chooses the next action to perform
func (gs *GameSimulator) selectAction(s *State, availableActions []Action) (Action, error) {
	for _, action := range availableActions {
		if !s.IsFinished(action) {
			return action, nil
		}
	}
	
	s.PrintCurrNext()
	return Action(0), fmt.Errorf("no valid action available")
} 