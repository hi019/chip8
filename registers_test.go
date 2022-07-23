package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister_GetBC(t *testing.T) {
	t.Run("different lengths", func(t *testing.T) {
		r := Registers{b: 0b1100100, c: 0b10010110}
		assert.EqualValues(t, 0b01100100_10010110, r.GetBC())
	})

	t.Run("simple", func(t *testing.T) {
		r := Registers{b: 0b11011010, c: 0b01101101}
		assert.EqualValues(t, 0b11011010_01101101, r.GetBC())
	})
}
