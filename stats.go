package main

type pairKey struct {
	player1 string
	player2 string
}

type pairStats struct {
	nrOfGames int
	sumScores float64
}

type resultStats struct {
	nrOfGames         int
	nrOfGamesByPlayer map[string]int

	statsByPair map[pairKey]pairStats
}

func pairingFromResult(result resultInfo) (p pairKey, invert bool) {

	if result.Player1 < result.Player2 {
		p.player1 = result.Player1
		p.player2 = result.Player2
		invert = false
	} else {
		p.player1 = result.Player2
		p.player2 = result.Player1
		invert = true
	}

	return p, invert
}

func calculateResultStatistics() (stats resultStats) {

	stats.nrOfGamesByPlayer = make(map[string]int)
	stats.statsByPair = make(map[pairKey]pairStats)

	stats.nrOfGames = len(config.Results)

	for _, result := range config.Results {

		p, invert := pairingFromResult(result)
		pStats := stats.statsByPair[p]

		pStats.nrOfGames += 1
		if invert {
			pStats.sumScores += 1.0 - result.Score
		} else {
			pStats.sumScores += result.Score
		}

		stats.statsByPair[p] = pStats

		stats.nrOfGamesByPlayer[result.Player1] += 1
		stats.nrOfGamesByPlayer[result.Player2] += 1
	}

	return stats
}
