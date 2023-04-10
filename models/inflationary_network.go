package models

import (
	"math/rand"

	"github.com/fatih/color"
)

const (
	maxTimeSteps = 10000
)

var (
	blockChars = []string{" ", "░", "▒", "▓", "█"}
)

// InflationaryNetwork struct representing the network of people
type InflationaryNetwork struct {
	StaticNetwork
	inflationRate float64
	boostedPeople []struct {
		row int
		col int
	}
}

// NewInflationaryNetwork creates a new network with gridSize x gridSize people
func NewInflationaryNetwork(gridSize, numCentralBankers int, inflationRate float64) InflationaryNetwork {
	boostedPeople := make([]struct {
		row int
		col int
	}, numCentralBankers)
	for i := 0; i < numCentralBankers; i++ {
		boostedPeople[i] = struct {
			row int
			col int
		}{
			row: rand.Intn(gridSize),
			col: rand.Intn(gridSize),
		}
	}

	return InflationaryNetwork{
		StaticNetwork: NewStaticNetwork(gridSize),
		inflationRate: inflationRate,
		boostedPeople: boostedPeople,
	}
}

// SimulateTransactions simulates transactions between people at each time step
func (n *InflationaryNetwork) SimulateTransactions() bool {
	if n.TimeStep >= maxTimeSteps {
		return false
	}

	for i := 0; i < n.gridSize; i++ {
		for j := 0; j < n.gridSize; j++ {
			person := n.people[i][j]
			if probPerson, ok := person.(*ProbabilisticPerson); ok {
				probPerson.UpdateSpending(n.people)
			}
		}
	}

	// for people along the diagonal, increase their balance by 10%
	// to simulate inflation.
	if n.TimeStep%10 == 0 {
		for _, p := range n.boostedPeople {
			person := n.people[p.row][p.col]
			balance := person.GetBalance()
			newBalance := balance * n.inflationRate
			person.SetBalance(newBalance)
		}
	}

	n.TimeStep++

	return true
}

// balanceToBlockChar maps a balance value to a Unicode block character
func balanceToBlockChar(normalizedBalance float64) string {
	blockChar := "█"

	// Ensure normalizedBalance is within the [0, 1] range
	if normalizedBalance < 0 {
		normalizedBalance = 0
	} else if normalizedBalance > 1 {
		normalizedBalance = 1
	}

	// Color the block character based on the normalized balance
	var coloredBlockChar *color.Color
	if normalizedBalance < 0.2 {
		coloredBlockChar = color.New(color.FgBlue)
	} else if normalizedBalance < 0.4 {
		coloredBlockChar = color.New(color.FgCyan)
	} else if normalizedBalance < 0.6 {
		coloredBlockChar = color.New(color.FgGreen)
	} else if normalizedBalance < 0.8 {
		coloredBlockChar = color.New(color.FgYellow)
	} else {
		coloredBlockChar = color.New(color.FgRed)
	}

	return coloredBlockChar.Sprint(blockChar)
}
