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
	network := models.NewStaticNetwork(gridSize) // e.g. BTC
	// network := models.NewEthNetwork(gridSize, ethFee) // e.g. ETH
	// network := models.NewInflationaryNetwork(gridSize, numCentralBankers, 1.10) // e.g. Fiat

	fmt.Printf("Money in circulation %.2f:\n", network.MoneyInCirculation())
	network.PrintBalanceHistogram(network.People(), 100)
	for network.SimulateTransactions() {
		// clearScreen()
		// fmt.Printf("Time step %d | Money in circulation %.2f:\n", network.TimeStep, network.MoneyInCirculation())
		// network.PrintState()

		// time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("-----------------------------------")
	network.PrintBalanceHistogram(network.People(), 100)
	fmt.Printf("Money in circulation %.2f:\n", network.MoneyInCirculation())
}

// clearScreen clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
