package cli

import (
  "gopkg.in/alecthomas/kingpin.v2"
  "fuzz-monkey/app/models"
  "fuzz-monkey/app/utils"
  "encoding/json"
  "io/ioutil"
)

var (
	configPath = kingpin.Arg("config", "Fuzz Monkey application configuration JSON file").Required().String()
)

func GetConfig() (*models.Config) {
  kingpin.Parse()
  fileContents := loadConfigFile()
  return mapFileToObject(fileContents)
}

func loadConfigFile() ([]byte) {
  file, err := ioutil.ReadFile(*configPath)

  utils.CheckError(err)

  return file
}

func mapFileToObject(contents []byte) (*models.Config) {
  config := &models.Config{}
  err := json.Unmarshal(contents, config)
  utils.CheckError(err)
  return config
}
