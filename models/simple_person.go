package models

// Example implementation of the Person interface
type SimplePerson struct {
	id               int
	balance          float64
	income           float64
	fixedExpenses    float64
	variableExpenses float64
}

// Implement the Person interface
func (p *SimplePerson) GetID() int {
	return p.id
}

func (p *SimplePerson) GetBalance() float64 {
	return p.balance
}

func (p *SimplePerson) SetBalance(balance float64) {
	p.balance = balance
}

func (p *SimplePerson) GetIncome() float64 {
	return p.income
}

func (p *SimplePerson) SetIncome(income float64) {
	p.income = income
}

func (p *SimplePerson) GetFixedExpenses() float64 {
	return p.fixedExpenses
}

func (p *SimplePerson) SetFixedExpenses(expenses float64) {
	p.fixedExpenses = expenses
}

func (p *SimplePerson) GetVariableExpenses() float64 {
	return p.variableExpenses
}

func (p *SimplePerson) SetVariableExpenses(expenses float64) {
	p.variableExpenses = expenses
}

func (p *SimplePerson) UpdateIncome() {
	// Update income logic here
}

func (p *SimplePerson) UpdateExpenses() {
	// Update expenses logic here
}

func (p *SimplePerson) ReceiveMoney(amount float64) {
	p.balance += amount
}

func (p *SimplePerson) SendMoney(amount float64, receiver Person) bool {
	if amount <= 0 || p.balance < amount {
		return false
	}

	if receiver.ReceiveMoneyFrom(amount, p) {
		p.balance -= amount
		return true
	}

	return false
}

func (p *SimplePerson) ReceiveMoneyFrom(amount float64, sender Person) bool {
	if amount <= 0 {
		return false
	}

	p.balance += amount
	return true
}

func NewSimplePerson() Person {
	person := SimplePerson{id: 1}
	person.SetBalance(1000)
	person.SetIncome(2000)
	person.SetFixedExpenses(500)
	person.SetVariableExpenses(300)

	return &person
}
