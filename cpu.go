package main

import "github.com/johncgriffin/overflow"

type CPU struct {
	registers Registers
}

type ArithmeticTarget string

const (
	ArithmeticTargetA ArithmeticTarget = "A"
	ArithmeticTargetB ArithmeticTarget = "b"
	ArithmeticTargetC ArithmeticTarget = "c"
	ArithmeticTargetD ArithmeticTarget = "d"
	ArithmeticTargetE ArithmeticTarget = "e"
	ArithmeticTargetH ArithmeticTarget = "h"
	ArithmeticTargetL ArithmeticTarget = "l"
)

type Instruction string

const (
	InstructionAdd Instruction = "add"
)

func (c *CPU) Execute(inst Instruction, targ ArithmeticTarget) {
	switch inst {
	case InstructionAdd:
		switch targ {
		case ArithmeticTargetC:
			c.add(c.registers.c)
		}
	}
}

func (c *CPU) add(val uint8) {
	res, didOverflow := overflow.Add8(int8(c.registers.a), int8(val))
	c.registers.a = uint8(res)

	c.registers.f.zero = res == 0
	c.registers.f.subtract = res == 0
	c.registers.f.carry = didOverflow
	c.registers.f.halfCarry = didOverflow

	// half carry is set if adding the lower nibbles of the value and register A
	// together result in a value bigger than 0xF (0b1111 / 0b00001111). if the result is larger than 0xF
	// than the addition caused a carry from the lower nibble to the upper nibble.
	c.registers.f.halfCarry = (c.registers.a&0xF)+(val&0xF) > 0xF
}
