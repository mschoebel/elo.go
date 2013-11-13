package main

import (
	"math"
	"testing"
)

var epsilon float64 = 0.00001

func TestExpectedScore(t *testing.T) {

	minScore := 1.0 / (1.0 + math.Pow10(1))
	maxScore := 1.0 / (1.0 + math.Pow10(-1))

	testParams := []struct {
		elo1          int
		elo2          int
		expectedScore float64
	}{
		{0, 0, 0.5},
		{1, 1, 0.5},
		{100, 100, 0.5},
		{1000, 1000, 0.5},

		{2806, 2577, 0.78889},

		{0, 400, minScore},
		{0, 500, minScore},
		{0, 1000, minScore},

		{400, 0, maxScore},
		{500, 0, maxScore},
		{1000, 0, maxScore},
	}

	for i, params := range testParams {

		score1 := expectedScore(params.elo1, params.elo2)
		score2 := expectedScore(params.elo2, params.elo1)

		if score1+score2 != 1.0 {
			t.Errorf("%d. sum of scores is %f - expected 1.0", i, score1+score2)
		}
		if math.Abs(score1-params.expectedScore) > epsilon {
			t.Errorf("%d. calculated %f - expected %f", i, score1, params.expectedScore)
		}
	}
}

func TestUpdatedElo(t *testing.T) {

	testParams := []struct {
		elo1            int
		elo2            int
		result          float64
		expectedNewElo1 int
		expectedNewElo2 int
	}{
		{0, 0, 0.5, 0, 0},

		{1000, 1000, 0.0, 995, 1005},
		{1000, 1000, 0.5, 1000, 1000},
		{1000, 1000, 1.0, 1005, 995},

		{2806, 2577, 0.0, 2798, 2585},
		{2806, 2577, 0.5, 2803, 2580},
		{2806, 2577, 1.0, 2808, 2575},
	}

	for i, params := range testParams {

		newElo1 := updatedElo(params.elo1, params.elo2, params.result, 10.0)
		newElo2 := updatedElo(params.elo2, params.elo1, 1.0-params.result, 10.0)

		if newElo1 != params.expectedNewElo1 {
			t.Errorf("%d. new elo1 is %d - expected %d", i, newElo1, params.expectedNewElo1)
		}
		if newElo2 != params.expectedNewElo2 {
			t.Errorf("%d. new elo2 is %d - expected %d", i, newElo2, params.expectedNewElo2)
		}
	}

}
