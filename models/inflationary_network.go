package models

const (
	maxTimeSteps = 500
)

var (
	blockChars = []string{" ", "░", "▒", "▓", "█"}
)

// InflationaryNetwork struct representing the network of people
type InflationaryNetwork struct {
	StaticNetwork
}

// NewInflationaryNetwork creates a new network with gridSize x gridSize people
func NewInflationaryNetwork(gridSize int) InflationaryNetwork {
	return InflationaryNetwork{StaticNetwork: NewStaticNetwork(gridSize)}
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

// balanceToBlockChar maps a balance value to a Unicode block character
func balanceToBlockChar(normalizedBalance float64) string {
	// Map the normalized balance value to an index in the blockChars slice.
	index := int(normalizedBalance * float64(len(blockChars)-1))

	return blockChars[index]
}
