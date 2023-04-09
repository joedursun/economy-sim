package models

import "math/rand"

type ProbabilisticPerson struct {
	SimplePerson
	spendProbability   float64
	minimumSpendAmount float64
	balancePercentage  float64
}

var directions = []struct{ x, y int }{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func (p *ProbabilisticPerson) UpdateSpending(grid [][]Person) {
	if rand.Float64() < p.spendProbability {
		minSpend := p.minimumSpendAmount
		maxSpend := p.balance * p.balancePercentage
		spendAmount := minSpend + rand.Float64()*(maxSpend-minSpend)

		if spendAmount > p.balance {
			spendAmount = p.balance
		}

		row, col := -1, -1
		for i := 0; i < len(grid) && row == -1; i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j].GetID() == p.GetID() {
					row, col = i, j
					break
				}
			}
		}

		adjacent := make([]Person, 0, 8)

		for _, d := range directions {
			newRow := row + d.x
			newCol := col + d.y

			if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) {
				adjacent = append(adjacent, grid[newRow][newCol])
			}
		}

		receiverIndex := rand.Intn(len(adjacent))
		receiver := adjacent[receiverIndex]

		if p.GetID() != receiver.GetID() {
			p.SendMoney(spendAmount, receiver)
		}
	}
}

func NewProbabilisticPerson(id int, spendProbability, minimumSpendAmount, balancePercentage float64) *ProbabilisticPerson {
	person := &ProbabilisticPerson{
		SimplePerson:       SimplePerson{id: id},
		spendProbability:   spendProbability,
		minimumSpendAmount: minimumSpendAmount,
		balancePercentage:  balancePercentage,
	}
	return person
}
