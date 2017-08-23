package main

import (
  "fmt"
  "net/http"
  "strings"
)

type Response struct {
  Passed bool
  Report string
  AttackConfig AttackConfig
}

func checkHttpResponse(httpResponse *http.Response, config AttackConfig) (bool, string) {
  if strings.Trim(httpResponse.Status, " ") != strings.Trim(config.ExpectedStatus, " ") {
    reason := fmt.Sprintf("Invalid status code of %s detected. Expected %s", httpResponse.Status, config.ExpectedStatus)
    return false, reason
  }

  /* TODO: Check response time against max response time */

  return true, ""
}

func checkHttpResponses(httpResponses []*http.Response, config AttackConfig) (bool, string) {
  for _,httpResp := range httpResponses {

    if httpResp == nil {
      return false, "Error occurred during HTTP request"
    }

    passed, reason := checkHttpResponse(httpResp, config)
    if !passed {
      return passed, reason
    }
  }

  return true, ""
}

func dispatchMultipleHttpRequests(endpoint string, c chan *http.Response, count int) {
  for i := 0; i < count; i++ {
    SendRandomHttpRequest(endpoint, c)
  }
}

func dispatchConcurrentHttpRequests(concurrentCount int, endpoint string, c chan *http.Response, count int) {
  for i:=0; i < concurrentCount; i++ {
    go dispatchMultipleHttpRequests(endpoint, c, count)
  }
}

func collectConcurrentHttpResponses(responses []*http.Response, c chan *http.Response, expectedCount int) []*http.Response {
  for len(responses) < (expectedCount) {
    responses = readResponseFromChannel(responses, c)

    if responses == nil {
      return nil
    }
  }

  return responses
}

func readResponseFromChannel(responses []*http.Response, c chan *http.Response) []*http.Response {
  response := <- c

  if response == nil {
    return nil
  }

  defer response.Body.Close()
  return append(responses, response)
}

func RunHttpSpam(endpointConfig EndpointConfig, attackConfig AttackConfig, responseChannel chan Response) error {
  fmt.Printf("🔥 Running HTTP Spam against %s 🔥\n", endpointConfig.Name)

  responses := []*http.Response{}
  c := make(chan *http.Response)

  endpoint := BuildNetworkPath(endpointConfig.Protocol, endpointConfig.Host, endpointConfig.Port, endpointConfig.Path)

  messageCount := attackConfig.Concurrents * attackConfig.MessagesPerConcurrent

  dispatchConcurrentHttpRequests(attackConfig.Concurrents, endpoint, c, attackConfig.MessagesPerConcurrent)
  collectConcurrentHttpResponses(responses, c, messageCount)

  if len(responses) == 0 {
    responseChannel <- Response{AttackConfig: attackConfig, Passed: false, Report: fmt.Sprintf("Error occurred during HTTP Spam.")}
    return nil
  }

  passed, reason := checkHttpResponses(responses, attackConfig)

  if !passed {
    responseChannel <- Response{AttackConfig: attackConfig, Passed: false, Report: fmt.Sprintf("Failure during HTTP Spam. %s", reason)}
    return nil
  }

  responseChannel <- Response{AttackConfig: attackConfig, Passed: true}
  return nil
}

func RunCorruptHttp(endpointConfig EndpointConfig, attackConfig AttackConfig, responseChannel chan Response) error {
  fmt.Printf("🔥 Running Corrupt HTTP against %s 🔥\n", endpointConfig.Name)
  c := make(chan string)
  endpoint := BuildNetworkPath(endpointConfig.Protocol, endpointConfig.Host, endpointConfig.Port, endpointConfig.Path)

  go SendCorruptHttpData(endpoint, c)

  rawResponse := <- c

  if rawResponse == "" {
    responseChannel <- Response{AttackConfig: attackConfig, Passed: false, Report: fmt.Sprintf("Error occurred during corrupt HTTP attack. Expected valid response but got empty String.", rawResponse)}
    return nil
  }
  
  if !strings.Contains(rawResponse, attackConfig.ExpectedStatus) {
    responseChannel <- Response{AttackConfig: attackConfig, Passed: false, Report: fmt.Sprintf("Failure during Corrupt HTTP. Expected Status = %s | Actual Response = %s", attackConfig.ExpectedStatus, rawResponse)}
  }

  responseChannel <- Response{AttackConfig: attackConfig, Passed: true, Report: fmt.Sprintf("Corrupt HTTP Test passed for endpoint %s", endpointConfig.Name)}
  return nil
}

func RunRandomRabbitJson(endpointConfig EndpointConfig, attackConfig AttackConfig, responseChannel chan Response) error {
  fmt.Printf("🔥 Running Random Rabbit JSON Attack against %s 🔥\n", endpointConfig.Name)
  responseChannel <- Response{AttackConfig: attackConfig, Passed: true, Report: fmt.Sprintf("Random JSON to Rabbit MQ passed for endpoint %s", endpointConfig.Name)}
  return nil
}
