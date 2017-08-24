package main

import (
  "testing"
)

func TestSetupAttackThreads(t *testing.T) {
  t.Run("Test that the monkey correctly reports failing attacks", func(t *testing.T) {
    MAX_TIME_BETWEEN_ATTACKS = 20
    responseChannel := make(chan Response)
    config := CreateFullTestConfiguration()

    go SetupTargets(&config, responseChannel)

    response := <- responseChannel

    if response.Passed {
      t.Error("Expected test response to not be passing. Test has passing = true")
    }
  })

}
