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
	// centralBankers are people who are given extra money periodically
	centralBankers []struct {
		row int
		col int
	}
}

// NewInflationaryNetwork creates a new network with gridSize x gridSize people
func NewInflationaryNetwork(opts SimulatorOptions) InflationaryNetwork {
	bankers := make([]struct {
		row int
		col int
	}, opts.NumCentralBankers)
	for i := 0; i < opts.NumCentralBankers; i++ {
		bankers[i] = struct {
			row int
			col int
		}{
			row: rand.Intn(opts.GridSize),
			col: rand.Intn(opts.GridSize),
		}
	}

	return InflationaryNetwork{
		StaticNetwork:  NewStaticNetwork(opts),
		inflationRate:  opts.InflationRate,
		centralBankers: bankers,
	}
}

// SimulateTransactions simulates transactions between people at each time step
func (n *InflationaryNetwork) SimulateTransactions() bool {
	if n.TimeStep >= maxTimeSteps {
		return false
	}

	for i := 0; i < n.opts.GridSize; i++ {
		for j := 0; j < n.opts.GridSize; j++ {
			person := n.people[i][j]
			if probPerson, ok := person.(*ProbabilisticPerson); ok {
				probPerson.UpdateSpending(n.people)
			}
		}
	}

	// for people along the diagonal, increase their balance by 10%
	// to simulate inflation.
	if n.TimeStep%10 == 0 {
		for _, p := range n.centralBankers {
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
