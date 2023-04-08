package models

import (
	"fmt"
	"math/rand"
	"time"
)

// Network struct representing the network of people
type Network struct {
	gridSize int
	people   [][]Person
}

// NewNetwork creates a new network with gridSize x gridSize people
func NewNetwork(gridSize int) *Network {
	network := &Network{gridSize: gridSize}
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
func (n *Network) SimulateTransactions() {
	peopleList := make([]Person, 0, n.gridSize*n.gridSize)

	for i := 0; i < n.gridSize; i++ {
		for j := 0; j < n.gridSize; j++ {
			peopleList = append(peopleList, n.people[i][j])
		}
	}

	for _, person := range peopleList {
		if probPerson, ok := person.(*ProbabilisticPerson); ok {
			probPerson.UpdateSpending(peopleList)
		}
	}
}

// PrintState prints the current state of the network in a gridSize x gridSize grid
func (n *Network) PrintState() {
	for i := 0; i < n.gridSize; i++ {
		for j := 0; j < n.gridSize; j++ {
			balance := n.people[i][j].GetBalance()
			fmt.Printf("%s", balanceToBlockChar(balance))
		}
		fmt.Println()
	}
}

// balanceToBlockChar maps a balance value to a Unicode block character
func balanceToBlockChar(balance float64) string {
	switch {
	case balance <= 200:
		return "░" // Light shade
	case balance <= 400:
		return "▒" // Medium shade
	case balance <= 600:
		return "▓" // Dark shade
	default:
		return "█" // Full block
	}
}
