package models

import "math/rand"

type ProbabilisticPerson struct {
	SimplePerson
	spendProbability   float64
	minimumSpendAmount float64
	balancePercentage  float64
}

func (p *ProbabilisticPerson) UpdateSpending(people []Person) {
	if rand.Float64() < p.spendProbability {
		minSpend := p.minimumSpendAmount
		maxSpend := p.balance * p.balancePercentage
		spendAmount := minSpend + rand.Float64()*(maxSpend-minSpend)

		if spendAmount > p.balance {
			spendAmount = p.balance
		}

		receiverIndex := rand.Intn(len(people))
		receiver := people[receiverIndex]

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
