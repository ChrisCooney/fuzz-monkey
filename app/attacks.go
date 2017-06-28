package main

import (
  "fmt"
  "net/http"
)

type Response struct {
  Passed bool
  Report string
}

func checkHttpResponse(httpResponse *http.Response, config *EndpointConfig) (bool, string) {
  if httpResponse.Status != config.SuccessStatus {
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

  for i:=0; i < 100; i++ {
    resp, err := SendRandomHttpRequest(endpointConfig.Endpoint)

    if err != nil {
      report := fmt.Sprintf("Error occurred firing spam at %s after %d requests. Error= %s", endpointConfig.Endpoint, i, err)
      return Response{Passed: false, Report: report}
    }

    responses = append(responses, resp)
  }

  passed, reason := checkHttpResponses(responses, endpointConfig)

  if !passed {
    return Response{Passed: false, Report: fmt.Sprintf("Failure during Garbage Spam. %s", reason)}
  }

  return Response{Passed: true}
}
