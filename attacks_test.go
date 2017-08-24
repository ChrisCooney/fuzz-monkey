package main

import (
  "testing"
  "gopkg.in/jarcoal/httpmock.v1"
  "net"
  "fmt"
)

func TestRunHttpSpam(t *testing.T) {
  t.Run("Test run HTTP spam correctly reports on endpoints", func(t *testing.T) {
    createMockHttpServer()
    defer httpmock.DeactivateAndReset()
    c := createResponseChannel()

    endpoint, attack := createConfiguration("200")

    go RunHttpSpam(endpoint, attack, c)

    response := <- c

    if !response.Passed {
      t.Error("Valid config provided to HTTP Spam. HTTP SPAM did not return a passing report.")
    }

    endpoint, attack = createConfiguration("404")

    go RunHttpSpam(endpoint, attack, c)

    response = <- c

    if response.Passed {
      t.Error("Failing config provided. Should have failed on expected status compare but returned passing report.")
    }
  })
}

func TestRunCorruptHttp(t *testing.T) {
  t.Run("Test run Corrupt HTTP correctly reports on endpoints", func(t *testing.T) {
    go createMockTcpServer()
    c := createResponseChannel()
    endpoint, attack := createConfiguration("200")

    go RunCorruptHttp(endpoint, attack, c)

    response := <- c

    if !response.Passed {
      t.Errorf("Valid config provided to corrupt HTTP. Corrupt HTTP did not return a passing report. %v", response)
    }

    endpoint, attack = createConfiguration("404")

    go RunCorruptHttp(endpoint, attack, c)

    response = <- c

    if response.Passed {
      t.Error("Invalid config provided. Should have failed on status compare but returned passed")
    }
  })
}

func createMockHttpServer() {
  httpmock.Activate()

  httpmock.RegisterResponder("GET", "http://localhost:8080/my-endpoint",
    httpmock.NewStringResponder(200, `[{"something": 1}]`))
}

func createMockTcpServer() {
  l, _ := net.Listen("tcp", ":8080")
  count := 0
  defer l.Close()
    for {
        conn, err := l.Accept()
        if err != nil {
            return
        }

        fmt.Println("Mock TCP Server returning 200")
        conn.Write([]byte("MOCK RESPONSE: 200\n"))
        defer conn.Close()

        if(count == 1) {
          return
        } else {
          count ++
        }
    }
}

func createConfiguration(responseStatus string) (EndpointConfig, AttackConfig) {
  endpoint := EndpointConfig{Name:"Test Endpoint", Protocol:"http",Host:"localhost",Port:"8080",Path:"/my-endpoint"}
  attack := AttackConfig{Type:"HTTP_SPAM",Concurrents:1,MessagesPerConcurrent:1,ExpectedStatus:responseStatus,Method:"GET"}

  return endpoint, attack
}

func createResponseChannel() (chan Response) {
  return make(chan Response)
}
