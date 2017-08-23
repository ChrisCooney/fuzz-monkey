package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_TIME_BETWEEN_ATTACKS = 60

var ATTACKS_STRATEGY = map[string](func(endpointConfig EndpointConfig, attackConfig AttackConfig, responseChannel chan Response) error){"HTTP_SPAM": RunHttpSpam,"CORRUPT_HTTP": RunCorruptHttp,"RANDOM_RABBIT_JSON": RunRandomRabbitJson}

func main() {
	config := GetConfigFromCli()
	wakeTheMonkey(config)
}

func wakeTheMonkey(config *Config) {
	fmt.Println("🐒 Waking the Monkey 🐒")
	responseChannel := make(chan Response)
	setupTargets(config, responseChannel)
	listenForResponses(responseChannel)
}

func listenForResponses(responseChannel chan Response) {
	for {
		response := <- responseChannel
		fmt.Printf("Response found. Passed = %v | Report = %s | Attack = %s\n", response.Passed, response.Report, response.AttackConfig.Type)
	}
}

func setupTargets(config *Config, responseChannel chan Response) {
	for _,endpoint := range config.Endpoints {
		fmt.Printf("🎯 Setting up %s 🎯\n", endpoint.Name)
		setupAttackThreads(endpoint, responseChannel)
	}
}

func setupAttackThreads(endpoint EndpointConfig, responseChannel chan Response) {
	for _,attack := range endpoint.Attacks {
		go beginHarassment(endpoint, attack, responseChannel)
	}
}

func beginHarassment(endpoint EndpointConfig, attack AttackConfig, responseChannel chan Response) {
	for {
		go ATTACKS_STRATEGY[attack.Type](endpoint, attack, responseChannel)
		pauseForRandomDuration()
	}
}

func pauseForRandomDuration() {
	time.Sleep(time.Duration(rand.Intn(MAX_TIME_BETWEEN_ATTACKS)) * time.Second)
}
