package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/joedursun/economy-sim/models"
)

const (
	numActors = 10
)

func main() {
	gridSize := 20
	network := models.NewInflationaryNetwork(gridSize, 1.15)

	// Simulate transactions and print state for 10 time steps
	for network.SimulateTransactions() {
		clearScreen()
		fmt.Printf("Time step %d:\n", network.TimeStep)
		network.PrintState()

		time.Sleep(50 * time.Millisecond)
	}
}

// clearScreen clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
