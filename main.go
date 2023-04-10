package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joedursun/economy-sim/models"
)

func main() {
	options := models.SimulatorOptions{
		GridSize:           20,
		NumCentralBankers:  5,
		BurnFee:            10.00,
		MinSpend:           10.0,
		MaxSpendPercentage: 0.5,
		SpendProbability:   0.5,
		InflationRate:      1.10,
	}

	// network := models.NewStaticNetwork(options) // e.g. BTC
	// network := models.NewEthNetwork(options) // e.g. ETH
	network := models.NewInflationaryNetwork(options) // e.g. Fiat

	fmt.Printf("Money in circulation %.2f:\n", network.MoneyInCirculation())
	network.SaveHistogram(100, "before-histogram.png")
	for network.SimulateTransactions() {
		// Uncomment these lines to see the simulation in action
		// clearScreen()
		// fmt.Printf("Time step %d | Money in circulation %.2f:\n", network.TimeStep, network.MoneyInCirculation())
		// network.PrintState()

		// time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("-----------------------------------")
	network.SaveHistogram(100, "after-histogram.png")
	fmt.Printf("Money in circulation %.2f:\n", network.MoneyInCirculation())
}

// clearScreen clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
