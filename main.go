package main

import (
	"fmt"

	. "github.com/binje/Fields/actions"
	. "github.com/binje/Fields/state"
)

func main() {
	//var input Action

	/*
		for !g.IsEnd() {
			fmt.Println()
			fmt.Println("Options:")
			availableActions := g.AvailableActions()
			for _, action := range availableActions {
				fmt.Printf("%d: %s\n", action, action)
			}
			fmt.Println()
			fmt.Scan(&input)
			if canDo(input, availableActions) {
				g.DoAction(input)
			} else {
				fmt.Println("Cannot do that action")
			}
		}
	*/

	root := Root()
	i := 0
	for !root.RootFinished() {
		i++
		g := NewGame()
		state := root
		for !g.IsEnd() {
			availableActions := g.AvailableActions()
			// needed to know hwen all actions have been taken
			state.LoadActions(availableActions)

			//randomAction := availableActions[rand.Intn(len(availableActions))]
			// take next action
			action := selectAction(state, availableActions)

			fmt.Printf("Taking action: %s\n", action)
			g.DoAction(action)
			state = state.TakeAction(action)

		}
		state.MarkFinished()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("finished ", i)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}

}

func selectAction(s *State, aa []Action) Action {
	for _, a := range aa {
		if !s.IsFinished(a) {
			// walk state machine
			return a
		}
	}
	s.PrintCurrNext()
	panic("no action")
}

func canDo(action Action, actions []Action) bool {
	for _, a := range actions {
		if action == a {
			return true
		}
	}
	return false
}

// TODO check if you have room for vehicle, throw away a vehicle if not

// TODO asset uniqueness on cureent choices?

// TODO travel exchanges? linking

//TODO all of buildings

//TODO breeding
