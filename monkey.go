package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

// Maximum time in seconds between two attacks when running in the background.
var MaxTimeBetweenAttacks = 60

// A map of attack names to attacks for use as a strategy pattern.
var AttacksStrategy = map[string](func(endpointConfig EndpointConfig, attackConfig AttackConfig, responseChannel chan Response) error){"HTTP_SPAM": RunHTTPSpam,"CORRUPT_HTTP": RunCorruptHTTP,"URL_QUERY_SPAM": RunURLQuery}

func main() {
	config := GetConfigFromCli()

	if IsCIMode() {
		fmt.Println("🔨 CI Mode detected. Each attack configuration will be ran in sequence for all endpoints.")
		PerformSequentialAttack(config)
	} else {
		wakeTheMonkey(config)
	}
}

// PerformSequentialAttack runs through each attack in config and run them in sequence.
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

// SetupTargets initialises the threads for each of the attacks.
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
	attackFunc, present := AttacksStrategy[attack.Type]

	if !present {
		panic(fmt.Errorf("Unknown attack type %s", attack.Type))
	}

	return attackFunc
}

func pauseForRandomDuration() {
	time.Sleep(time.Duration(rand.Intn(MaxTimeBetweenAttacks)) * time.Second)
}
