package main

import (
	"fmt"
	"math/rand"

	. "github.com/binje/Fields/actions"
)

func main() {
	//var input Action

	g := NewGame()

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

	for !g.IsEnd() {
		availableActions := g.AvailableActions()
		randomAction := availableActions[rand.Intn(len(availableActions))]
		fmt.Printf("%s\n", randomAction)
		g.DoAction(randomAction)
	}

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
