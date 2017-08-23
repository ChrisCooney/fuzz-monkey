package main

import (
  "testing"
)

func TestBuildNetworkPath(t *testing.T) {
  t.Run("Build Network Path correctly builds path", func(t *testing.T) {
    host := "google.com"
    protocol := "http"
    port := "8080"
    path := "/chris"

    url := BuildNetworkPath(protocol, host, port, path)
    expectedUrl := "http://google.com:8080/chris"

    if url != expectedUrl  {
      t.Errorf("Expected %s, Got %s", expectedUrl, url)
    }
  })
}

func TestBuildTcpUrl(t *testing.T) {
  t.Run("Test BuildTcpUrl builds tcp url correctly", func(t *testing.T) {
    host := "google.com"
    port := "8080"
    path := "/chris"

    url := BuildTcpUrl(host, port, path)
    expectedUrl := "google.com:8080/chris"

    if url != expectedUrl {
      t.Errorf("Expected %s, Got %s", expectedUrl, url)
    }
  })
}
