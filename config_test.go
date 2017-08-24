package main

import (
  "testing"
)

func TestGetConfig(t *testing.T) {
  t.Run("Test ability to load in correct config", func(t *testing.T) {
    configPath := "examples/config.json"
    config := GetConfig(configPath)

    if config == nil {
      t.Error("Expected Config to be parsed. Got nil")
    }

    if len(config.Endpoints) == 0 {
      t.Error("Expected Config to contain endpoints. Endpoints length = 0")
    }

    endpoint := config.Endpoints[0]

    if endpoint.Name != "QA List orders Endpoint" {
      t.Errorf("Incorrect name for first endpoint. Expected QA List Orders Endpoint. Got %s", endpoint.Name)
    }

    attack := endpoint.Attacks[0]

    if attack.Type != "CORRUPT_HTTP" {
      t.Errorf("Expected attack to be CORRUPT_HTTP. Got %s", attack.Type)
    }
  })
}

func TestIsValidConfig(t *testing.T) {
  t.Run("Test ability to check for config errors", func(t *testing.T) {
    config := &Config{}
    valid, _ := IsValidConfig(config)

    errorIfValid(valid, "Config with no endpoints should not be valid", t)

    config.Endpoints = append([]EndpointConfig{}, EndpointConfig{})
    valid, _ = IsValidConfig(config)

    errorIfValid(valid, "Config endpoint with no name should not be valid", t)

    config.Endpoints[0].Name = "Chris"
    valid, _ = IsValidConfig(config)

    errorIfValid(valid, "Config endpoint with no host should not be valid", t)

    config.Endpoints[0].Host = "host"

    valid, _ = IsValidConfig(config)

    errorIfValid(valid, "Config endpoint with no attacks should not be valid", t)

    config.Endpoints[0].Attacks = append([]AttackConfig{}, AttackConfig{})

    valid, _ = IsValidConfig(config)

    errorIfValid(valid, "Config endpoint attack with no type should not be valid", t)

    config.Endpoints[0].Attacks[0].Type = "Chris"

    valid, _ = IsValidConfig(config)

    if valid != true {
      t.Error("Valid config provided but validator returned false")
    }
  })
}

func errorIfValid(valid bool, message string, t *testing.T ) {
  if valid {
    t.Error(message)
  }
}
