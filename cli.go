package main

import (
  "gopkg.in/alecthomas/kingpin.v2"
)

var (
  ciMode = kingpin.Flag("ci-mode", "CI Mode").Short('c').Bool()
	configPath = kingpin.Arg("config", "📚 Fuzz Monkey application configuration JSON file").String()
)

func GetConfigFromCli() (*Config) {
  kingpin.Parse()
  return GetConfig(*configPath)
}

func IsCIMode() bool {
  return *ciMode
}
