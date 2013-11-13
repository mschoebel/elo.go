package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type resultInfo struct {
	Player1 string
	Player2 string
	Score   float64
}

type tournamentConfig struct {
	Participants map[string]string
	Results      []resultInfo

	InitialRating int
	K_factor      float64
	MinNrOfGames  int
}

var config tournamentConfig

func initializeConfiguration() {

	var tournamentFilePath string

	flag.StringVar(&tournamentFilePath, "t", "tournament.json", "Tournament description file (JSON)")
	flag.Parse()

	tournamentJSON, err := ioutil.ReadFile(tournamentFilePath)
	if err != nil {
		log.Fatalf("Could not read tournament description file:\n%v", err)
	}

	err = json.Unmarshal(tournamentJSON, &config)
	if err != nil {
		log.Fatalf("Could not parse tournament description:\n%v", err)
	}

	if config.InitialRating == 0 {
		config.InitialRating = 1000
	}
	if config.K_factor == 0.0 {
		config.K_factor = 30.0
	}
}
