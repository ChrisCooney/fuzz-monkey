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
  Endpoint string `json:"endpoint"`
  MaxResponseTime int `json:"maxResponseTime"`
  SuccessStatus string `json:"successStatus"`
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
      return false, errors.New(fmt.Sprintf("Endpoint name can not be empty for endpoint #%d", i + 1))
    }

    if endpoint.Endpoint == "" {
      return false, errors.New(fmt.Sprintf("Endpoint can not be null for endpoint with name %s", endpoint.Name))
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
