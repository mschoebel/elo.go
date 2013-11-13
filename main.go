package main

import (
	"fmt"
)

func main() {

	initializeConfiguration()

	// initialize elo rating for all participants
	eloRating := make(map[string]int)
	for shortName := range config.Participants {
		eloRating[shortName] = config.InitialRating
	}

	// apply results
	nrOfGames := make(map[string]int)
	for _, result := range config.Results {
		fmt.Printf("%2s vs. %2s = %2.1f\n", result.Player1, result.Player2, result.Score)

		newPlayer1Elo := updatedElo(
			eloRating[result.Player1], eloRating[result.Player2],
			result.Score, config.K_factor)
		newPlayer2Elo := updatedElo(
			eloRating[result.Player2], eloRating[result.Player1],
			1.0-result.Score, config.K_factor)

		eloRating[result.Player1] = newPlayer1Elo
		eloRating[result.Player2] = newPlayer2Elo

		nrOfGames[result.Player1] += 1
		nrOfGames[result.Player2] += 1
	}

	// check nr of games
	for shortName := range eloRating {
		if nrOfGames[shortName] < config.MinNrOfGames {
			// not enough games
			eloRating[shortName] = 0
		}
	}

	// print result ranking
	fmt.Println()
	for i, shortName := range sortedKeys(eloRating) {
		fmt.Printf("%2d. %-20s (%d)\n", i+1, config.Participants[shortName], eloRating[shortName])
	}
}
