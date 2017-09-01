package main

import (
	"fmt"
	"math/rand"
	"time"
	"errors"
	"os"
)

var MAX_TIME_BETWEEN_ATTACKS = 60

var ATTACKS_STRATEGY = map[string](func(endpointConfig EndpointConfig, attackConfig AttackConfig, responseChannel chan Response) error){"HTTP_SPAM": RunHTTPSpam,"CORRUPT_HTTP": RunCorruptHTTP,"URL_QUERY_SPAM": RunUrlQuery}

func main() {
	config := GetConfigFromCli()

	if IsCIMode() {
		fmt.Println("🔨 CI Mode detected. Each attack configuration will be ran in sequence for all endpoints.")
		PerformSequentialAttack(config)
	} else {
		wakeTheMonkey(config)
	}
}

func PerformSequentialAttack(config *Config) {
	isFailure := false;

	for _,endpoint := range config.Endpoints {
		for _,attack := range endpoint.Attacks {
			response := executeAttackSync(endpoint, attack)

			isFailure = isFailure && response.Passed
			logResponse(response)
		}
	}

	if isFailure {
		os.Exit(1)
	}

	os.Exit(0)
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
		logResponse(response)
	}
}

func logResponse(response Response) {
	if response.Passed {
		fmt.Printf("✅ Attack %s Passed\n", response.AttackConfig.Type)
	} else {
		fmt.Printf("❌ Attack %s Failed: %s\n", response.AttackConfig.Type, response.Report)
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
		executeAttack(endpoint, attack, responseChannel)
		pauseForRandomDuration()
	}
}

func executeAttack(endpoint EndpointConfig, attack AttackConfig, responseChannel chan Response) {
	attackFunc := getAttackFunction(attack)
	go attackFunc(endpoint, attack, responseChannel)
}

func executeAttackSync(endpoint EndpointConfig, attack AttackConfig) Response {
	attackFunc := getAttackFunction(attack)
	responseChannel := make(chan Response)
	go attackFunc(endpoint, attack, responseChannel)
	response := <- responseChannel

	return response
}

func getAttackFunction(attack AttackConfig) (func(endpointConfig EndpointConfig, attackConfig AttackConfig, responseChannel chan Response) error) {
	attackFunc, present := ATTACKS_STRATEGY[attack.Type]

	if !present {
		panic(errors.New(fmt.Sprintf("Unknown attack type %s", attack.Type)))
	}

	return attackFunc
}

func pauseForRandomDuration() {
	time.Sleep(time.Duration(rand.Intn(MAX_TIME_BETWEEN_ATTACKS)) * time.Second)
}
