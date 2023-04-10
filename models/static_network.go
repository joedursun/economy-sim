package models

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// StaticNetwork struct representing the network of people
type StaticNetwork struct {
	opts     SimulatorOptions
	gridSize int
	people   [][]Person
	TimeStep int
}

// NewStaticNetwork creates a new network with gridSize x gridSize people
func NewStaticNetwork(opts SimulatorOptions) StaticNetwork {
	network := StaticNetwork{opts: opts}
	rand.Seed(time.Now().UnixNano())

	network.people = NewNormalPopulation(opts.GridSize, opts.BurnFee)

	return network
}

func (n *StaticNetwork) MoneyInCirculation() float64 {
	total := 0.0
	for i := 0; i < len(n.people); i++ {
		for j := 0; j < len(n.people[i]); j++ {
			total += n.people[i][j].GetBalance()
		}
	}
	return total
}

// SimulateTransactions simulates transactions between people at each time step
func (n *StaticNetwork) SimulateTransactions() bool {
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

	n.TimeStep++

	return true
}

// PrintState prints the current state of the network in a gridSize x gridSize grid
func (n *StaticNetwork) PrintState() {
	normalizedBalances := n.normalizeBalances()

	for i := 0; i < n.gridSize; i++ {
		for j := 0; j < n.gridSize; j++ {
			normalizedBalance := normalizedBalances[i][j]
			blockChar := balanceToBlockChar(normalizedBalance)
			fmt.Printf("%s", blockChar)
		}
		fmt.Println()
	}
}

func (n *StaticNetwork) normalizeBalances() [][]float64 {
	minBalance := math.MaxFloat64
	maxBalance := -math.MaxFloat64

	for i := 0; i < n.opts.GridSize; i++ {
		for j := 0; j < n.opts.GridSize; j++ {
			balance := n.people[i][j].GetBalance()
			if balance < minBalance {
				minBalance = balance
			}
			if balance > maxBalance {
				maxBalance = balance
			}
		}
	}

	normalized := make([][]float64, n.opts.GridSize)
	for i := 0; i < n.opts.GridSize; i++ {
		row := make([]float64, n.opts.GridSize)
		for j := 0; j < n.opts.GridSize; j++ {
			balance := n.people[i][j].GetBalance()
			normalizedValue := (balance - minBalance) / (maxBalance - minBalance)
			row[j] = normalizedValue
		}
		normalized[i] = row
	}

	return normalized
}

func (n *StaticNetwork) People() [][]Person {
	return n.people
}

func (n *StaticNetwork) PrintBalanceHistogram(people [][]Person, numBins int) {
	balances := []float64{}

	// Collect all balances
	for i := 0; i < len(people); i++ {
		for j := 0; j < len(people[i]); j++ {
			balances = append(balances, people[i][j].GetBalance())
		}
	}

	// Sort the balances
	sort.Float64s(balances)

	// Find the maximum balance
	maxBalance := balances[len(balances)-1]

	// Calculate bin width
	binWidth := maxBalance / float64(numBins)

	// Create bins and count balances in each bin
	binCounts := make([]int, numBins)
	for _, balance := range balances {
		binIndex := int(math.Floor(balance / binWidth))
		if binIndex >= numBins {
			binIndex = numBins - 1
		}
		binCounts[binIndex]++
	}

	// Find the maximum count to scale the histogram
	maxCount := 0
	for _, count := range binCounts {
		if count > maxCount {
			maxCount = count
		}
	}

	// Print the histogram
	for binIndex, count := range binCounts {
		barLength := int(math.Round(float64(count) / float64(maxCount) * 50))
		bar := strings.Repeat("█", barLength)
		fmt.Printf("%6.2f - %6.2f: %s\n", binWidth*float64(binIndex), binWidth*float64(binIndex+1), bar)
	}
}
