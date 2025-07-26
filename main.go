package main

import (
	"fmt"
)

func main() {
	simulator := NewGameSimulator()
	gameCount, maxVP := simulator.RunSimulation()
	
	fmt.Printf("Finished: %d games\n", gameCount)
	fmt.Printf("Max VP: %d\n", maxVP)
}
