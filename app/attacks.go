package main

type Response struct {
  Passed bool
  Response string
}

func RunCorrupt(endpoint *EndpointConfig) Response {
    return Response{Passed: true, Response: "hello"}
}
