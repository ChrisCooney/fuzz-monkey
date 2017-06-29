package main

import (
  "fmt"
  "net/http"
  "strings"
)

type Response struct {
  Passed bool
  Report string
}

func checkHttpResponse(httpResponse *http.Response, config EndpointConfig) (bool, string) {
  if strings.Trim(httpResponse.Status, " ") != strings.Trim(config.SuccessStatus, " ") {
    reason := fmt.Sprintf("Invalid status code of %s detected. Expected %s", httpResponse.Status, config.SuccessStatus)
    return false, reason
  }

  /* TODO: Check response time against max response time */

  return true, ""
}

func checkHttpResponses(httpResponses []*http.Response, config EndpointConfig) (bool, string) {
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
  }

  return responses
}

func readResponseFromChannel(responses []*http.Response, c chan *http.Response) []*http.Response {
  response := <- c
  defer response.Body.Close()
  return append(responses, response)
}

var NUM_OF_CONCURRENTS = 20
var MESSAGES_PER_CONCURRENT = 100

func RunHttpSpam(endpointConfig EndpointConfig, responseChannel chan Response) error {
  fmt.Printf("Running HTTP Spam against %s\n", endpointConfig.Name)

  responses := []*http.Response{}
  c := make(chan *http.Response)

  endpoint := BuildHttpUrl(endpointConfig.Host, endpointConfig.Port, endpointConfig.Path)

  dispatchConcurrentHttpRequests(NUM_OF_CONCURRENTS, endpoint, c, MESSAGES_PER_CONCURRENT)
  collectConcurrentHttpResponses(responses, c, NUM_OF_CONCURRENTS * MESSAGES_PER_CONCURRENT)

  passed, reason := checkHttpResponses(responses, endpointConfig)

  if !passed {
    responseChannel <- Response{Passed: false, Report: fmt.Sprintf("Failure during HTTP Spam. %s", reason)}
    return nil
  }

  responseChannel <- Response{Passed: true}
  return nil
}

func RunCorruptHttp(endpointConfig EndpointConfig, responseChannel chan Response) error {
  fmt.Printf("Running Corrupt HTTP against %s\n", endpointConfig.Name)
  c := make(chan string)
  endpoint := BuildTcpUrl(endpointConfig.Host, endpointConfig.Port, endpointConfig.Path)
  go SendCorruptHttpData(endpoint, c)

  response := <- c

  fmt.Printf("Output was %s", response)

  return nil
}
