package main

import (
  "bytes"
  "net/http"
  "net"
  "math/rand"
  "fmt"
  "bufio"
)

func BuildTcpUrl(host string, port string, path string) string {
  return fmt.Sprintf("%s:%s%s", host, port, path)
}

func BuildNetworkPath(protocol string, host string, port string, path string) string {

  if protocol == "" {
    return fmt.Sprintf("%s:%s%s", host, port, path)
  }

  return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func SendHTTPRequest(endpoint string, c chan *http.Response, method string) (*http.Response, error) {
  client := &http.Client{}
  request, _ := http.NewRequest(method, endpoint, bytes.NewBufferString("hello"))

  response, _ := client.Do(request)

  c <- response
  return response, nil
}

func SendRandomHTTPRequest(endpoint string, c chan *http.Response) (*http.Response, error) {
  method := getRandomRequestMethod()
  return SendHTTPRequest(endpoint, c, method)
}

var MAX_JUNK_LENGTH = 100

func SendCorruptHTTPData(endpoint string, c chan string) error {
  conn, err := net.Dial("tcp", endpoint)

  if err != nil {
    fmt.Printf("%v", err)
    c <- ""
    return err
  }

  junkLength := rand.Intn(MAX_JUNK_LENGTH)
  junkStr := CreateRandomString(junkLength)

  fmt.Fprintf(conn, "%s\n", junkStr)
  status, err := bufio.NewReader(conn).ReadString('\n')
  c <- status
  return err
}

func getRandomRequestMethod() string {
  diceRoll := rand.Intn(100)

  if diceRoll < 25 {
    return "GET"
  } else if diceRoll < 50 {
    return "POST"
  } else {
    return "HEAD"
  }
}
