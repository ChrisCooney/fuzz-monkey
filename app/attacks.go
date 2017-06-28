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

func checkHttpResponse(httpResponse *http.Response, config *EndpointConfig) (bool, string) {
  if strings.Trim(httpResponse.Status, " ") != strings.Trim(config.SuccessStatus, " ") {
    reason := fmt.Sprintf("Invalid status code of %s detected. Expected %s", httpResponse.Status, config.SuccessStatus)
    return false, reason
  }

  return true, ""
}

func checkHttpResponses(httpResponses []*http.Response, config *EndpointConfig) (bool, string) {
  for _,httpResp := range httpResponses {
    passed, reason := checkHttpResponse(httpResp, config)
    if !passed {
      return passed, reason
    }
  }

  return true, ""
}

func RunGarbageSpam(endpointConfig *EndpointConfig) Response {
  fmt.Printf("Running Garbage Spam against %s\n", endpointConfig.Name)

  responses := []*http.Response{}
  c := make(chan *http.Response)

  for i:=0; i < 1000; i++ {
    go SendRandomHttpRequest(endpointConfig.Endpoint, c)
    resp := <- c
    responses = append(responses, resp)
  }

  passed, reason := checkHttpResponses(responses, endpointConfig)

  if !passed {
    return Response{Passed: false, Report: fmt.Sprintf("Failure during Garbage Spam. %s", reason)}
  }

  return Response{Passed: true}
}
