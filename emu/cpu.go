package emu

import (
	"io"
	"math/rand"
	"os"
)

const (
	VideoHeight = 32
	VideoWidth  = 64
)

type CPU struct {
	Registers [16]byte
	Memory    [4096]byte
	Key       [16]byte
	Video     [64 * 32]uint32
	Opcode    uint16
	// the program counter, which tells us what instruction the cpu is currently executing
	PC uint16
	// stack pointer
	SP    uint16
	Stack [16]uint16
	// delay timer
	DT byte
	// sound timer
	ST byte
	// flag register
	Flags FlagRegister
	// index
	I      uint16
	Keypad [16]uint8
}

func (c *CPU) Init() {
	c.PC = 0x200

	// init fonts, starting at 0x50
	for i, b := range fontset {
		c.Memory[FontsetStartAddress+i] = b
	}
}

func (c *CPU) LoadROM(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	for i := 0; true; i++ {
		// rom starts at 0x200
		_, err := f.Read(c.Memory[0x200:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	return nil
}

func (c *CPU) Random() uint8 {
	return uint8(rand.Intn(255))
}
