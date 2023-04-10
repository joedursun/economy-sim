package models

type SimulatorOptions struct {
	GridSize          int
	NumCentralBankers int
	// BurnFee is the percentage of money that is burned when a transaction occurs. Similar to how ETH burns gas fees.
	BurnFee            float64
	MinSpend           float64
	MaxSpendPercentage float64
	SpendProbability   float64
	InflationRate      float64
}
