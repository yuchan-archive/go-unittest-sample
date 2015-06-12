package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
  // assert equality
  assert.Equal(t, add(2, 3), 5, "they should be equal")
  assert.NotEqual(t, add(2, 3), 6, "they should not be equal")
}
