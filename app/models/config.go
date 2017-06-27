package models

type Config struct {
  Endpoints []EndpointConfig `json:"endpoints"`
}

type EndpointConfig struct {
  Name string `json:"name"`
  Endpoint string `json:"endpoint"`
  MaxResponseTime int `json:"maxResponseTime"`
  MinResponseTime int `json:"minResponseTime"`
  SuccessCode int `json:"successCode"`
}
