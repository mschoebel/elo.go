package main

import (
	"math"
)

func round(value float64) int {

	if value < 0.0 {
		value -= 0.5
	} else {
		value += 0.5
	}

	return int(value)
}

func expectedScore(elo int, eloOpponent int) float64 {

	eloDifference := eloOpponent - elo

	if eloDifference > 400 {
		eloDifference = 400
	}
	if eloDifference < -400 {
		eloDifference = -400
	}

	return 1.0 / (1.0 + math.Pow(10, float64(eloDifference)/400.0))
}

func updatedElo(elo int, eloOpponent int, score float64, k float64) int {

	return elo + round(k*(score-expectedScore(elo, eloOpponent)))
}
