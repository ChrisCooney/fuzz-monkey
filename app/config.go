package main

import (
  "encoding/json"
  "io/ioutil"
  "errors"
  "fmt"
)

type Config struct {
  Endpoints []EndpointConfig `json:"endpoints"`
}

type EndpointConfig struct {
  Name string `json:"name"`
  Host string `json:"host"`
  Port string `json:"port"`
  Path string `json:"path"`
  Protocol string `json:"protocol"`
  Attacks []AttackConfig `json:"attacks"`
}

type AttackConfig struct {
  Type string `json:"type"`
  MaxResponseTime int `json:"maxResponseTime"`
  ExpectedStatus string `json:"expectedStatus"`
  Concurrents int `json:"concurrent"`
  MessagesPerConcurrent int `json:"messagesPerConcurrent"`
}

func GetConfig(configPath string) (*Config) {
  fileContents := loadConfigFile(configPath)
  return mapFileToObject(fileContents)
}

func loadConfigFile(configPath string) ([]byte) {
  file, err := ioutil.ReadFile(configPath)

  CheckError(err)

  return file
}

func isValidConfig(config *Config) (bool, error) {
  for i,endpoint := range config.Endpoints {
    if endpoint.Name == "" {
      return false, errors.New(fmt.Sprintf("⚠️ Endpoint name can not be empty for endpoint #%d ⚠️", i + 1))
    }

    if endpoint.Host == "" {
      return false, errors.New(fmt.Sprintf("⚠️ Host can not be null for endpoint with name %s ⚠️", endpoint.Name))
    }
  }

  return true, nil
}

func mapFileToObject(contents []byte) (*Config) {
  config := &Config{}
  err := json.Unmarshal(contents, config)
  CheckError(err)

  valid, err := isValidConfig(config)

  if !valid {
    panic(err)
  }

  return config
}
