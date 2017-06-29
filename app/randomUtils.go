package main

import (
  "math/rand"
)

func CreateRandomString(length int) string {
  b := make([]byte, length)
  rand.Read(b)
  return string(b)
}
