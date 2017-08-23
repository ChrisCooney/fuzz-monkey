package main

import (
  "testing"
)

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
