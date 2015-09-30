package main

import (
  "testing"
)

func TestCalcCorrect(t *testing.T) {
  // assert equality
  actual := add(2, 3)
  expected := 5

  if (actual != expected) {
    t.Error("Expected ", expected," got ", actual)
  }
}

func TestCalcInCorrect(t *testing.T) {
  // assert equality
  actual := add(2, 3)
  expected := 6

  if (actual == expected) {
    t.Error("Expected ", expected," got ", actual)
  }
}
