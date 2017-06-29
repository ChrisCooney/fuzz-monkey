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
    passed, reason := checkHttpResponse(httpResp, config)
    if !passed {
      return passed, reason
    }
  }

  return true, ""
}

func dispatchMultipleHttpRequests(endpoint string, c chan *http.Response, count int) {
  for i := 0; i < count; i++ {
    _, err := SendRandomHttpRequest(endpoint, c)
    if err != nil {
      fmt.Printf("Attempt %d failed with error %v\n", i, err)
    }
  }
}

func RunHttpSpam(endpointConfig EndpointConfig, responseChannel chan Response) error {
  fmt.Printf("Running HTTP Spam against %s\n", endpointConfig.Name)

  responses := []*http.Response{}
  c := make(chan *http.Response)

  NUM_OF_CONCURRENTS := 20
  MESSAGES_PER_CONCURRENT := 500

  for i:=0; i < NUM_OF_CONCURRENTS; i++ {
    go dispatchMultipleHttpRequests(endpointConfig.Endpoint, c, MESSAGES_PER_CONCURRENT)
  }

  for len(responses) < (NUM_OF_CONCURRENTS * MESSAGES_PER_CONCURRENT) {
    responses = append(responses, <- c)
  }

  passed, reason := checkHttpResponses(responses, endpointConfig)

  if !passed {
    responseChannel <- Response{Passed: false, Report: fmt.Sprintf("Failure during HTTP Spam. %s", reason)}
    return nil
  }

  responseChannel <- Response{Passed: true}
  return nil
}
