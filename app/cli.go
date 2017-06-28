package main

import (
  "gopkg.in/alecthomas/kingpin.v2"
)

var (
	configPath = kingpin.Arg("config", "Fuzz Monkey application configuration JSON file").Required().String()
)

func GetConfigFromCli() (*Config) {
  kingpin.Parse()
  return GetConfig(*configPath)
}
