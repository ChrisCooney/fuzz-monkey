package main

import (
	"fmt"
	"math/rand"
	"time"
	"errors"
)

var MAX_TIME_BETWEEN_ATTACKS = 60

var ATTACKS_STRATEGY = map[string](func(endpointConfig EndpointConfig, attackConfig AttackConfig, responseChannel chan Response) error){"HTTP_SPAM": RunHTTPSpam,"CORRUPT_HTTP": RunCorruptHTTP}

func main() {
	config := GetConfigFromCli()
	wakeTheMonkey(config)
}

func wakeTheMonkey(config *Config) {
	fmt.Println("🐒 Waking the Monkey")
	responseChannel := make(chan Response)
	SetupTargets(config, responseChannel)
	listenForResponses(responseChannel)
}

func listenForResponses(responseChannel chan Response) {
	for {
		response := <- responseChannel

		if response.Passed {
			fmt.Printf("✅ Attack %s Passed\n", response.AttackConfig.Type)
		} else {
			fmt.Printf("❌ Attack %s Failed\n", response.AttackConfig.Type)
			fmt.Printf("❌ Reason: %s\n", response.Report)
		}
	}
}

// Sets up the targets in the config file for attack.
func SetupTargets(config *Config, responseChannel chan Response) {
	for _,endpoint := range config.Endpoints {
		fmt.Printf("🎯 Setting up %s\n", endpoint.Name)
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
		attackFunc, present := ATTACKS_STRATEGY[attack.Type]

		if !present {
			panic(errors.New(fmt.Sprintf("Unknown attack type %s", attack.Type)))
		}

		go attackFunc(endpoint, attack, responseChannel)
		pauseForRandomDuration()
	}
}

func pauseForRandomDuration() {
	time.Sleep(time.Duration(rand.Intn(MAX_TIME_BETWEEN_ATTACKS)) * time.Second)
}
