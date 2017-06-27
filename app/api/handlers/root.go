package handlers

import (
  "net/http"
  "github.com/gernest/alien"
)

func InitRootHandler(mux *alien.Mux) {
  mux.Get("/", handleRootGet)
}

func handleRootGet(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("The Fuzz Monkey has no root. He has no beginning or end. He is endless. Formless. Fuzzy."))
}
