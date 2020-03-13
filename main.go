package main

import (
	"fmt"

	. "github.com/binje/Fields/actions"
)

func main() {
	var input Action

	g := NewGame()

	for !g.IsEnd() {
		fmt.Println()
		g.PrintDate()
		fmt.Println("Options:")
		availableActions := g.AvailableActions()
		for _, action := range availableActions {
			fmt.Printf("%d: %s\n", action, action)
		}
		fmt.Println()
		fmt.Scan(&input)
		if canDo(input, availableActions) {
			g.DoAction(input)
			g.NextMonth()
		} else {
			fmt.Println("Cannot do that action")
		}
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
