package main

import (
  "testing"
)

func TestCreateRandomString(t *testing.T) {
  t.Run("Test random string is correct length", func(t *testing.T) {
    randomString := CreateRandomString(10)

    if len(randomString) != 10 {
      t.Errorf("Expected Random string of length 10. Got random string %s of length %v", randomString, len(randomString))
    }
  })
}
