package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_TIME_BETWEEN_ATTACKS = 60

func main() {
	config := GetConfigFromCli()
	wakeTheMonkey(config)
}

func wakeTheMonkey(config *Config) {
	responseChannel := make(chan Response)
	setupTargets(config, responseChannel)
	listenForResponses(responseChannel)
}

func listenForResponses(responseChannel chan Response) {
	for {
		response := <- responseChannel
		fmt.Printf("Response found. Passed = %v Report = %s\n", response.Passed, response.Report)
	}
}

func setupTargets(config *Config, responseChannel chan Response) {
	for _,endpoint := range config.Endpoints {
		fmt.Printf("Setting up %s\n", endpoint.Name)
		go beginHarassment(endpoint, responseChannel)
	}
}

func beginHarassment(endpoint EndpointConfig, responseChannel chan Response) {
	for {
		randomAttack := pickRandomAttack()
		go randomAttack(endpoint, responseChannel)
		pauseForRandomDuration()
	}
}

func pickRandomAttack() (func(endpointConfig EndpointConfig, responseChannel chan Response) error) {
	diceRoll := rand.Intn(100)

	if diceRoll <= 50 {
		return RunHttpSpam
	} else {
		return RunCorruptHttp
	}
}

func pauseForRandomDuration() {
	time.Sleep(time.Duration(rand.Intn(MAX_TIME_BETWEEN_ATTACKS)) * time.Second)
}
