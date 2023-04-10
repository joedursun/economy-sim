package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joedursun/economy-sim/models"
)

const (
	gridSize          = 20
	numCentralBankers = 5
	ethFee            = 10.00
)

func main() {
	network := models.NewEthNetwork(gridSize, ethFee)
	// network := models.NewInflationaryNetwork(gridSize, numCentralBankers, 1.10)

	network.PrintBalanceHistogram(network.People(), 100)
	for network.SimulateTransactions() {
		// clearScreen()
		// fmt.Printf("Time step %d | Money in circulation %.2f:\n", network.TimeStep, network.MoneyInCirculation())
		// network.PrintState()

		// time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("-----------------------------------")
	network.PrintBalanceHistogram(network.People(), 100)
}

// clearScreen clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
