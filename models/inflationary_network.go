package models

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	maxTimeSteps = 500
)

var (
	blockChars = []string{" ", "░", "▒", "▓", "█"}
)

// InflationaryNetwork struct representing the network of people
type InflationaryNetwork struct {
	gridSize int
	people   [][]Person
	TimeStep int
}

// NewInflationaryNetwork creates a new network with gridSize x gridSize people
func NewInflationaryNetwork(gridSize int) *InflationaryNetwork {
	network := &InflationaryNetwork{gridSize: gridSize}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < gridSize; i++ {
		row := make([]Person, gridSize)
		for j := 0; j < gridSize; j++ {
			id := i*gridSize + j + 1
			spendProbability := 0.5
			minimumSpendAmount := 20.0
			balancePercentage := 0.1

			person := NewProbabilisticPerson(id, spendProbability, minimumSpendAmount, balancePercentage)
			person.SetBalance(rand.Float64() * 1000)
			person.SetIncome(rand.Float64() * 2000)
			person.SetFixedExpenses(rand.Float64() * 500)
			person.SetVariableExpenses(rand.Float64() * 300)
			row[j] = person
		}
		network.people = append(network.people, row)
	}

	return network
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

	if n.TimeStep%10 == 0 {
		for i := 0; i < n.gridSize; i++ {
			for j := 0; j < n.gridSize; j++ {
				if i == j || i == n.gridSize-1-j {
					person := n.people[i][j]
					balance := person.GetBalance()
					newBalance := balance * 1.1
					person.SetBalance(newBalance)
				}
			}
		}
	}

	n.TimeStep++

	return true
}

// PrintState prints the current state of the network in a gridSize x gridSize grid
func (n *InflationaryNetwork) PrintState() {
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

func (n *InflationaryNetwork) normalizeBalances() [][]float64 {
	minBalance := math.MaxFloat64
	maxBalance := -math.MaxFloat64

	for i := 0; i < n.gridSize; i++ {
		for j := 0; j < n.gridSize; j++ {
			balance := n.people[i][j].GetBalance()
			if balance < minBalance {
				minBalance = balance
			}
			if balance > maxBalance {
				maxBalance = balance
			}
		}
	}

	normalized := make([][]float64, n.gridSize)
	for i := 0; i < n.gridSize; i++ {
		row := make([]float64, n.gridSize)
		for j := 0; j < n.gridSize; j++ {
			balance := n.people[i][j].GetBalance()
			normalizedValue := (balance - minBalance) / (maxBalance - minBalance)
			row[j] = normalizedValue
		}
		normalized[i] = row
	}

	return normalized
}

// balanceToBlockChar maps a balance value to a Unicode block character
func balanceToBlockChar(normalizedBalance float64) string {
	// Map the normalized balance value to an index in the blockChars slice.
	index := int(normalizedBalance * float64(len(blockChars)-1))

	return blockChars[index]
}
