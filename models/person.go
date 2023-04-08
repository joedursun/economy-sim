package models

// Person interface representing an individual in the network
type Person interface {
	GetID() int
	GetBalance() float64
	SetBalance(balance float64)
	GetIncome() float64
	SetIncome(income float64)
	GetFixedExpenses() float64
	SetFixedExpenses(expenses float64)
	GetVariableExpenses() float64
	SetVariableExpenses(expenses float64)
	UpdateIncome()
	UpdateExpenses()
	ReceiveMoney(amount float64)
	SendMoney(amount float64, receiver Person) bool
	ReceiveMoneyFrom(amount float64, sender Person) bool
}
