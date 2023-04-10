package models

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func getBalances(people [][]Person) []float64 {
	balances := make([]float64, 0, len(people)*len(people[0]))

	for i := 0; i < len(people); i++ {
		for j := 0; j < len(people[i]); j++ {
			balances = append(balances, people[i][j].GetBalance())
		}
	}
	return balances
}

func (n *StaticNetwork) SaveHistogram(bins int, filename string) error {
	// Create a new plot
	p := plot.New()

	p.Title.Text = "Balance Histogram"
	p.X.Label.Text = "Balance"
	p.Y.Label.Text = "Frequency"

	balances := getBalances(n.people)
	// Create the histogram of the data
	hist, err := plotter.NewHist(plotter.Values(balances), bins)
	if err != nil {
		return err
	}

	hist.Normalize(1) // Normalize the histogram

	// Set the histogram fill color
	hist.FillColor = color.RGBA{R: 196, G: 196, B: 196, A: 255}
	hist.LineStyle.Width = vg.Length(0) // Remove the outline

	// Add the histogram to the plot
	p.Add(hist)

	// Save the plot to a PNG file
	img := vgimg.New(6*vg.Inch, 4*vg.Inch)
	dc := draw.New(img)
	dc = draw.Crop(dc, 0, -0.5*vg.Centimeter, 0, -0.5*vg.Centimeter) // Add some padding
	p.Draw(dc)
	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = vgimg.PngCanvas{Canvas: img}.WriteTo(w)
	if err != nil {
		return err
	}

	fmt.Printf("Histogram saved to %s\n", filename)
	return nil
}
