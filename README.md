# About

This project is a simulation of an economy with different monetary regimes,
implemented using the concept of cellular automata. The simulation demonstrates
the behavior of a network of people with different spending habits in an
inflationary and a static money supply environment.

This is strictly for satisfying my curiosity and is not meant to accurately simulate
a real economy. Eventually this project may include more complex behavior but
for now we'll keep it simple.

## Overview

The economy consists of a network of people, represented as a grid. Each person
in the network has a balance and an income (either stable or variable). People
can exchange money with their neighbors (adjacent and diagonal).

The simulation features two monetary regimes:

1. Inflationary: A specific set of people have their balances increased by some % every 10 time steps.
2. Static money supply: No additional money is introduced into the economy.

Here's an example of it in action with an "inflationary" network that adds 10%
to people's balance if they're located along the diagonals. You can see that
over time these people tend to accumulate a much higher balance and essentially
inflate away the others' savings. Surprise!

![before-histogram](https://user-images.githubusercontent.com/1846807/230923067-c958375d-3c78-4885-95bd-274d9b3f1cdd.png)
![after-histogram](https://user-images.githubusercontent.com/1846807/230923043-d03d86c7-46a2-45d8-8ccb-58f80df51ad8.png)



![inflation](https://user-images.githubusercontent.com/1846807/230749657-079f47a3-9903-403d-9e24-ed20b72adf0e.gif)

## How it works

Each network is made of "people" who start with a random balance and income (eventually including fixed,
and variable expenses). Then for every time step _t_ each person spends money with one of their neighbors
in the grid. Their spending is a limited percentage of their total balance.

In the inflationary network if you set the inflation rate to something low like 1% (1.01) then
you don't see much of an impact. If you set it to 10% (1.10) then you'll quickly see a pattern
emerge where people closest to the money printer have significantly more money than anyone else.


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
```
