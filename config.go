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
  Concurrents int `json:"concurrents"`
  MessagesPerConcurrent int `json:"messagesPerConcurrent"`
  Method string `json:"method"`
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

func IsValidConfig(config *Config) (bool, error) {

  if len(config.Endpoints) == 0 {
    return false, errors.New(fmt.Sprintf("⚠️ Endpoints can not be empty. The monkey needs victims. ⚠️"))
  }

  for i,endpoint := range config.Endpoints {
    if endpoint.Name == "" {
      return false, errors.New(fmt.Sprintf("⚠️ Endpoint name can not be empty for endpoint #%d. The monkey is like Arya Stark. It needs a name. ⚠️", i + 1))
    }

    if endpoint.Host == "" {
      return false, errors.New(fmt.Sprintf("⚠️ Host can not be null for endpoint with name %s. The monkey needs an address to go after. ⚠️", endpoint.Name))
    }

    if len(endpoint.Attacks) == 0 {
      return false, errors.New(fmt.Sprintf("⚠️ Endpoint must have attacks associated with it. The monkey kills all it sees. ⚠️"))
    }

    for j,attack := range endpoint.Attacks {
      if attack.Type == "" {
        return false, errors.New(fmt.Sprintf("⚠️ Attack config #%d for endpoint %s needs a type. Future versions will interpret this as an all access pass for the monkey. ⚠️", endpoint.Name, j + 1))
      }
    }
  }

  return true, nil
}

func mapFileToObject(contents []byte) (*Config) {
  config := &Config{}
  err := json.Unmarshal(contents, config)
  CheckError(err)

  valid, err := IsValidConfig(config)

  if !valid {
    panic(err)
  }

  return config
}
