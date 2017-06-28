package main

import (
  "bytes"
  "net/http"
  "math/rand"
)

func SendRandomHttpRequest(endpoint string) (*http.Response, error) {
  client := &http.Client{}
  method := getRandomRequestMethod()

  request, err := http.NewRequest(method, endpoint, bytes.NewBufferString("hello"))

  if err != nil {
    return nil, err
  }

  return client.Do(request)
}

func getRandomRequestMethod() string {
  diceRoll := rand.Intn(100)

  if diceRoll < 25 {
    return "GET"
  } else if diceRoll < 50 {
    return "HEAD"
  } else {
    return "OPTIONS"
  }
}
