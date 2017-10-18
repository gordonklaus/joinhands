package main

import (
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func main() {
	rand.Seed(time.Now().Unix())

	const numPeople = 100
	groupSizes := make([]int, numPeople+1)
	for i := 0; i < 10000; i++ {
		groups := joinHands(numPeople)
		for _, count := range groups {
			groupSizes[count]++
		}
	}

	// for size := range groupSizes {
	// 	groupSizes[size] *= size
	// }

	saveBarChart(groupSizes[:100])
}

func joinHands(numPeople int) []int {
	done := []int{}
	free := make([]int, numPeople)
	for i := range free {
		free[i] = 1
	}

	for len(free) > 0 {
		i := rand.Intn(len(free))
		j := rand.Intn(len(free))
		if i == j {
			done = append(done, free[i])
		} else {
			free[i] += free[j]
		}
		free = append(free[:j], free[j+1:]...)
	}

	return done
}

func saveBarChart(values []int) {
	vals := make(plotter.Values, len(values))
	for i := range vals {
		vals[i] = float64(values[i])
	}

	p, err := plot.New()
	chk(err)
	p.Title.Text = "Number of groups by size"
	p.Y.Label.Text = "Count"
	p.X.Label.Text = "Group size"

	chart, err := plotter.NewBarChart(vals, 1)
	chk(err)

	p.Add(chart)
	chk(p.Save(400, 300, "groupSizes.png"))
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
