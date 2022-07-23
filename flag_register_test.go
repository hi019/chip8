package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagRegister_ToUint8(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		reg := FlagRegister{zero: true}
		assert.EqualValues(t, 0b1000_0000, reg.ToUint8())
	})
	t.Run("subtract", func(t *testing.T) {
		reg := FlagRegister{subtract: true}
		assert.EqualValues(t, 0b0100_0000, reg.ToUint8())
	})
	t.Run("half carry", func(t *testing.T) {
		reg := FlagRegister{halfCarry: true}
		assert.EqualValues(t, 0b0010_0000, reg.ToUint8())
	})
	t.Run("carry", func(t *testing.T) {
		reg := FlagRegister{carry: true}
		assert.EqualValues(t, 0b0001_0000, reg.ToUint8())
	})
}

func TestFlagRegister_ParseUint8Uint8(t *testing.T) {
	t.Run("a", func(t *testing.T) {
		reg := &FlagRegister{}
		reg.ParseUint8(0b0101_0000)

		assert.EqualValues(t, &FlagRegister{subtract: true, carry: true}, reg)
	})
	t.Run("a", func(t *testing.T) {
		reg := &FlagRegister{}
		reg.ParseUint8(0b1111_0000)

		assert.EqualValues(t, &FlagRegister{subtract: true, carry: true, halfCarry: true, zero: true}, reg)
	})
	t.Run("a", func(t *testing.T) {
		reg := &FlagRegister{}
		reg.ParseUint8(0b1001_0000)

		assert.EqualValues(t, &FlagRegister{zero: true, carry: true}, reg)
	})
}
