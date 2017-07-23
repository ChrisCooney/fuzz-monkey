package main

import (
  "bytes"
  "net/http"
  "net"
  "math/rand"
  "fmt"
  "bufio"
)

func BuildHttpUrl(host string, port string, path string) string {
  return buildNetworkPath("http", host, port, path)
}

func BuildTcpUrl(host string, port string, path string) string {
  return fmt.Sprintf("%s:%s", host, port)
}

func buildNetworkPath(protocol string, host string, port string, path string) string {
  return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func SendRandomHttpRequest(endpoint string, c chan *http.Response) (*http.Response, error) {
  client := &http.Client{}
  method := getRandomRequestMethod()

  request, _ := http.NewRequest(method, endpoint, bytes.NewBufferString("hello"))

  response, _ := client.Do(request)

  c <- response
  return response, nil
}

var MAX_JUNK_LENGTH = 100

func SendCorruptHttpData(endpoint string, c chan string) error {
  conn, err := net.Dial("tcp", endpoint)

  if err != nil {
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
