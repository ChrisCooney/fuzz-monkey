package main

import (
  "testing"
  "os"
  "os/exec"
)

func TestSetupAttackThreads(t *testing.T) {
  t.Run("Test that the monkey correctly reports failing attacks", func(t *testing.T) {
    MAX_TIME_BETWEEN_ATTACKS = 20
    responseChannel := make(chan Response)
    config := CreateFullTestConfiguration("200", "HTTP_SPAM")

    go SetupTargets(&config, responseChannel)

    response := <- responseChannel

    if response.Passed {
      t.Error("Expected test response to not be passing. Test has passing = true")
    }
  })
}

func TestPerformSequentialAttack(t *testing.T) {
  t.Run("Test that the monkey correctly runs each of the attacks in config", func(t *testing.T) {
    ok := getStatusFromSequentialAttack()

    if ok {
      t.Error("Expected error status from the command line. Got ok status.")
    }
  })
}

func getStatusFromSequentialAttack() bool {
  config := CreateFullTestConfiguration("200", "HTTP_SPAM")
  if os.Getenv("TEST_SEQ_ATTACK") == "1" {
      PerformSequentialAttack(&config)
      return false
  }

  cmd := exec.Command(os.Args[0])
  cmd.Env = append(os.Environ(), "TEST_SEQ_ATTACK=1")
  err := cmd.Run()
  _, ok := err.(*exec.ExitError)

  return ok
}
