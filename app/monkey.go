package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		randomAttack := pickRandomAttack()
		randomAttack(endpoint, responseChannel)
	}
}

func pickRandomAttack() (func(endpointConfig EndpointConfig, responseChannel chan Response) error) {
	diceRoll := rand.Intn(100)

	if diceRoll <= 100 {
		return RunHttpSpam
	}

	return RunHttpSpam
}
