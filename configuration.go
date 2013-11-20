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

type kFactorSpec struct {
	Start float64

	Amateur        float64
	AmateurNrGames int

	Master    float64
	MasterElo int
}

type tournamentConfig struct {
	Participants map[string]string
	Results      []resultInfo

	InitialRating int
	K_factor      kFactorSpec
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

	if config.K_factor.Start == 0.0 {
		config.K_factor.Start = 30.0
	}
	if config.K_factor.Amateur == 0.0 {
		config.K_factor.Amateur = 20.0
	}
	if config.K_factor.AmateurNrGames == 0 {
		config.K_factor.AmateurNrGames = 10
	}
	if config.K_factor.Master == 0.0 {
		config.K_factor.Master = 10.0
	}
	if config.K_factor.MasterElo == 0 {
		config.K_factor.MasterElo = 2000
	}

}
