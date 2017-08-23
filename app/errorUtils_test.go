package main

import (
  "testing"
)

func TestCheckError(t *testing.T) {
  t.Run("Test that check does not panic when the error is not present", func(t *testing.T) {
    var err error
    CheckError(err)
  })
}
