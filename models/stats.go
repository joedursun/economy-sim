package models

import (
	"math"
	"math/rand"
	"time"
)

func geometricDistribution2DSlice(rows, cols int, p float64, maxBalance float64) [][]float64 {
	slice := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		row := make([]float64, cols)
		for j := 0; j < cols; j++ {
			k := i*cols + j + 1 // Add 1 since the geometric distribution starts at 1
			value := math.Pow(1-p, float64(k))
			row[j] = maxBalance * value
		}
		slice[i] = row
	}

	return shuffle2DSlice(slice)
}

func shuffle2DSlice(slice [][]float64) [][]float64 {
	rows := len(slice)
	cols := len(slice[0])

	// Flatten the 2D slice into a 1D slice
	flatSlice := make([]float64, rows*cols)
	idx := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			flatSlice[idx] = slice[i][j]
			idx++
		}
	}

	// Shuffle the flattened slice
	rand.Shuffle(len(flatSlice), func(i, j int) {
		flatSlice[i], flatSlice[j] = flatSlice[j], flatSlice[i]
	})

	// Convert the shuffled flattened slice back to a 2D slice
	shuffledSlice := make([][]float64, rows)
	idx = 0
	for i := 0; i < rows; i++ {
		row := make([]float64, cols)
		for j := 0; j < cols; j++ {
			row[j] = flatSlice[idx]
			idx++
		}
		shuffledSlice[i] = row
	}

	return shuffledSlice
}

func normalDistribution2DSlice(rows, cols int, mean, stddev float64) [][]float64 {
	slice := make([][]float64, rows)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < rows; i++ {
		row := make([]float64, cols)
		for j := 0; j < cols; j++ {
			row[j] = rand.NormFloat64()*stddev + mean
		}
		slice[i] = row
	}

	return slice
}

func NewNormalPopulation(gridSize int, spendFee float64) [][]Person {
	people := make([][]Person, gridSize)
	balances := normalDistribution2DSlice(gridSize, gridSize, 1000.0, 100.0)

	for i := 0; i < gridSize; i++ {
		row := make([]Person, gridSize)
		for j := 0; j < gridSize; j++ {
			startingBalance := balances[i][j]
			income := startingBalance * 0.1
			id := i*gridSize + j + 1
			minSpend := 10.0
			maxSpendPercentage := 0.5
			spendProbability := 0.5
			person := NewProbabilisticPerson(id, spendProbability, minSpend, maxSpendPercentage, spendFee)
			person.SetBalance(startingBalance)
			person.SetIncome(income)
			row[j] = person
		}
		people[i] = row
	}

	return people
}

func NewGeometricPopulation(gridSize int, spendFee float64) [][]Person {
	people := make([][]Person, gridSize)

	p := 0.1 // Probability parameter for the geometric distribution
	maxBalance := 1000.0
	geometricBalances := geometricDistribution2DSlice(gridSize, gridSize, p, maxBalance)

	for i := 0; i < gridSize; i++ {
		row := make([]Person, gridSize)
		for j := 0; j < gridSize; j++ {
			startingBalance := geometricBalances[i][j]
			income := startingBalance * 0.1
			id := i*gridSize + j + 1
			minSpend := 10.0
			maxSpendPercentage := 0.5
			spendProbability := 0.5
			person := NewProbabilisticPerson(id, spendProbability, minSpend, maxSpendPercentage, spendFee)
			person.SetBalance(startingBalance)
			person.SetIncome(income)
			row[j] = person
		}
		people[i] = row
	}

	return people
}
