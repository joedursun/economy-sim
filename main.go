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
	gridSize := 5
	network := models.NewNetwork(gridSize)

	// Simulate transactions and print state for 10 time steps
	for t := 1; t <= 10; t++ {
		clearScreen()
		fmt.Printf("Time step %d:\n", t)
		network.PrintState()
		network.SimulateTransactions()
		time.Sleep(time.Second)
	}
}

// clearScreen clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
