package api

import (
    "net/http"
    "fuzz-monkey/api/handlers"
    "github.com/gernest/alien"
)

func Init() {
  multiplexer := alien.New()
  initRoutes(multiplexer)
  http.ListenAndServe(":8080", multiplexer)
}

func initRoutes(multiplexer *alien.Mux) {
  handlers.InitRootHandler(multiplexer)
}
