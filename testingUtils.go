package main

func CreateTestConfiguration(expectedStatus string) (EndpointConfig, AttackConfig) {
  endpoint := EndpointConfig{Name:"Test Endpoint", Protocol:"http",Host:"localhost",Port:"8080",Path:"/my-endpoint"}
  attack := AttackConfig{Type:"HTTP_SPAM",Concurrents:1,MessagesPerConcurrent:1,ExpectedStatus:expectedStatus,Method:"GET"}

  return endpoint, attack
}

func CreateFullTestConfiguration() Config {
  config := Config{}

  endpoint, attack := CreateTestConfiguration("200")

  config.Endpoints = append(config.Endpoints, endpoint)
  config.Endpoints[0].Attacks = append(config.Endpoints[0].Attacks, attack)

  return config
}
