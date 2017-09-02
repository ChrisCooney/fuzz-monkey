package main

import (
  "testing"
  "gopkg.in/jarcoal/httpmock.v1"
  "net"
  "fmt"
)

func TestRunHTTPSpam(t *testing.T) {
  t.Run("Test run HTTP spam correctly reports on endpoints", func(t *testing.T) {
    createMockHTTPServer(200)
    defer httpmock.DeactivateAndReset()
    c := createResponseChannel()

    endpoint, attack := CreateTestEndpointAndAttackConfiguration("200", "HTTP_SPAM")

    go RunHTTPSpam(endpoint, attack, c)

    response := <- c

    if !response.Passed {
      t.Error("Valid config provided to HTTP Spam. HTTP SPAM did not return a passing report.")
    }

    endpoint, attack = CreateTestEndpointAndAttackConfiguration("404", "HTTP_SPAM")

    go RunHTTPSpam(endpoint, attack, c)

    response = <- c

    if response.Passed {
      t.Error("Failing config provided. Should have failed on expected status compare but returned passing report.")
    }
  })
}

func TestRunCorruptHTTP(t *testing.T) {
  t.Run("Test run Corrupt HTTP correctly reports on endpoints", func(t *testing.T) {
    go createMockTcpServer()
    c := createResponseChannel()
    endpoint, attack := CreateTestEndpointAndAttackConfiguration("200", "CORRUPT_HTTP")

    go RunCorruptHTTP(endpoint, attack, c)

    response := <- c

    if !response.Passed {
      t.Errorf("Valid config provided to corrupt HTTP. Corrupt HTTP did not return a passing report. %v", response)
    }

    endpoint, attack = CreateTestEndpointAndAttackConfiguration("404", "CORRUPT_HTTP")

    go RunCorruptHTTP(endpoint, attack, c)

    response = <- c

    if response.Passed {
      t.Error("Invalid config provided. Should have failed on status compare but returned passed")
    }
  })
}

func TestRunURLQuery(t *testing.T) {
  t.Run("Test run URL Query attack correctly reports on endpoints", func(t *testing.T) {
    // Spin up a HTTP server with the right values
    createMockHTTPServer(400)
    defer httpmock.DeactivateAndReset()
    c := createResponseChannel()

    endpoint, attack := CreateTestEndpointAndAttackConfiguration("400", "URL_QUERY_SPAM")

    go RunURLQuery(endpoint, attack, c)

    response := <- c

    if !response.Passed {
      t.Error("Valid config which should have passed provided, but reported as failing")
    }

    endpoint, attack = CreateTestEndpointAndAttackConfiguration("200", "URL_QUERY_SPAM")

    go RunURLQuery(endpoint, attack, c)

    response = <- c

    if response.Passed {
      t.Error("Valid config which should not have passed provided, but reported as passing.")
    }
  })
}

func createMockHTTPServer(status int) {
  httpmock.Activate()

  httpmock.RegisterResponder("GET", "http://localhost:8080/my-endpoint",
    httpmock.NewStringResponder(status, `[{"something": 1}]`))
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

        count ++

        if(count == 2) {
          return
        }
    }
}

func createResponseChannel() (chan Response) {
  return make(chan Response)
}
