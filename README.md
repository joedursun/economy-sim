# About

This project is a simulation of an economy with different monetary regimes,
implemented using the concept of cellular automata. The simulation demonstrates
the behavior of a network of people with different spending habits in an
inflationary and a static money supply environment.

## Overview

The economy consists of a network of people, represented as a grid. Each person
in the network has a balance, an income (either stable or variable), and expenses
(fixed + variable). People can exchange money with their neighbors (adjacent and diagonal).

The simulation features two monetary regimes:

1. Inflationary: A specific set of people have their balances increased by some % every 10 time steps.
2. Static money supply: No additional money is introduced into the economy.

## Example Usage

```go
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
	network := models.NewInflationaryNetwork(gridSize, 1.15)

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
```
