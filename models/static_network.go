package models

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// StaticNetwork struct representing the network of people
type StaticNetwork struct {
	gridSize int
	people   [][]Person
	TimeStep int
}

// NewStaticNetwork creates a new network with gridSize x gridSize people
func NewStaticNetwork(gridSize int) StaticNetwork {
	network := StaticNetwork{gridSize: gridSize}
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
func (n *StaticNetwork) SimulateTransactions() bool {
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
