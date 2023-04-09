package models

import (
	"math/rand"
	"time"
)

// EthNetwork struct representing the network of people
type EthNetwork struct {
	StaticNetwork
}

// NewEthNetwork creates a new network with gridSize x gridSize people
func NewEthNetwork(gridSize int, spendFee float64) EthNetwork {
	rand.Seed(time.Now().UnixNano())
	network := EthNetwork{
		StaticNetwork: StaticNetwork{gridSize: gridSize},
	}

	network.people = NewGeometricPopulation(gridSize, spendFee)

	return network
}

// SimulateTransactions simulates transactions between people at each time step
func (n *EthNetwork) SimulateTransactions() bool {
	if n.TimeStep >= maxTimeSteps {
		return false
	}

	for i := 0; i < n.gridSize; i++ {
		for j := 0; j < n.gridSize; j++ {
			person := n.people[i][j]
			if probPerson, ok := person.(*ProbabilisticPerson); ok {
				probPerson.UpdateSpending(n.people)
				// simulate staking rewards proportional to balance
				probPerson.balance = probPerson.balance * 1.02
			}
		}
	}

	n.TimeStep++

	return true
}
