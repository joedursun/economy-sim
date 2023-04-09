package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/joedursun/economy-sim/models"
)

const (
	gridSize = 20
)

func main() {
	network := models.NewInflationaryNetwork(gridSize, 1.10)

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
