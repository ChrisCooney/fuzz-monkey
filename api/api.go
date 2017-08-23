package api

import (
    "net/http"
    "github.com/gernest/alien"
)

func InitApi() {
  multiplexer := alien.New()
  initRoutes(multiplexer)
  http.ListenAndServe(":8080", multiplexer)
}

func initRoutes(multiplexer *alien.Mux) {
  InitRootHandler(multiplexer)
}
