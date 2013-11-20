package main

import (
	"fmt"
)

func calculateElo() (rating map[string]int) {

	rating = make(map[string]int)
	for shortName := range config.Participants {
		rating[shortName] = config.InitialRating
	}

	nrOfGamesByPlayer := make(map[string]int)

	for _, result := range config.Results {

		fmt.Printf("%s vs %s (%.2f) -- ", result.Player1, result.Player2, result.Score)

		kFactor := func(player string) (k float64) {
			k = config.K_factor.Start
			if nrOfGamesByPlayer[player] >= config.K_factor.AmateurNrGames {
				k = config.K_factor.Amateur
			}
			if rating[player] >= config.K_factor.MasterElo {
				k = config.K_factor.Master
			}
			fmt.Printf("%s %.0f ", player, k)
			return k
		}

		kPlayer1 := kFactor(result.Player1)
		kPlayer2 := kFactor(result.Player2)

		newPlayer1Elo := updatedElo(
			rating[result.Player1], rating[result.Player2],
			result.Score, kPlayer1)
		newPlayer2Elo := updatedElo(
			rating[result.Player2], rating[result.Player1],
			1.0-result.Score, kPlayer2)

		fmt.Printf("-- %s %5d %s %5d\n", result.Player1, newPlayer1Elo, result.Player2, newPlayer2Elo)

		rating[result.Player1] = newPlayer1Elo
		rating[result.Player2] = newPlayer2Elo

		nrOfGamesByPlayer[result.Player1] += 1
		nrOfGamesByPlayer[result.Player2] += 1
	}

	return rating
}

func calculateElo_alternative(stats resultStats) (rating map[string]int) {

	rating = make(map[string]int)
	for shortName := range config.Participants {
		rating[shortName] = config.InitialRating
	}

	for p, pStats := range stats.statsByPair {

		if pStats.nrOfGames < 2 {
			fmt.Printf("%2v   at least 2 games required\n", p)
			continue
		}

		avgResult := pStats.sumScores / float64(pStats.nrOfGames)

		deltaPlayer1 := updatedElo(config.InitialRating, config.InitialRating,
			avgResult, config.K_factor.Start) - config.InitialRating
		deltaPlayer2 := updatedElo(config.InitialRating, config.InitialRating,
			1.0-avgResult, config.K_factor.Start) - config.InitialRating

		rating[p.player1] += deltaPlayer1
		rating[p.player2] += deltaPlayer2

		fmt.Printf("%2v   %.2f d1 %5d d2 %5d\n", p, avgResult, deltaPlayer1, deltaPlayer2)
	}

	return rating
}

func main() {

	initializeConfiguration()

	stats := calculateResultStatistics()

	eloRating := calculateElo()
	eloRating_alternative := calculateElo_alternative(stats)

	// print statistics
	fmt.Printf("\nNumber of games by player (%d games overall)\n", stats.nrOfGames)
	for i, shortName := range sortedKeys(stats.nrOfGamesByPlayer) {
		fmt.Printf("%2d. %-20s (%d)\n", i+1, config.Participants[shortName], stats.nrOfGamesByPlayer[shortName])
	}

	// print pair statistics
	fmt.Println("\nScores by pair")
	for p, pStats := range stats.statsByPair {
		fmt.Printf("%2v   games %3d  - avg. score %.2f\n", p, pStats.nrOfGames, pStats.sumScores/float64(pStats.nrOfGames))
	}

	// print result ranking
	fmt.Println("\nELO ranking")
	for i, shortName := range sortedKeys(eloRating) {
		sign := " "
		if stats.nrOfGamesByPlayer[shortName] < config.MinNrOfGames {
			sign = "!"
		}
		fmt.Printf("%s%2d. %-20s (%d)\n", sign, i+1, config.Participants[shortName], eloRating[shortName])
	}

	fmt.Println("\nELO ranking alternative")
	for i, shortName := range sortedKeys(eloRating_alternative) {
		sign := " "
		if stats.nrOfGamesByPlayer[shortName] < config.MinNrOfGames {
			sign = "!"
		}
		fmt.Printf("%s%2d. %-20s (%d)\n", sign, i+1, config.Participants[shortName], eloRating_alternative[shortName])
	}
}
